## MTProto Client

Golang MTProto client (with personalized modification)

TL layer: 116

Commit: https://github.com/tdlib/td/commit/5b69e72b097573bda599ff040616eb735d2a2888

## Usage
 
Setup Client
```golang
config := mtproto.NewAppConfig(appID, appHash)
m := mtproto.NewMTProto(config)
m.SetSession(session) // use session
m.UseIPv6(true) // use ipv6

if err := m.InitSessAndConnect(); err != nil {
	return err
}
```

Login (as Bot)
```golang
for {
	res := m.SendSync(mtproto.UpdatesGetState{})
	if mtproto.IsErrorType(res, mtproto.ErrUnauthorized) { 
		if err := m.AuthBot(token); err != nil {
			log.Errorln(err)
		}
		continue
	}
	_, ok := res.(mtproto.UpdatesState)
	if !ok {
		log.Errorln(err)
        continue
	}	
	break
}
```

Login (as User)
```golang
for {
	res := m.SendSync(mtproto.UpdatesGetState{})
	if mtproto.IsErrorType(res, mtproto.ErrUnauthorized) { 
		if err := m.Auth(mtproto.BaseAuth{}); err != nil {
			log.Errorln(err)
		}
		continue
	}
	_, ok := res.(mtproto.UpdatesState)
	if !ok {
		log.Errorln(err)
        continue
	}	
	break
}
```

Save Session
```golang
_ = m.CopySession().Save("cred.json")
```

Handle Updates
```golang
m.SetEventsHandler(updateHandler)

func updateHandler(update mtproto.TL) {
	switch u := update.(type) {
	case mtproto.UpdateNewChannelMessage:
		log.Printf("Channel Message %+v\n", u.Message)
	case mtproto.UpdateNewMessage:
		log.Printf("Message %+v\n", u.Message)
	case mtproto.UpdateChatParticipantAdd:
		log.Printf("New Member: %d %d\n", u.ChatID, u.UserID)
	case mtproto.Updates:
		for _, item := range u.Updates {
			updateHandler(item)
		}
	default:
		log.Printf("Unhandled Updates: %T", update)
	}
}
```

Send Message
```golang
userID := 391829189

m.Send(mtproto.MessagesSendMessage{
	Peer:     mtproto.InputPeerUser{UserID: userID},
	RandomID: rand.Int63(),
	Message:  "ACK",
})
```

Send Buttons
```golang
button1 := mtproto.KeyboardButtonCallback{
    Text: "OK?",	
    Data: []byte("ok"),
}
m.Send(mtproto.MessagesSendMessage{
	Flags:    1<<2,
	Peer:     mtproto.InputPeerUser{UserID: msg.FromID},
	RandomID: rand.Int63(),
	Message:  "Button Test",
	ReplyMarkup: mtproto.ReplyInlineMarkup{
		Rows: mtproto.SliceToTLStable([]mtproto.KeyboardButtonRow{
			{
				Buttons: []mtproto.TL{button1, button1, button1},
			},
		}),
	},
})
```

Upload files(Max: 2 GB)
```golang
fileName := ""

fs, err := os.Stat(fileName)
if err!= nil {
	panic(err)
}
size := fs.Size()
	
file, err := os.Open(fileName)
if err != nil {
	panic(err)
}
	
ps := int32(0)
bs := []int64{524288, 262144, 131072, 65536, 32768, 16384, 8192, 4096, 1024, 512, 64, 2}
var bsf int32
for _, p := range bs {
	if size / p < 3000 && size / p > 2 {
		bsf = int32(p)
		break
	}
}
total := int32(math.Ceil(float64(size) / float64(bsf)))
buf := make([]byte, bsf)
for {
	n, err := file.Read(buf)
	if err != nil {
		if err != io.EOF {
			panic(err)
		}
		break
	}
	log.Debug("Send FilePart: ", fileID, ps, total, n)
	res := m.SendSyncRetry(mtproto.UploadSaveBigFilePart{
		FileID:         fileID,
		FilePart:       ps,
		FileTotalParts: total,
		Bytes:          buf[:n],
	}, time.Second,5, time.Minute)
	log.Debug(res)
	ps++
}

res := mtproto.InputFileBig{
	ID:    fileID,
	Parts: total,
	Name:  fs.Name(),
}
```

Send Media (Using uploaded file)
```golang
attr := mtproto.DocumentAttributeVideo{
	Flags:             1 << 1,
	Duration:          120,
	SupportsStreaming: true,
	H: 720,
	W: 1280,
}
r := m.SendSync(mtproto.MessagesSendMedia{
	Peer:     mtproto.InputPeerUser{UserID: userID},
	RandomID: rand.Int63(),
	Message:      "Test",
	Media: mtproto.InputMediaUploadedDocument{
		File:         res,	
		MimeType:     "video/mp4",
		Attributes:   []mtproto.TL{attr},
	},
})
```


## Information

* https://core.telegram.org/mtproto
* https://core.telegram.org/api

## Acknowledgement

* https://github.com/3bl3gamer/tgclient
* https://github.com/sdidyk/mtproto
* https://github.com/ronaksoft/mtproto
* https://github.com/shelomentsevd/mtproto

