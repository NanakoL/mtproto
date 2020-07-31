package mtproto

import (
	"github.com/ansel1/merry"
)

const (
	Layer                                                                = 116
	crcResPQ                                                             = 0x05162463
	crcPQInnerDataDc                                                     = 0xa9f55f95
	crcPQInnerDataTempDc                                                 = 0x56fddf88
	crcServerDHParamsFail                                                = 0x79cb045d
	crcServerDHParamsOk                                                  = 0xd0e8075c
	crcServerDHInnerData                                                 = 0xb5890dba
	crcClientDHInnerData                                                 = 0x6643b654
	crcDhGenOk                                                           = 0x3bcbf734
	crcDhGenRetry                                                        = 0x46dc1fb9
	crcDhGenFail                                                         = 0xa69dae02
	crcBindAuthKeyInner                                                  = 0x75a3f765
	crcRpcError                                                          = 0x2144ca19
	crcRpcAnswerUnknown                                                  = 0x5e2ad36e
	crcRpcAnswerDroppedRunning                                           = 0xcd78e586
	crcRpcAnswerDropped                                                  = 0xa43ad8b7
	crcFutureSalt                                                        = 0x0949d9dc
	crcFutureSalts                                                       = 0xae500895
	crcPong                                                              = 0x347773c5
	crcDestroySessionOk                                                  = 0xe22045fc
	crcDestroySessionNone                                                = 0x62d350c9
	crcNewSessionCreated                                                 = 0x9ec20908
	crcMsgsAck                                                           = 0x62d6b459
	crcBadMsgNotification                                                = 0xa7eff811
	crcBadServerSalt                                                     = 0xedab447b
	crcMsgResendReq                                                      = 0x7d861a08
	crcMsgsStateReq                                                      = 0xda69fb52
	crcMsgsStateInfo                                                     = 0x04deb57d
	crcMsgsAllInfo                                                       = 0x8cc0d131
	crcMsgDetailedInfo                                                   = 0x276d3ec6
	crcMsgNewDetailedInfo                                                = 0x809db6df
	crcDestroyAuthKeyOk                                                  = 0xf660e1d4
	crcDestroyAuthKeyNone                                                = 0x0a9f2259
	crcDestroyAuthKeyFail                                                = 0xea109b13
	crcReqPqMulti                                                        = 0xbe7e8ef1
	crcReqDHParams                                                       = 0xd712e4be
	crcSetClientDHParams                                                 = 0xf5045f1f
	crcRpcDropAnswer                                                     = 0x58e4a740
	crcGetFutureSalts                                                    = 0xb921bd04
	crcPing                                                              = 0x7abe77ec
	crcPingDelayDisconnect                                               = 0xf3427b8c
	crcDestroySession                                                    = 0xe7512126
	crcHttpWait                                                          = 0x9299359f
	crcDestroyAuthKey                                                    = 0xd1435160
	crcTrue                                                              = 0x3fedd339
	crcBoolFalse                                                         = 0xbc799737
	crcBoolTrue                                                          = 0x997275b5
	crcError                                                             = 0xc4b9f9bb
	crcIpPort                                                            = 0xd433ad73
	crcInputPeerEmpty                                                    = 0x7f3b18ea
	crcInputPeerSelf                                                     = 0x7da07ec9
	crcInputPeerChat                                                     = 0x179be863
	crcInputPeerUser                                                     = 0x7b8e7de6
	crcInputPeerChannel                                                  = 0x20adaef8
	crcInputPeerUserFromMessage                                          = 0x17bae2e6
	crcInputPeerChannelFromMessage                                       = 0x9c95f7bb
	crcInputUserEmpty                                                    = 0xb98886cf
	crcInputUserSelf                                                     = 0xf7c1b13f
	crcInputUser                                                         = 0xd8292816
	crcInputUserFromMessage                                              = 0x2d117597
	crcInputPhoneContact                                                 = 0xf392b7f4
	crcInputFile                                                         = 0xf52ff27f
	crcInputFileBig                                                      = 0xfa4f0bb5
	crcInputMediaEmpty                                                   = 0x9664f57f
	crcInputMediaUploadedPhoto                                           = 0x1e287d04
	crcInputMediaPhoto                                                   = 0xb3ba0635
	crcInputMediaGeoPoint                                                = 0xf9c44144
	crcInputMediaContact                                                 = 0xf8ab7dfb
	crcInputMediaUploadedDocument                                        = 0x5b38c6c1
	crcInputMediaDocument                                                = 0x23ab23d2
	crcInputMediaVenue                                                   = 0xc13d1c11
	crcInputMediaPhotoExternal                                           = 0xe5bbfe1a
	crcInputMediaDocumentExternal                                        = 0xfb52dc99
	crcInputMediaGame                                                    = 0xd33f43f3
	crcInputMediaInvoice                                                 = 0xf4e096c3
	crcInputMediaGeoLive                                                 = 0xce4e82fd
	crcInputMediaPoll                                                    = 0x0f94e5f1
	crcInputMediaDice                                                    = 0xe66fbf7b
	crcInputChatPhotoEmpty                                               = 0x1ca48f57
	crcInputChatUploadedPhoto                                            = 0xc642724e
	crcInputChatPhoto                                                    = 0x8953ad37
	crcInputGeoPointEmpty                                                = 0xe4c123d6
	crcInputGeoPoint                                                     = 0xf3b7acc9
	crcInputPhotoEmpty                                                   = 0x1cd7bf0d
	crcInputPhoto                                                        = 0x3bb3b94a
	crcInputFileLocation                                                 = 0xdfdaabe1
	crcInputEncryptedFileLocation                                        = 0xf5235d55
	crcInputDocumentFileLocation                                         = 0xbad07584
	crcInputSecureFileLocation                                           = 0xcbc7ee28
	crcInputTakeoutFileLocation                                          = 0x29be5899
	crcInputPhotoFileLocation                                            = 0x40181ffe
	crcInputPhotoLegacyFileLocation                                      = 0xd83466f3
	crcInputPeerPhotoFileLocation                                        = 0x27d69997
	crcInputStickerSetThumb                                              = 0x0dbaeae9
	crcPeerUser                                                          = 0x9db1bc6d
	crcPeerChat                                                          = 0xbad0e5bb
	crcPeerChannel                                                       = 0xbddde532
	crcStorageFileUnknown                                                = 0xaa963b05
	crcStorageFilePartial                                                = 0x40bc6f52
	crcStorageFileJpeg                                                   = 0x007efe0e
	crcStorageFileGif                                                    = 0xcae1aadf
	crcStorageFilePng                                                    = 0x0a4f63c0
	crcStorageFilePdf                                                    = 0xae1e508d
	crcStorageFileMp3                                                    = 0x528a0677
	crcStorageFileMov                                                    = 0x4b09ebbc
	crcStorageFileMp4                                                    = 0xb3cea0e4
	crcStorageFileWebp                                                   = 0x1081464c
	crcUserEmpty                                                         = 0x200250ba
	crcUser                                                              = 0x938458c1
	crcUserProfilePhotoEmpty                                             = 0x4f11bae1
	crcUserProfilePhoto                                                  = 0x69d3ab26
	crcUserStatusEmpty                                                   = 0x09d05049
	crcUserStatusOnline                                                  = 0xedb93949
	crcUserStatusOffline                                                 = 0x008c703f
	crcUserStatusRecently                                                = 0xe26f42f1
	crcUserStatusLastWeek                                                = 0x07bf09fc
	crcUserStatusLastMonth                                               = 0x77ebc742
	crcChatEmpty                                                         = 0x9ba2d800
	crcChat                                                              = 0x3bda1bde
	crcChatForbidden                                                     = 0x07328bdb
	crcChannel                                                           = 0xd31a961e
	crcChannelForbidden                                                  = 0x289da732
	crcChatFull                                                          = 0x1b7c9db3
	crcChannelFull                                                       = 0xf0e6672a
	crcChatParticipant                                                   = 0xc8d7493e
	crcChatParticipantCreator                                            = 0xda13538a
	crcChatParticipantAdmin                                              = 0xe2d6e436
	crcChatParticipantsForbidden                                         = 0xfc900c2b
	crcChatParticipants                                                  = 0x3f460fed
	crcChatPhotoEmpty                                                    = 0x37c1011c
	crcChatPhoto                                                         = 0xd20b9f3c
	crcMessageEmpty                                                      = 0x83e5de54
	crcMessage                                                           = 0x452c0e65
	crcMessageService                                                    = 0x9e19a1f6
	crcMessageMediaEmpty                                                 = 0x3ded6320
	crcMessageMediaPhoto                                                 = 0x695150d7
	crcMessageMediaGeo                                                   = 0x56e0d474
	crcMessageMediaContact                                               = 0xcbf24940
	crcMessageMediaUnsupported                                           = 0x9f84f49e
	crcMessageMediaDocument                                              = 0x9cb070d7
	crcMessageMediaWebPage                                               = 0xa32dd600
	crcMessageMediaVenue                                                 = 0x2ec0533f
	crcMessageMediaGame                                                  = 0xfdb19008
	crcMessageMediaInvoice                                               = 0x84551347
	crcMessageMediaGeoLive                                               = 0x7c3c2609
	crcMessageMediaPoll                                                  = 0x4bd6e798
	crcMessageMediaDice                                                  = 0x3f7ee58b
	crcMessageActionEmpty                                                = 0xb6aef7b0
	crcMessageActionChatCreate                                           = 0xa6638b9a
	crcMessageActionChatEditTitle                                        = 0xb5a1ce5a
	crcMessageActionChatEditPhoto                                        = 0x7fcb13a8
	crcMessageActionChatDeletePhoto                                      = 0x95e3fbef
	crcMessageActionChatAddUser                                          = 0x488a7337
	crcMessageActionChatDeleteUser                                       = 0xb2ae9b0c
	crcMessageActionChatJoinedByLink                                     = 0xf89cf5e8
	crcMessageActionChannelCreate                                        = 0x95d2ac92
	crcMessageActionChatMigrateTo                                        = 0x51bdb021
	crcMessageActionChannelMigrateFrom                                   = 0xb055eaee
	crcMessageActionPinMessage                                           = 0x94bd38ed
	crcMessageActionHistoryClear                                         = 0x9fbab604
	crcMessageActionGameScore                                            = 0x92a72876
	crcMessageActionPaymentSentMe                                        = 0x8f31b327
	crcMessageActionPaymentSent                                          = 0x40699cd0
	crcMessageActionPhoneCall                                            = 0x80e11a7f
	crcMessageActionScreenshotTaken                                      = 0x4792929b
	crcMessageActionCustomAction                                         = 0xfae69f56
	crcMessageActionBotAllowed                                           = 0xabe9affe
	crcMessageActionSecureValuesSentMe                                   = 0x1b287353
	crcMessageActionSecureValuesSent                                     = 0xd95c6154
	crcMessageActionContactSignUp                                        = 0xf3f25f76
	crcDialog                                                            = 0x2c171f72
	crcDialogFolder                                                      = 0x71bd134c
	crcPhotoEmpty                                                        = 0x2331b22d
	crcPhoto                                                             = 0xfb197a65
	crcPhotoSizeEmpty                                                    = 0x0e17e23c
	crcPhotoSize                                                         = 0x77bfb61b
	crcPhotoCachedSize                                                   = 0xe9a734fa
	crcPhotoStrippedSize                                                 = 0xe0b0bc2e
	crcGeoPointEmpty                                                     = 0x1117dd5f
	crcGeoPoint                                                          = 0x0296f104
	crcAuthSentCode                                                      = 0x5e002502
	crcAuthAuthorization                                                 = 0xcd050916
	crcAuthAuthorizationSignUpRequired                                   = 0x44747e9a
	crcAuthExportedAuthorization                                         = 0xdf969c2d
	crcInputNotifyPeer                                                   = 0xb8bc5b0c
	crcInputNotifyUsers                                                  = 0x193b4417
	crcInputNotifyChats                                                  = 0x4a95e84e
	crcInputNotifyBroadcasts                                             = 0xb1db7c7e
	crcInputPeerNotifySettings                                           = 0x9c3d198e
	crcPeerNotifySettings                                                = 0xaf509d20
	crcPeerSettings                                                      = 0x733f2961
	crcWallPaper                                                         = 0xa437c3ed
	crcWallPaperNoFile                                                   = 0x8af40b25
	crcInputReportReasonSpam                                             = 0x58dbcab8
	crcInputReportReasonViolence                                         = 0x1e22c78d
	crcInputReportReasonPornography                                      = 0x2e59d922
	crcInputReportReasonChildAbuse                                       = 0xadf44ee3
	crcInputReportReasonOther                                            = 0xe1746d0a
	crcInputReportReasonCopyright                                        = 0x9b89f93a
	crcInputReportReasonGeoIrrelevant                                    = 0xdbd4feed
	crcUserFull                                                          = 0xedf17c12
	crcContact                                                           = 0xf911c994
	crcImportedContact                                                   = 0xd0028438
	crcContactBlocked                                                    = 0x561bc879
	crcContactStatus                                                     = 0xd3680c61
	crcContactsContactsNotModified                                       = 0xb74ba9d2
	crcContactsContacts                                                  = 0xeae87e42
	crcContactsImportedContacts                                          = 0x77d01c3b
	crcContactsBlocked                                                   = 0x1c138d15
	crcContactsBlockedSlice                                              = 0x900802a1
	crcMessagesDialogs                                                   = 0x15ba6c40
	crcMessagesDialogsSlice                                              = 0x71e094f3
	crcMessagesDialogsNotModified                                        = 0xf0e3e596
	crcMessagesMessages                                                  = 0x8c718e87
	crcMessagesMessagesSlice                                             = 0xc8edce1e
	crcMessagesChannelMessages                                           = 0x99262e37
	crcMessagesMessagesNotModified                                       = 0x74535f21
	crcMessagesChats                                                     = 0x64ff9fd5
	crcMessagesChatsSlice                                                = 0x9cd81144
	crcMessagesChatFull                                                  = 0xe5d7d19c
	crcMessagesAffectedHistory                                           = 0xb45c69d1
	crcInputMessagesFilterEmpty                                          = 0x57e2f66c
	crcInputMessagesFilterPhotos                                         = 0x9609a51c
	crcInputMessagesFilterVideo                                          = 0x9fc00e65
	crcInputMessagesFilterPhotoVideo                                     = 0x56e9f0e4
	crcInputMessagesFilterDocument                                       = 0x9eddf188
	crcInputMessagesFilterUrl                                            = 0x7ef0dd87
	crcInputMessagesFilterGif                                            = 0xffc86587
	crcInputMessagesFilterVoice                                          = 0x50f5c392
	crcInputMessagesFilterMusic                                          = 0x3751b49e
	crcInputMessagesFilterChatPhotos                                     = 0x3a20ecb8
	crcInputMessagesFilterPhoneCalls                                     = 0x80c99768
	crcInputMessagesFilterRoundVoice                                     = 0x7a7c17a4
	crcInputMessagesFilterRoundVideo                                     = 0xb549da53
	crcInputMessagesFilterMyMentions                                     = 0xc1f8e69a
	crcInputMessagesFilterGeo                                            = 0xe7026d0d
	crcInputMessagesFilterContacts                                       = 0xe062db83
	crcUpdateNewMessage                                                  = 0x1f2b0afd
	crcUpdateMessageID                                                   = 0x4e90bfd6
	crcUpdateDeleteMessages                                              = 0xa20db0e5
	crcUpdateUserTyping                                                  = 0x5c486927
	crcUpdateChatUserTyping                                              = 0x9a65ea1f
	crcUpdateChatParticipants                                            = 0x07761198
	crcUpdateUserStatus                                                  = 0x1bfbd823
	crcUpdateUserName                                                    = 0xa7332b73
	crcUpdateUserPhoto                                                   = 0x95313b0c
	crcUpdateNewEncryptedMessage                                         = 0x12bcbd9a
	crcUpdateEncryptedChatTyping                                         = 0x1710f156
	crcUpdateEncryption                                                  = 0xb4a2e88d
	crcUpdateEncryptedMessagesRead                                       = 0x38fe25b7
	crcUpdateChatParticipantAdd                                          = 0xea4b0e5c
	crcUpdateChatParticipantDelete                                       = 0x6e5f8c22
	crcUpdateDcOptions                                                   = 0x8e5e9873
	crcUpdateUserBlocked                                                 = 0x80ece81a
	crcUpdateNotifySettings                                              = 0xbec268ef
	crcUpdateServiceNotification                                         = 0xebe46819
	crcUpdatePrivacy                                                     = 0xee3b272a
	crcUpdateUserPhone                                                   = 0x12b9417b
	crcUpdateReadHistoryInbox                                            = 0x9c974fdf
	crcUpdateReadHistoryOutbox                                           = 0x2f2f21bf
	crcUpdateWebPage                                                     = 0x7f891213
	crcUpdateReadMessagesContents                                        = 0x68c13933
	crcUpdateChannelTooLong                                              = 0xeb0467fb
	crcUpdateChannel                                                     = 0xb6d45656
	crcUpdateNewChannelMessage                                           = 0x62ba04d9
	crcUpdateReadChannelInbox                                            = 0x330b5424
	crcUpdateDeleteChannelMessages                                       = 0xc37521c9
	crcUpdateChannelMessageViews                                         = 0x98a12b4b
	crcUpdateChatParticipantAdmin                                        = 0xb6901959
	crcUpdateNewStickerSet                                               = 0x688a30aa
	crcUpdateStickerSetsOrder                                            = 0x0bb2d201
	crcUpdateStickerSets                                                 = 0x43ae3dec
	crcUpdateSavedGifs                                                   = 0x9375341e
	crcUpdateBotInlineQuery                                              = 0x54826690
	crcUpdateBotInlineSend                                               = 0x0e48f964
	crcUpdateEditChannelMessage                                          = 0x1b3f4df7
	crcUpdateChannelPinnedMessage                                        = 0x98592475
	crcUpdateBotCallbackQuery                                            = 0xe73547e1
	crcUpdateEditMessage                                                 = 0xe40370a3
	crcUpdateInlineBotCallbackQuery                                      = 0xf9d27a5a
	crcUpdateReadChannelOutbox                                           = 0x25d6c9c7
	crcUpdateDraftMessage                                                = 0xee2bb969
	crcUpdateReadFeaturedStickers                                        = 0x571d2742
	crcUpdateRecentStickers                                              = 0x9a422c20
	crcUpdateConfig                                                      = 0xa229dd06
	crcUpdatePtsChanged                                                  = 0x3354678f
	crcUpdateChannelWebPage                                              = 0x40771900
	crcUpdateDialogPinned                                                = 0x6e6fe51c
	crcUpdatePinnedDialogs                                               = 0xfa0f3ca2
	crcUpdateBotWebhookJSON                                              = 0x8317c0c3
	crcUpdateBotWebhookJSONQuery                                         = 0x9b9240a6
	crcUpdateBotShippingQuery                                            = 0xe0cdc940
	crcUpdateBotPrecheckoutQuery                                         = 0x5d2f3aa9
	crcUpdatePhoneCall                                                   = 0xab0f6b1e
	crcUpdateLangPackTooLong                                             = 0x46560264
	crcUpdateLangPack                                                    = 0x56022f4d
	crcUpdateFavedStickers                                               = 0xe511996d
	crcUpdateChannelReadMessagesContents                                 = 0x89893b45
	crcUpdateContactsReset                                               = 0x7084a7be
	crcUpdateChannelAvailableMessages                                    = 0x70db6837
	crcUpdateDialogUnreadMark                                            = 0xe16459c3
	crcUpdateUserPinnedMessage                                           = 0x4c43da18
	crcUpdateChatPinnedMessage                                           = 0xe10db349
	crcUpdateMessagePoll                                                 = 0xaca1657b
	crcUpdateChatDefaultBannedRights                                     = 0x54c01850
	crcUpdateFolderPeers                                                 = 0x19360dc0
	crcUpdatePeerSettings                                                = 0x6a7e7366
	crcUpdatePeerLocated                                                 = 0xb4afcfb0
	crcUpdateNewScheduledMessage                                         = 0x39a51dfb
	crcUpdateDeleteScheduledMessages                                     = 0x90866cee
	crcUpdateTheme                                                       = 0x8216fba3
	crcUpdateGeoLiveViewed                                               = 0x871fb939
	crcUpdateLoginToken                                                  = 0x564fe691
	crcUpdateMessagePollVote                                             = 0x42f88f2c
	crcUpdateDialogFilter                                                = 0x26ffde7d
	crcUpdateDialogFilterOrder                                           = 0xa5d72105
	crcUpdateDialogFilters                                               = 0x3504914f
	crcUpdatePhoneCallSignalingData                                      = 0x2661bf09
	crcUpdateChannelParticipant                                          = 0x65d2b464
	crcUpdatesState                                                      = 0xa56c2a3e
	crcUpdatesDifferenceEmpty                                            = 0x5d75a138
	crcUpdatesDifference                                                 = 0x00f49ca0
	crcUpdatesDifferenceSlice                                            = 0xa8fb1981
	crcUpdatesDifferenceTooLong                                          = 0x4afe8f6d
	crcUpdatesTooLong                                                    = 0xe317af7e
	crcUpdateShortMessage                                                = 0x914fbf11
	crcUpdateShortChatMessage                                            = 0x16812688
	crcUpdateShort                                                       = 0x78d4dec1
	crcUpdatesCombined                                                   = 0x725b04c3
	crcUpdates                                                           = 0x74ae4240
	crcUpdateShortSentMessage                                            = 0x11f1331c
	crcPhotosPhotos                                                      = 0x8dca6aa5
	crcPhotosPhotosSlice                                                 = 0x15051f54
	crcPhotosPhoto                                                       = 0x20212ca8
	crcUploadFile                                                        = 0x096a18d5
	crcUploadFileCdnRedirect                                             = 0xf18cda44
	crcDcOption                                                          = 0x18b7a10d
	crcConfig                                                            = 0x330b4067
	crcNearestDc                                                         = 0x8e1a1775
	crcHelpAppUpdate                                                     = 0x1da7158f
	crcHelpNoAppUpdate                                                   = 0xc45a6536
	crcHelpInviteText                                                    = 0x18cb9f78
	crcEncryptedChatEmpty                                                = 0xab7ec0a0
	crcEncryptedChatWaiting                                              = 0x3bf703dc
	crcEncryptedChatRequested                                            = 0x62718a82
	crcEncryptedChat                                                     = 0xfa56ce36
	crcEncryptedChatDiscarded                                            = 0x13d6dd27
	crcInputEncryptedChat                                                = 0xf141b5e1
	crcEncryptedFileEmpty                                                = 0xc21f497e
	crcEncryptedFile                                                     = 0x4a70994c
	crcInputEncryptedFileEmpty                                           = 0x1837c364
	crcInputEncryptedFileUploaded                                        = 0x64bd0306
	crcInputEncryptedFile                                                = 0x5a17b5e5
	crcInputEncryptedFileBigUploaded                                     = 0x2dc173c8
	crcEncryptedMessage                                                  = 0xed18c118
	crcEncryptedMessageService                                           = 0x23734b06
	crcMessagesDhConfigNotModified                                       = 0xc0e24635
	crcMessagesDhConfig                                                  = 0x2c221edd
	crcMessagesSentEncryptedMessage                                      = 0x560f8935
	crcMessagesSentEncryptedFile                                         = 0x9493ff32
	crcInputDocumentEmpty                                                = 0x72f0eaae
	crcInputDocument                                                     = 0x1abfb575
	crcDocumentEmpty                                                     = 0x36f8c871
	crcDocument                                                          = 0x1e87342b
	crcHelpSupport                                                       = 0x17c6b5f6
	crcNotifyPeer                                                        = 0x9fd40bd8
	crcNotifyUsers                                                       = 0xb4c83b4c
	crcNotifyChats                                                       = 0xc007cec3
	crcNotifyBroadcasts                                                  = 0xd612e8ef
	crcSendMessageTypingAction                                           = 0x16bf744e
	crcSendMessageCancelAction                                           = 0xfd5ec8f5
	crcSendMessageRecordVideoAction                                      = 0xa187d66f
	crcSendMessageUploadVideoAction                                      = 0xe9763aec
	crcSendMessageRecordAudioAction                                      = 0xd52f73f7
	crcSendMessageUploadAudioAction                                      = 0xf351d7ab
	crcSendMessageUploadPhotoAction                                      = 0xd1d34a26
	crcSendMessageUploadDocumentAction                                   = 0xaa0cd9e4
	crcSendMessageGeoLocationAction                                      = 0x176f8ba1
	crcSendMessageChooseContactAction                                    = 0x628cbc6f
	crcSendMessageGamePlayAction                                         = 0xdd6a8f48
	crcSendMessageRecordRoundAction                                      = 0x88f27fbc
	crcSendMessageUploadRoundAction                                      = 0x243e1c66
	crcContactsFound                                                     = 0xb3134d9d
	crcInputPrivacyKeyStatusTimestamp                                    = 0x4f96cb18
	crcInputPrivacyKeyChatInvite                                         = 0xbdfb0426
	crcInputPrivacyKeyPhoneCall                                          = 0xfabadc5f
	crcInputPrivacyKeyPhoneP2P                                           = 0xdb9e70d2
	crcInputPrivacyKeyForwards                                           = 0xa4dd4c08
	crcInputPrivacyKeyProfilePhoto                                       = 0x5719bacc
	crcInputPrivacyKeyPhoneNumber                                        = 0x0352dafa
	crcInputPrivacyKeyAddedByPhone                                       = 0xd1219bdd
	crcPrivacyKeyStatusTimestamp                                         = 0xbc2eab30
	crcPrivacyKeyChatInvite                                              = 0x500e6dfa
	crcPrivacyKeyPhoneCall                                               = 0x3d662b7b
	crcPrivacyKeyPhoneP2P                                                = 0x39491cc8
	crcPrivacyKeyForwards                                                = 0x69ec56a3
	crcPrivacyKeyProfilePhoto                                            = 0x96151fed
	crcPrivacyKeyPhoneNumber                                             = 0xd19ae46d
	crcPrivacyKeyAddedByPhone                                            = 0x42ffd42b
	crcInputPrivacyValueAllowContacts                                    = 0x0d09e07b
	crcInputPrivacyValueAllowAll                                         = 0x184b35ce
	crcInputPrivacyValueAllowUsers                                       = 0x131cc67f
	crcInputPrivacyValueDisallowContacts                                 = 0x0ba52007
	crcInputPrivacyValueDisallowAll                                      = 0xd66b66c9
	crcInputPrivacyValueDisallowUsers                                    = 0x90110467
	crcInputPrivacyValueAllowChatParticipants                            = 0x4c81c1ba
	crcInputPrivacyValueDisallowChatParticipants                         = 0xd82363af
	crcPrivacyValueAllowContacts                                         = 0xfffe1bac
	crcPrivacyValueAllowAll                                              = 0x65427b82
	crcPrivacyValueAllowUsers                                            = 0x4d5bbe0c
	crcPrivacyValueDisallowContacts                                      = 0xf888fa1a
	crcPrivacyValueDisallowAll                                           = 0x8b73e763
	crcPrivacyValueDisallowUsers                                         = 0x0c7f49b7
	crcPrivacyValueAllowChatParticipants                                 = 0x18be796b
	crcPrivacyValueDisallowChatParticipants                              = 0xacae0690
	crcAccountPrivacyRules                                               = 0x50a04e45
	crcAccountDaysTTL                                                    = 0xb8d0afdf
	crcDocumentAttributeImageSize                                        = 0x6c37c15c
	crcDocumentAttributeAnimated                                         = 0x11b58939
	crcDocumentAttributeSticker                                          = 0x6319d612
	crcDocumentAttributeVideo                                            = 0x0ef02ce6
	crcDocumentAttributeAudio                                            = 0x9852f9c6
	crcDocumentAttributeFilename                                         = 0x15590068
	crcDocumentAttributeHasStickers                                      = 0x9801d2f7
	crcMessagesStickersNotModified                                       = 0xf1749a22
	crcMessagesStickers                                                  = 0xe4599bbd
	crcStickerPack                                                       = 0x12b299d4
	crcMessagesAllStickersNotModified                                    = 0xe86602c3
	crcMessagesAllStickers                                               = 0xedfd405f
	crcMessagesAffectedMessages                                          = 0x84d19185
	crcWebPageEmpty                                                      = 0xeb1477e8
	crcWebPagePending                                                    = 0xc586da1c
	crcWebPage                                                           = 0xe89c45b2
	crcWebPageNotModified                                                = 0x7311ca11
	crcAuthorization                                                     = 0xad01d61d
	crcAccountAuthorizations                                             = 0x1250abde
	crcAccountPassword                                                   = 0xad2641f8
	crcAccountPasswordSettings                                           = 0x9a5c33e5
	crcAccountPasswordInputSettings                                      = 0xc23727c9
	crcAuthPasswordRecovery                                              = 0x137948a5
	crcReceivedNotifyMessage                                             = 0xa384b779
	crcChatInviteEmpty                                                   = 0x69df3769
	crcChatInviteExported                                                = 0xfc2e05bc
	crcChatInviteAlready                                                 = 0x5a686d7c
	crcChatInvite                                                        = 0xdfc2f58e
	crcChatInvitePeek                                                    = 0x61695cb0
	crcInputStickerSetEmpty                                              = 0xffb62b95
	crcInputStickerSetID                                                 = 0x9de7a269
	crcInputStickerSetShortName                                          = 0x861cc8a0
	crcInputStickerSetAnimatedEmoji                                      = 0x028703c8
	crcInputStickerSetDice                                               = 0xe67f520e
	crcStickerSet                                                        = 0xeeb46f27
	crcMessagesStickerSet                                                = 0xb60a24a6
	crcBotCommand                                                        = 0xc27ac8c7
	crcBotInfo                                                           = 0x98e81d3a
	crcKeyboardButton                                                    = 0xa2fa4880
	crcKeyboardButtonUrl                                                 = 0x258aff05
	crcKeyboardButtonCallback                                            = 0x683a5e46
	crcKeyboardButtonRequestPhone                                        = 0xb16a6c29
	crcKeyboardButtonRequestGeoLocation                                  = 0xfc796b3f
	crcKeyboardButtonSwitchInline                                        = 0x0568a748
	crcKeyboardButtonGame                                                = 0x50f41ccf
	crcKeyboardButtonBuy                                                 = 0xafd93fbb
	crcKeyboardButtonUrlAuth                                             = 0x10b78d29
	crcInputKeyboardButtonUrlAuth                                        = 0xd02e7fd4
	crcKeyboardButtonRequestPoll                                         = 0xbbc7515d
	crcKeyboardButtonRow                                                 = 0x77608b83
	crcReplyKeyboardHide                                                 = 0xa03e5b85
	crcReplyKeyboardForceReply                                           = 0xf4108aa0
	crcReplyKeyboardMarkup                                               = 0x3502758c
	crcReplyInlineMarkup                                                 = 0x48a30254
	crcMessageEntityUnknown                                              = 0xbb92ba95
	crcMessageEntityMention                                              = 0xfa04579d
	crcMessageEntityHashtag                                              = 0x6f635b0d
	crcMessageEntityBotCommand                                           = 0x6cef8ac7
	crcMessageEntityUrl                                                  = 0x6ed02538
	crcMessageEntityEmail                                                = 0x64e475c2
	crcMessageEntityBold                                                 = 0xbd610bc9
	crcMessageEntityItalic                                               = 0x826f8b60
	crcMessageEntityCode                                                 = 0x28a20571
	crcMessageEntityPre                                                  = 0x73924be0
	crcMessageEntityTextUrl                                              = 0x76a6d327
	crcMessageEntityMentionName                                          = 0x352dca58
	crcInputMessageEntityMentionName                                     = 0x208e68c9
	crcMessageEntityPhone                                                = 0x9b69e34b
	crcMessageEntityCashtag                                              = 0x4c4e743f
	crcMessageEntityUnderline                                            = 0x9c4e7e8b
	crcMessageEntityStrike                                               = 0xbf0693d4
	crcMessageEntityBlockquote                                           = 0x020df5d0
	crcMessageEntityBankCard                                             = 0x761e6af4
	crcInputChannelEmpty                                                 = 0xee8c1e86
	crcInputChannel                                                      = 0xafeb712e
	crcInputChannelFromMessage                                           = 0x2a286531
	crcContactsResolvedPeer                                              = 0x7f077ad9
	crcMessageRange                                                      = 0x0ae30253
	crcUpdatesChannelDifferenceEmpty                                     = 0x3e11affb
	crcUpdatesChannelDifferenceTooLong                                   = 0xa4bcc6fe
	crcUpdatesChannelDifference                                          = 0x2064674e
	crcChannelMessagesFilterEmpty                                        = 0x94d42ee7
	crcChannelMessagesFilter                                             = 0xcd77d957
	crcChannelParticipant                                                = 0x15ebac1d
	crcChannelParticipantSelf                                            = 0xa3289a6d
	crcChannelParticipantCreator                                         = 0x808d15a4
	crcChannelParticipantAdmin                                           = 0xccbebbaf
	crcChannelParticipantBanned                                          = 0x1c0facaf
	crcChannelParticipantsRecent                                         = 0xde3f3c79
	crcChannelParticipantsAdmins                                         = 0xb4608969
	crcChannelParticipantsKicked                                         = 0xa3b54985
	crcChannelParticipantsBots                                           = 0xb0d1865b
	crcChannelParticipantsBanned                                         = 0x1427a5e1
	crcChannelParticipantsSearch                                         = 0x0656ac4b
	crcChannelParticipantsContacts                                       = 0xbb6ae88d
	crcChannelsChannelParticipants                                       = 0xf56ee2a8
	crcChannelsChannelParticipantsNotModified                            = 0xf0173fe9
	crcChannelsChannelParticipant                                        = 0xd0d9b163
	crcHelpTermsOfService                                                = 0x780a0310
	crcMessagesSavedGifsNotModified                                      = 0xe8025ca2
	crcMessagesSavedGifs                                                 = 0x2e0709a5
	crcInputBotInlineMessageMediaAuto                                    = 0x3380c786
	crcInputBotInlineMessageText                                         = 0x3dcd7a87
	crcInputBotInlineMessageMediaGeo                                     = 0xc1b15d65
	crcInputBotInlineMessageMediaVenue                                   = 0x417bbf11
	crcInputBotInlineMessageMediaContact                                 = 0xa6edbffd
	crcInputBotInlineMessageGame                                         = 0x4b425864
	crcInputBotInlineResult                                              = 0x88bf9319
	crcInputBotInlineResultPhoto                                         = 0xa8d864a7
	crcInputBotInlineResultDocument                                      = 0xfff8fdc4
	crcInputBotInlineResultGame                                          = 0x4fa417f2
	crcBotInlineMessageMediaAuto                                         = 0x764cf810
	crcBotInlineMessageText                                              = 0x8c7f65e2
	crcBotInlineMessageMediaGeo                                          = 0xb722de65
	crcBotInlineMessageMediaVenue                                        = 0x8a86659c
	crcBotInlineMessageMediaContact                                      = 0x18d1cdc2
	crcBotInlineResult                                                   = 0x11965f3a
	crcBotInlineMediaResult                                              = 0x17db940b
	crcMessagesBotResults                                                = 0x947ca848
	crcExportedMessageLink                                               = 0x5dab1af4
	crcMessageFwdHeader                                                  = 0x353a686b
	crcAuthCodeTypeSms                                                   = 0x72a3158c
	crcAuthCodeTypeCall                                                  = 0x741cd3e3
	crcAuthCodeTypeFlashCall                                             = 0x226ccefb
	crcAuthSentCodeTypeApp                                               = 0x3dbb5986
	crcAuthSentCodeTypeSms                                               = 0xc000bba2
	crcAuthSentCodeTypeCall                                              = 0x5353e5a7
	crcAuthSentCodeTypeFlashCall                                         = 0xab03c6d9
	crcMessagesBotCallbackAnswer                                         = 0x36585ea4
	crcMessagesMessageEditData                                           = 0x26b5dde6
	crcInputBotInlineMessageID                                           = 0x890c3d89
	crcInlineBotSwitchPM                                                 = 0x3c20629f
	crcMessagesPeerDialogs                                               = 0x3371c354
	crcTopPeer                                                           = 0xedcdc05b
	crcTopPeerCategoryBotsPM                                             = 0xab661b5b
	crcTopPeerCategoryBotsInline                                         = 0x148677e2
	crcTopPeerCategoryCorrespondents                                     = 0x0637b7ed
	crcTopPeerCategoryGroups                                             = 0xbd17a14a
	crcTopPeerCategoryChannels                                           = 0x161d9628
	crcTopPeerCategoryPhoneCalls                                         = 0x1e76a78c
	crcTopPeerCategoryForwardUsers                                       = 0xa8406ca9
	crcTopPeerCategoryForwardChats                                       = 0xfbeec0f0
	crcTopPeerCategoryPeers                                              = 0xfb834291
	crcContactsTopPeersNotModified                                       = 0xde266ef5
	crcContactsTopPeers                                                  = 0x70b772a8
	crcContactsTopPeersDisabled                                          = 0xb52c939d
	crcDraftMessageEmpty                                                 = 0x1b0c841a
	crcDraftMessage                                                      = 0xfd8e711f
	crcMessagesFeaturedStickersNotModified                               = 0xc6dc0c66
	crcMessagesFeaturedStickers                                          = 0xb6abc341
	crcMessagesRecentStickersNotModified                                 = 0x0b17f890
	crcMessagesRecentStickers                                            = 0x22f3afb3
	crcMessagesArchivedStickers                                          = 0x4fcba9c8
	crcMessagesStickerSetInstallResultSuccess                            = 0x38641628
	crcMessagesStickerSetInstallResultArchive                            = 0x35e410a8
	crcStickerSetCovered                                                 = 0x6410a5d2
	crcStickerSetMultiCovered                                            = 0x3407e51b
	crcMaskCoords                                                        = 0xaed6dbb2
	crcInputStickeredMediaPhoto                                          = 0x4a992157
	crcInputStickeredMediaDocument                                       = 0x0438865b
	crcGame                                                              = 0xbdf9653b
	crcInputGameID                                                       = 0x032c3e77
	crcInputGameShortName                                                = 0xc331e80a
	crcHighScore                                                         = 0x58fffcd0
	crcMessagesHighScores                                                = 0x9a3bfd99
	crcTextEmpty                                                         = 0xdc3d824f
	crcTextPlain                                                         = 0x744694e0
	crcTextBold                                                          = 0x6724abc4
	crcTextItalic                                                        = 0xd912a59c
	crcTextUnderline                                                     = 0xc12622c4
	crcTextStrike                                                        = 0x9bf8bb95
	crcTextFixed                                                         = 0x6c3f19b9
	crcTextUrl                                                           = 0x3c2884c1
	crcTextEmail                                                         = 0xde5a0dd6
	crcTextConcat                                                        = 0x7e6260d7
	crcTextSubscript                                                     = 0xed6a8504
	crcTextSuperscript                                                   = 0xc7fb5e01
	crcTextMarked                                                        = 0x034b8621
	crcTextPhone                                                         = 0x1ccb966a
	crcTextImage                                                         = 0x081ccf4f
	crcTextAnchor                                                        = 0x35553762
	crcPageBlockUnsupported                                              = 0x13567e8a
	crcPageBlockTitle                                                    = 0x70abc3fd
	crcPageBlockSubtitle                                                 = 0x8ffa9a1f
	crcPageBlockAuthorDate                                               = 0xbaafe5e0
	crcPageBlockHeader                                                   = 0xbfd064ec
	crcPageBlockSubheader                                                = 0xf12bb6e1
	crcPageBlockParagraph                                                = 0x467a0766
	crcPageBlockPreformatted                                             = 0xc070d93e
	crcPageBlockFooter                                                   = 0x48870999
	crcPageBlockDivider                                                  = 0xdb20b188
	crcPageBlockAnchor                                                   = 0xce0d37b0
	crcPageBlockList                                                     = 0xe4e88011
	crcPageBlockBlockquote                                               = 0x263d7c26
	crcPageBlockPullquote                                                = 0x4f4456d3
	crcPageBlockPhoto                                                    = 0x1759c560
	crcPageBlockVideo                                                    = 0x7c8fe7b6
	crcPageBlockCover                                                    = 0x39f23300
	crcPageBlockEmbed                                                    = 0xa8718dc5
	crcPageBlockEmbedPost                                                = 0xf259a80b
	crcPageBlockCollage                                                  = 0x65a0fa4d
	crcPageBlockSlideshow                                                = 0x031f9590
	crcPageBlockChannel                                                  = 0xef1751b5
	crcPageBlockAudio                                                    = 0x804361ea
	crcPageBlockKicker                                                   = 0x1e148390
	crcPageBlockTable                                                    = 0xbf4dea82
	crcPageBlockOrderedList                                              = 0x9a8ae1e1
	crcPageBlockDetails                                                  = 0x76768bed
	crcPageBlockRelatedArticles                                          = 0x16115a96
	crcPageBlockMap                                                      = 0xa44f3ef6
	crcPhoneCallDiscardReasonMissed                                      = 0x85e42301
	crcPhoneCallDiscardReasonDisconnect                                  = 0xe095c1a0
	crcPhoneCallDiscardReasonHangup                                      = 0x57adc690
	crcPhoneCallDiscardReasonBusy                                        = 0xfaf7e8c9
	crcDataJSON                                                          = 0x7d748d04
	crcLabeledPrice                                                      = 0xcb296bf8
	crcInvoice                                                           = 0xc30aa358
	crcPaymentCharge                                                     = 0xea02c27e
	crcPostAddress                                                       = 0x1e8caaeb
	crcPaymentRequestedInfo                                              = 0x909c3f94
	crcPaymentSavedCredentialsCard                                       = 0xcdc27a1f
	crcWebDocument                                                       = 0x1c570ed1
	crcWebDocumentNoProxy                                                = 0xf9c8bcc6
	crcInputWebDocument                                                  = 0x9bed434d
	crcInputWebFileLocation                                              = 0xc239d686
	crcInputWebFileGeoPointLocation                                      = 0x9f2221c9
	crcUploadWebFile                                                     = 0x21e753bc
	crcPaymentsPaymentForm                                               = 0x3f56aea3
	crcPaymentsValidatedRequestedInfo                                    = 0xd1451883
	crcPaymentsPaymentResult                                             = 0x4e5f810d
	crcPaymentsPaymentVerificationNeeded                                 = 0xd8411139
	crcPaymentsPaymentReceipt                                            = 0x500911e1
	crcPaymentsSavedInfo                                                 = 0xfb8fe43c
	crcInputPaymentCredentialsSaved                                      = 0xc10eb2cf
	crcInputPaymentCredentials                                           = 0x3417d728
	crcInputPaymentCredentialsApplePay                                   = 0x0aa1c39f
	crcInputPaymentCredentialsAndroidPay                                 = 0xca05d50e
	crcAccountTmpPassword                                                = 0xdb64fd34
	crcShippingOption                                                    = 0xb6213cdf
	crcInputStickerSetItem                                               = 0xffa0a496
	crcInputPhoneCall                                                    = 0x1e36fded
	crcPhoneCallEmpty                                                    = 0x5366c915
	crcPhoneCallWaiting                                                  = 0x1b8f4ad1
	crcPhoneCallRequested                                                = 0x87eabb53
	crcPhoneCallAccepted                                                 = 0x997c454a
	crcPhoneCall                                                         = 0x8742ae7f
	crcPhoneCallDiscarded                                                = 0x50ca4de1
	crcPhoneConnection                                                   = 0x9d4c17c0
	crcPhoneCallProtocol                                                 = 0xfc878fc8
	crcPhonePhoneCall                                                    = 0xec82e140
	crcUploadCdnFileReuploadNeeded                                       = 0xeea8e46e
	crcUploadCdnFile                                                     = 0xa99fca4f
	crcCdnPublicKey                                                      = 0xc982eaba
	crcCdnConfig                                                         = 0x5725e40a
	crcLangPackString                                                    = 0xcad181f6
	crcLangPackStringPluralized                                          = 0x6c47ac9f
	crcLangPackStringDeleted                                             = 0x2979eeb2
	crcLangPackDifference                                                = 0xf385c1f6
	crcLangPackLanguage                                                  = 0xeeca5ce3
	crcChannelAdminLogEventActionChangeTitle                             = 0xe6dfb825
	crcChannelAdminLogEventActionChangeAbout                             = 0x55188a2e
	crcChannelAdminLogEventActionChangeUsername                          = 0x6a4afc38
	crcChannelAdminLogEventActionChangePhoto                             = 0x434bd2af
	crcChannelAdminLogEventActionToggleInvites                           = 0x1b7907ae
	crcChannelAdminLogEventActionToggleSignatures                        = 0x26ae0971
	crcChannelAdminLogEventActionUpdatePinned                            = 0xe9e82c18
	crcChannelAdminLogEventActionEditMessage                             = 0x709b2405
	crcChannelAdminLogEventActionDeleteMessage                           = 0x42e047bb
	crcChannelAdminLogEventActionParticipantJoin                         = 0x183040d3
	crcChannelAdminLogEventActionParticipantLeave                        = 0xf89777f2
	crcChannelAdminLogEventActionParticipantInvite                       = 0xe31c34d8
	crcChannelAdminLogEventActionParticipantToggleBan                    = 0xe6d83d7e
	crcChannelAdminLogEventActionParticipantToggleAdmin                  = 0xd5676710
	crcChannelAdminLogEventActionChangeStickerSet                        = 0xb1c3caa7
	crcChannelAdminLogEventActionTogglePreHistoryHidden                  = 0x5f5c95f1
	crcChannelAdminLogEventActionDefaultBannedRights                     = 0x2df5fc0a
	crcChannelAdminLogEventActionStopPoll                                = 0x8f079643
	crcChannelAdminLogEventActionChangeLinkedChat                        = 0xa26f881b
	crcChannelAdminLogEventActionChangeLocation                          = 0x0e6b76ae
	crcChannelAdminLogEventActionToggleSlowMode                          = 0x53909779
	crcChannelAdminLogEvent                                              = 0x3b5a3e40
	crcChannelsAdminLogResults                                           = 0xed8af74d
	crcChannelAdminLogEventsFilter                                       = 0xea107ae4
	crcPopularContact                                                    = 0x5ce14175
	crcMessagesFavedStickersNotModified                                  = 0x9e8fa6d3
	crcMessagesFavedStickers                                             = 0xf37f2f16
	crcRecentMeUrlUnknown                                                = 0x46e1d13d
	crcRecentMeUrlUser                                                   = 0x8dbc3336
	crcRecentMeUrlChat                                                   = 0xa01b22f9
	crcRecentMeUrlChatInvite                                             = 0xeb49081d
	crcRecentMeUrlStickerSet                                             = 0xbc0a57dc
	crcHelpRecentMeUrls                                                  = 0x0e0310d7
	crcInputSingleMedia                                                  = 0x1cc6e91f
	crcWebAuthorization                                                  = 0xcac943f2
	crcAccountWebAuthorizations                                          = 0xed56c9fc
	crcInputMessageID                                                    = 0xa676a322
	crcInputMessageReplyTo                                               = 0xbad88395
	crcInputMessagePinned                                                = 0x86872538
	crcInputDialogPeer                                                   = 0xfcaafeb7
	crcInputDialogPeerFolder                                             = 0x64600527
	crcDialogPeer                                                        = 0xe56dbf05
	crcDialogPeerFolder                                                  = 0x514519e2
	crcMessagesFoundStickerSetsNotModified                               = 0x0d54b65d
	crcMessagesFoundStickerSets                                          = 0x5108d648
	crcFileHash                                                          = 0x6242c773
	crcInputClientProxy                                                  = 0x75588b3f
	crcHelpTermsOfServiceUpdateEmpty                                     = 0xe3309f7f
	crcHelpTermsOfServiceUpdate                                          = 0x28ecf961
	crcInputSecureFileUploaded                                           = 0x3334b0f0
	crcInputSecureFile                                                   = 0x5367e5be
	crcSecureFileEmpty                                                   = 0x64199744
	crcSecureFile                                                        = 0xe0277a62
	crcSecureData                                                        = 0x8aeabec3
	crcSecurePlainPhone                                                  = 0x7d6099dd
	crcSecurePlainEmail                                                  = 0x21ec5a5f
	crcSecureValueTypePersonalDetails                                    = 0x9d2a81e3
	crcSecureValueTypePassport                                           = 0x3dac6a00
	crcSecureValueTypeDriverLicense                                      = 0x06e425c4
	crcSecureValueTypeIdentityCard                                       = 0xa0d0744b
	crcSecureValueTypeInternalPassport                                   = 0x99a48f23
	crcSecureValueTypeAddress                                            = 0xcbe31e26
	crcSecureValueTypeUtilityBill                                        = 0xfc36954e
	crcSecureValueTypeBankStatement                                      = 0x89137c0d
	crcSecureValueTypeRentalAgreement                                    = 0x8b883488
	crcSecureValueTypePassportRegistration                               = 0x99e3806a
	crcSecureValueTypeTemporaryRegistration                              = 0xea02ec33
	crcSecureValueTypePhone                                              = 0xb320aadb
	crcSecureValueTypeEmail                                              = 0x8e3ca7ee
	crcSecureValue                                                       = 0x187fa0ca
	crcInputSecureValue                                                  = 0xdb21d0a7
	crcSecureValueHash                                                   = 0xed1ecdb0
	crcSecureValueErrorData                                              = 0xe8a40bd9
	crcSecureValueErrorFrontSide                                         = 0x00be3dfa
	crcSecureValueErrorReverseSide                                       = 0x868a2aa5
	crcSecureValueErrorSelfie                                            = 0xe537ced6
	crcSecureValueErrorFile                                              = 0x7a700873
	crcSecureValueErrorFiles                                             = 0x666220e9
	crcSecureValueError                                                  = 0x869d758f
	crcSecureValueErrorTranslationFile                                   = 0xa1144770
	crcSecureValueErrorTranslationFiles                                  = 0x34636dd8
	crcSecureCredentialsEncrypted                                        = 0x33f0ea47
	crcAccountAuthorizationForm                                          = 0xad2e1cd8
	crcAccountSentEmailCode                                              = 0x811f854f
	crcHelpDeepLinkInfoEmpty                                             = 0x66afa166
	crcHelpDeepLinkInfo                                                  = 0x6a4ee832
	crcSavedPhoneContact                                                 = 0x1142bd56
	crcAccountTakeout                                                    = 0x4dba4501
	crcPasswordKdfAlgoUnknown                                            = 0xd45ab096
	crcPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow = 0x3a912d4a
	crcSecurePasswordKdfAlgoUnknown                                      = 0x004a8537
	crcSecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000                   = 0xbbf2dda0
	crcSecurePasswordKdfAlgoSHA512                                       = 0x86471d92
	crcSecureSecretSettings                                              = 0x1527bcac
	crcInputCheckPasswordEmpty                                           = 0x9880f658
	crcInputCheckPasswordSRP                                             = 0xd27ff082
	crcSecureRequiredType                                                = 0x829d99da
	crcSecureRequiredTypeOneOf                                           = 0x027477b4
	crcHelpPassportConfigNotModified                                     = 0xbfb9f457
	crcHelpPassportConfig                                                = 0xa098d6af
	crcInputAppEvent                                                     = 0x1d1b1245
	crcJsonObjectValue                                                   = 0xc0de1bd9
	crcJsonNull                                                          = 0x3f6d7b68
	crcJsonBool                                                          = 0xc7345e6a
	crcJsonNumber                                                        = 0x2be0dfa4
	crcJsonString                                                        = 0xb71e767a
	crcJsonArray                                                         = 0xf7444763
	crcJsonObject                                                        = 0x99c1d49d
	crcPageTableCell                                                     = 0x34566b6a
	crcPageTableRow                                                      = 0xe0c0c5e5
	crcPageCaption                                                       = 0x6f747657
	crcPageListItemText                                                  = 0xb92fb6cd
	crcPageListItemBlocks                                                = 0x25e073fc
	crcPageListOrderedItemText                                           = 0x5e068047
	crcPageListOrderedItemBlocks                                         = 0x98dd8936
	crcPageRelatedArticle                                                = 0xb390dc08
	crcPage                                                              = 0x98657f0d
	crcHelpSupportName                                                   = 0x8c05f1c9
	crcHelpUserInfoEmpty                                                 = 0xf3ae2eed
	crcHelpUserInfo                                                      = 0x01eb3758
	crcPollAnswer                                                        = 0x6ca9c2e9
	crcPoll                                                              = 0x86e18161
	crcPollAnswerVoters                                                  = 0x3b6ddad2
	crcPollResults                                                       = 0xbadcc1a3
	crcChatOnlines                                                       = 0xf041e250
	crcStatsURL                                                          = 0x47a971e0
	crcChatAdminRights                                                   = 0x5fb224d5
	crcChatBannedRights                                                  = 0x9f120418
	crcInputWallPaper                                                    = 0xe630b979
	crcInputWallPaperSlug                                                = 0x72091c80
	crcInputWallPaperNoFile                                              = 0x8427bbac
	crcAccountWallPapersNotModified                                      = 0x1c199183
	crcAccountWallPapers                                                 = 0x702b65a9
	crcCodeSettings                                                      = 0xdebebe83
	crcWallPaperSettings                                                 = 0x05086cf8
	crcAutoDownloadSettings                                              = 0xe04232f3
	crcAccountAutoDownloadSettings                                       = 0x63cacf26
	crcEmojiKeyword                                                      = 0xd5b3b9f9
	crcEmojiKeywordDeleted                                               = 0x236df622
	crcEmojiKeywordsDifference                                           = 0x5cc761bd
	crcEmojiURL                                                          = 0xa575739d
	crcEmojiLanguage                                                     = 0xb3fb5361
	crcFileLocationToBeDeprecated                                        = 0xbc7fc6cd
	crcFolder                                                            = 0xff544e65
	crcInputFolderPeer                                                   = 0xfbd2c296
	crcFolderPeer                                                        = 0xe9baa668
	crcMessagesSearchCounter                                             = 0xe844ebff
	crcUrlAuthResultRequest                                              = 0x92d33a0e
	crcUrlAuthResultAccepted                                             = 0x8f8c0e4e
	crcUrlAuthResultDefault                                              = 0xa9d6db1f
	crcChannelLocationEmpty                                              = 0xbfb5ad8b
	crcChannelLocation                                                   = 0x209b82db
	crcPeerLocated                                                       = 0xca461b5d
	crcPeerSelfLocated                                                   = 0xf8ec284b
	crcRestrictionReason                                                 = 0xd072acb4
	crcInputTheme                                                        = 0x3c5693e9
	crcInputThemeSlug                                                    = 0xf5890df1
	crcTheme                                                             = 0x028f1114
	crcAccountThemesNotModified                                          = 0xf41eb622
	crcAccountThemes                                                     = 0x7f676421
	crcAuthLoginToken                                                    = 0x629f1980
	crcAuthLoginTokenMigrateTo                                           = 0x068e9916
	crcAuthLoginTokenSuccess                                             = 0x390d5c5e
	crcAccountContentSettings                                            = 0x57e28221
	crcMessagesInactiveChats                                             = 0xa927fec5
	crcBaseThemeClassic                                                  = 0xc3a12462
	crcBaseThemeDay                                                      = 0xfbd81688
	crcBaseThemeNight                                                    = 0xb7b31ea8
	crcBaseThemeTinted                                                   = 0x6d5f77ee
	crcBaseThemeArctic                                                   = 0x5b11125a
	crcInputThemeSettings                                                = 0xbd507cd1
	crcThemeSettings                                                     = 0x9c14984a
	crcWebPageAttributeTheme                                             = 0x54b56617
	crcMessageUserVote                                                   = 0xa28e5559
	crcMessageUserVoteInputOption                                        = 0x36377430
	crcMessageUserVoteMultiple                                           = 0x0e8fe0de
	crcMessagesVotesList                                                 = 0x0823f649
	crcBankCardOpenUrl                                                   = 0xf568028a
	crcPaymentsBankCardData                                              = 0x3e24e573
	crcDialogFilter                                                      = 0x7438f7e8
	crcDialogFilterSuggested                                             = 0x77744d4a
	crcStatsDateRangeDays                                                = 0xb637edaf
	crcStatsAbsValueAndPrev                                              = 0xcb43acde
	crcStatsPercentValue                                                 = 0xcbce2fe0
	crcStatsGraphAsync                                                   = 0x4a27eb2d
	crcStatsGraphError                                                   = 0xbedc9822
	crcStatsGraph                                                        = 0x8ea464b6
	crcMessageInteractionCounters                                        = 0xad4fc9bd
	crcStatsBroadcastStats                                               = 0xbdf78394
	crcHelpPromoDataEmpty                                                = 0x98f6ac75
	crcHelpPromoData                                                     = 0x8c39793f
	crcVideoSize                                                         = 0xe831c556
	crcStatsGroupTopPoster                                               = 0x18f3d0f7
	crcStatsGroupTopAdmin                                                = 0x6014f412
	crcStatsGroupTopInviter                                              = 0x31962a4c
	crcStatsMegagroupStats                                               = 0xef7ff916
	crcGlobalPrivacySettings                                             = 0xbea2f424
	crcInvokeAfterMsg                                                    = 0xcb9f372d
	crcInvokeAfterMsgs                                                   = 0x3dc4b4f0
	crcInitConnection                                                    = 0xc1cd5ea9
	crcInvokeWithLayer                                                   = 0xda9b0d0d
	crcInvokeWithoutUpdates                                              = 0xbf9459b7
	crcInvokeWithMessagesRange                                           = 0x365275f2
	crcInvokeWithTakeout                                                 = 0xaca9fd2e
	crcAuthSendCode                                                      = 0xa677244f
	crcAuthSignUp                                                        = 0x80eee427
	crcAuthSignIn                                                        = 0xbcd51581
	crcAuthLogOut                                                        = 0x5717da40
	crcAuthResetAuthorizations                                           = 0x9fab0d1a
	crcAuthExportAuthorization                                           = 0xe5bfffcd
	crcAuthImportAuthorization                                           = 0xe3ef9613
	crcAuthBindTempAuthKey                                               = 0xcdd42a05
	crcAuthImportBotAuthorization                                        = 0x67a3ff2c
	crcAuthCheckPassword                                                 = 0xd18b4d16
	crcAuthRequestPasswordRecovery                                       = 0xd897bc66
	crcAuthRecoverPassword                                               = 0x4ea56e92
	crcAuthResendCode                                                    = 0x3ef1a9bf
	crcAuthCancelCode                                                    = 0x1f040578
	crcAuthDropTempAuthKeys                                              = 0x8e48a188
	crcAuthExportLoginToken                                              = 0xb1b41517
	crcAuthImportLoginToken                                              = 0x95ac5ce4
	crcAuthAcceptLoginToken                                              = 0xe894ad4d
	crcAccountRegisterDevice                                             = 0x68976c6f
	crcAccountUnregisterDevice                                           = 0x3076c4bf
	crcAccountUpdateNotifySettings                                       = 0x84be5b93
	crcAccountGetNotifySettings                                          = 0x12b3ad31
	crcAccountResetNotifySettings                                        = 0xdb7e1747
	crcAccountUpdateProfile                                              = 0x78515775
	crcAccountUpdateStatus                                               = 0x6628562c
	crcAccountGetWallPapers                                              = 0xaabb1763
	crcAccountReportPeer                                                 = 0xae189d5f
	crcAccountCheckUsername                                              = 0x2714d86c
	crcAccountUpdateUsername                                             = 0x3e0bdd7c
	crcAccountGetPrivacy                                                 = 0xdadbc950
	crcAccountSetPrivacy                                                 = 0xc9f81ce8
	crcAccountDeleteAccount                                              = 0x418d4e0b
	crcAccountGetAccountTTL                                              = 0x08fc711d
	crcAccountSetAccountTTL                                              = 0x2442485e
	crcAccountSendChangePhoneCode                                        = 0x82574ae5
	crcAccountChangePhone                                                = 0x70c32edb
	crcAccountUpdateDeviceLocked                                         = 0x38df3532
	crcAccountGetAuthorizations                                          = 0xe320c158
	crcAccountResetAuthorization                                         = 0xdf77f3bc
	crcAccountGetPassword                                                = 0x548a30f5
	crcAccountGetPasswordSettings                                        = 0x9cd4eaf9
	crcAccountUpdatePasswordSettings                                     = 0xa59b102f
	crcAccountSendConfirmPhoneCode                                       = 0x1b3faa88
	crcAccountConfirmPhone                                               = 0x5f2178c3
	crcAccountGetTmpPassword                                             = 0x449e0b51
	crcAccountGetWebAuthorizations                                       = 0x182e6d6f
	crcAccountResetWebAuthorization                                      = 0x2d01b9ef
	crcAccountResetWebAuthorizations                                     = 0x682d2594
	crcAccountGetAllSecureValues                                         = 0xb288bc7d
	crcAccountGetSecureValue                                             = 0x73665bc2
	crcAccountSaveSecureValue                                            = 0x899fe31d
	crcAccountDeleteSecureValue                                          = 0xb880bc4b
	crcAccountGetAuthorizationForm                                       = 0xb86ba8e1
	crcAccountAcceptAuthorization                                        = 0xe7027c94
	crcAccountSendVerifyPhoneCode                                        = 0xa5a356f9
	crcAccountVerifyPhone                                                = 0x4dd3a7f6
	crcAccountSendVerifyEmailCode                                        = 0x7011509f
	crcAccountVerifyEmail                                                = 0xecba39db
	crcAccountInitTakeoutSession                                         = 0xf05b4804
	crcAccountFinishTakeoutSession                                       = 0x1d2652ee
	crcAccountConfirmPasswordEmail                                       = 0x8fdf1920
	crcAccountResendPasswordEmail                                        = 0x7a7f2a15
	crcAccountCancelPasswordEmail                                        = 0xc1cbd5b6
	crcAccountGetContactSignUpNotification                               = 0x9f07c728
	crcAccountSetContactSignUpNotification                               = 0xcff43f61
	crcAccountGetNotifyExceptions                                        = 0x53577479
	crcAccountGetWallPaper                                               = 0xfc8ddbea
	crcAccountUploadWallPaper                                            = 0xdd853661
	crcAccountSaveWallPaper                                              = 0x6c5a5b37
	crcAccountInstallWallPaper                                           = 0xfeed5769
	crcAccountResetWallPapers                                            = 0xbb3b9804
	crcAccountGetAutoDownloadSettings                                    = 0x56da0b3f
	crcAccountSaveAutoDownloadSettings                                   = 0x76f36233
	crcAccountUploadTheme                                                = 0x1c3db333
	crcAccountCreateTheme                                                = 0x8432c21f
	crcAccountUpdateTheme                                                = 0x5cb367d5
	crcAccountSaveTheme                                                  = 0xf257106c
	crcAccountInstallTheme                                               = 0x7ae43737
	crcAccountGetTheme                                                   = 0x8d9d742b
	crcAccountGetThemes                                                  = 0x285946f8
	crcAccountSetContentSettings                                         = 0xb574b16b
	crcAccountGetContentSettings                                         = 0x8b9b4dae
	crcAccountGetMultiWallPapers                                         = 0x65ad71dc
	crcAccountGetGlobalPrivacySettings                                   = 0xeb2b4cf6
	crcAccountSetGlobalPrivacySettings                                   = 0x1edaaac2
	crcUsersGetUsers                                                     = 0x0d91a548
	crcUsersGetFullUser                                                  = 0xca30a5b1
	crcUsersSetSecureValueErrors                                         = 0x90c894b5
	crcContactsGetContactIDs                                             = 0x2caa4a42
	crcContactsGetStatuses                                               = 0xc4a353ee
	crcContactsGetContacts                                               = 0xc023849f
	crcContactsImportContacts                                            = 0x2c800be5
	crcContactsDeleteContacts                                            = 0x096a0e00
	crcContactsDeleteByPhones                                            = 0x1013fd9e
	crcContactsBlock                                                     = 0x332b49fc
	crcContactsUnblock                                                   = 0xe54100bd
	crcContactsGetBlocked                                                = 0xf57c350f
	crcContactsSearch                                                    = 0x11f812d8
	crcContactsResolveUsername                                           = 0xf93ccba3
	crcContactsGetTopPeers                                               = 0xd4982db5
	crcContactsResetTopPeerRating                                        = 0x1ae373ac
	crcContactsResetSaved                                                = 0x879537f1
	crcContactsGetSaved                                                  = 0x82f1e39f
	crcContactsToggleTopPeers                                            = 0x8514bdda
	crcContactsAddContact                                                = 0xe8f463d0
	crcContactsAcceptContact                                             = 0xf831a20f
	crcContactsGetLocated                                                = 0xd348bc44
	crcMessagesGetMessages                                               = 0x63c66506
	crcMessagesGetDialogs                                                = 0xa0ee3b73
	crcMessagesGetHistory                                                = 0xdcbb8260
	crcMessagesSearch                                                    = 0x8614ef68
	crcMessagesReadHistory                                               = 0x0e306d3a
	crcMessagesDeleteHistory                                             = 0x1c015b09
	crcMessagesDeleteMessages                                            = 0xe58e95d2
	crcMessagesReceivedMessages                                          = 0x05a954c0
	crcMessagesSetTyping                                                 = 0xa3825e50
	crcMessagesSendMessage                                               = 0x520c3870
	crcMessagesSendMedia                                                 = 0x3491eba9
	crcMessagesForwardMessages                                           = 0xd9fee60e
	crcMessagesReportSpam                                                = 0xcf1592db
	crcMessagesGetPeerSettings                                           = 0x3672e09c
	crcMessagesReport                                                    = 0xbd82b658
	crcMessagesGetChats                                                  = 0x3c6aa187
	crcMessagesGetFullChat                                               = 0x3b831c66
	crcMessagesEditChatTitle                                             = 0xdc452855
	crcMessagesEditChatPhoto                                             = 0xca4c79d8
	crcMessagesAddChatUser                                               = 0xf9a0aa09
	crcMessagesDeleteChatUser                                            = 0xe0611f16
	crcMessagesCreateChat                                                = 0x09cb126e
	crcMessagesGetDhConfig                                               = 0x26cf8950
	crcMessagesRequestEncryption                                         = 0xf64daf43
	crcMessagesAcceptEncryption                                          = 0x3dbc0415
	crcMessagesDiscardEncryption                                         = 0xedd923c5
	crcMessagesSetEncryptedTyping                                        = 0x791451ed
	crcMessagesReadEncryptedHistory                                      = 0x7f4b690a
	crcMessagesSendEncrypted                                             = 0xa9776773
	crcMessagesSendEncryptedFile                                         = 0x9a901b66
	crcMessagesSendEncryptedService                                      = 0x32d439a4
	crcMessagesReceivedQueue                                             = 0x55a5bb66
	crcMessagesReportEncryptedSpam                                       = 0x4b0c8c0f
	crcMessagesReadMessageContents                                       = 0x36a73f77
	crcMessagesGetStickers                                               = 0x043d4f2c
	crcMessagesGetAllStickers                                            = 0x1c9618b1
	crcMessagesGetWebPagePreview                                         = 0x8b68b0cc
	crcMessagesExportChatInvite                                          = 0x0df7534c
	crcMessagesCheckChatInvite                                           = 0x3eadb1bb
	crcMessagesImportChatInvite                                          = 0x6c50051c
	crcMessagesGetStickerSet                                             = 0x2619a90e
	crcMessagesInstallStickerSet                                         = 0xc78fe460
	crcMessagesUninstallStickerSet                                       = 0xf96e55de
	crcMessagesStartBot                                                  = 0xe6df7378
	crcMessagesGetMessagesViews                                          = 0xc4c8a55d
	crcMessagesEditChatAdmin                                             = 0xa9e69f2e
	crcMessagesMigrateChat                                               = 0x15a3b8e3
	crcMessagesSearchGlobal                                              = 0xbf7225a4
	crcMessagesReorderStickerSets                                        = 0x78337739
	crcMessagesGetDocumentByHash                                         = 0x338e2464
	crcMessagesGetSavedGifs                                              = 0x83bf3d52
	crcMessagesSaveGif                                                   = 0x327a30cb
	crcMessagesGetInlineBotResults                                       = 0x514e999d
	crcMessagesSetInlineBotResults                                       = 0xeb5ea206
	crcMessagesSendInlineBotResult                                       = 0x220815b0
	crcMessagesGetMessageEditData                                        = 0xfda68d36
	crcMessagesEditMessage                                               = 0x48f71778
	crcMessagesEditInlineBotMessage                                      = 0x83557dba
	crcMessagesGetBotCallbackAnswer                                      = 0x810a9fec
	crcMessagesSetBotCallbackAnswer                                      = 0xd58f130a
	crcMessagesGetPeerDialogs                                            = 0xe470bcfd
	crcMessagesSaveDraft                                                 = 0xbc39e14b
	crcMessagesGetAllDrafts                                              = 0x6a3f8d65
	crcMessagesGetFeaturedStickers                                       = 0x2dacca4f
	crcMessagesReadFeaturedStickers                                      = 0x5b118126
	crcMessagesGetRecentStickers                                         = 0x5ea192c9
	crcMessagesSaveRecentSticker                                         = 0x392718f8
	crcMessagesClearRecentStickers                                       = 0x8999602d
	crcMessagesGetArchivedStickers                                       = 0x57f17692
	crcMessagesGetMaskStickers                                           = 0x65b8c79f
	crcMessagesGetAttachedStickers                                       = 0xcc5b67cc
	crcMessagesSetGameScore                                              = 0x8ef8ecc0
	crcMessagesSetInlineGameScore                                        = 0x15ad9f64
	crcMessagesGetGameHighScores                                         = 0xe822649d
	crcMessagesGetInlineGameHighScores                                   = 0x0f635e1b
	crcMessagesGetCommonChats                                            = 0x0d0a48c4
	crcMessagesGetAllChats                                               = 0xeba80ff0
	crcMessagesGetWebPage                                                = 0x32ca8f91
	crcMessagesToggleDialogPin                                           = 0xa731e257
	crcMessagesReorderPinnedDialogs                                      = 0x3b1adf37
	crcMessagesGetPinnedDialogs                                          = 0xd6b94df2
	crcMessagesSetBotShippingResults                                     = 0xe5f672fa
	crcMessagesSetBotPrecheckoutResults                                  = 0x09c2dd95
	crcMessagesUploadMedia                                               = 0x519bc2b1
	crcMessagesSendScreenshotNotification                                = 0xc97df020
	crcMessagesGetFavedStickers                                          = 0x21ce0b0e
	crcMessagesFaveSticker                                               = 0xb9ffc55b
	crcMessagesGetUnreadMentions                                         = 0x46578472
	crcMessagesReadMentions                                              = 0x0f0189d3
	crcMessagesGetRecentLocations                                        = 0xbbc45b09
	crcMessagesSendMultiMedia                                            = 0xcc0110cb
	crcMessagesUploadEncryptedFile                                       = 0x5057c497
	crcMessagesSearchStickerSets                                         = 0xc2b7d08b
	crcMessagesGetSplitRanges                                            = 0x1cff7e08
	crcMessagesMarkDialogUnread                                          = 0xc286d98f
	crcMessagesGetDialogUnreadMarks                                      = 0x22e24e22
	crcMessagesClearAllDrafts                                            = 0x7e58ee9c
	crcMessagesUpdatePinnedMessage                                       = 0xd2aaf7ec
	crcMessagesSendVote                                                  = 0x10ea6184
	crcMessagesGetPollResults                                            = 0x73bb643b
	crcMessagesGetOnlines                                                = 0x6e2be050
	crcMessagesGetStatsURL                                               = 0x812c2ae6
	crcMessagesEditChatAbout                                             = 0xdef60797
	crcMessagesEditChatDefaultBannedRights                               = 0xa5866b41
	crcMessagesGetEmojiKeywords                                          = 0x35a0e062
	crcMessagesGetEmojiKeywordsDifference                                = 0x1508b6af
	crcMessagesGetEmojiKeywordsLanguages                                 = 0x4e9963b2
	crcMessagesGetEmojiURL                                               = 0xd5b10c26
	crcMessagesGetSearchCounters                                         = 0x732eef00
	crcMessagesRequestUrlAuth                                            = 0xe33f5613
	crcMessagesAcceptUrlAuth                                             = 0xf729ea98
	crcMessagesHidePeerSettingsBar                                       = 0x4facb138
	crcMessagesGetScheduledHistory                                       = 0xe2c2685b
	crcMessagesGetScheduledMessages                                      = 0xbdbb0464
	crcMessagesSendScheduledMessages                                     = 0xbd38850a
	crcMessagesDeleteScheduledMessages                                   = 0x59ae2b16
	crcMessagesGetPollVotes                                              = 0xb86e380e
	crcMessagesToggleStickerSets                                         = 0xb5052fea
	crcMessagesGetDialogFilters                                          = 0xf19ed96d
	crcMessagesGetSuggestedDialogFilters                                 = 0xa29cd42c
	crcMessagesUpdateDialogFilter                                        = 0x1ad4a04a
	crcMessagesUpdateDialogFiltersOrder                                  = 0xc563c1e4
	crcMessagesGetOldFeaturedStickers                                    = 0x5fe7025b
	crcUpdatesGetState                                                   = 0xedd4882a
	crcUpdatesGetDifference                                              = 0x25939651
	crcUpdatesGetChannelDifference                                       = 0x03173d78
	crcPhotosUpdateProfilePhoto                                          = 0x72d4742c
	crcPhotosUploadProfilePhoto                                          = 0x89f30f69
	crcPhotosDeletePhotos                                                = 0x87cf7f2f
	crcPhotosGetUserPhotos                                               = 0x91cd32a8
	crcUploadSaveFilePart                                                = 0xb304a621
	crcUploadGetFile                                                     = 0xb15a9afc
	crcUploadSaveBigFilePart                                             = 0xde7b673d
	crcUploadGetWebFile                                                  = 0x24e6818d
	crcUploadGetCdnFile                                                  = 0x2000bcc3
	crcUploadReuploadCdnFile                                             = 0x9b2754a8
	crcUploadGetCdnFileHashes                                            = 0x4da54231
	crcUploadGetFileHashes                                               = 0xc7025931
	crcHelpGetConfig                                                     = 0xc4f9186b
	crcHelpGetNearestDc                                                  = 0x1fb33026
	crcHelpGetAppUpdate                                                  = 0x522d5a7d
	crcHelpGetInviteText                                                 = 0x4d392343
	crcHelpGetSupport                                                    = 0x9cdf08cd
	crcHelpGetAppChangelog                                               = 0x9010ef6f
	crcHelpSetBotUpdatesStatus                                           = 0xec22cfcd
	crcHelpGetCdnConfig                                                  = 0x52029342
	crcHelpGetRecentMeUrls                                               = 0x3dc0f114
	crcHelpGetTermsOfServiceUpdate                                       = 0x2ca51fd1
	crcHelpAcceptTermsOfService                                          = 0xee72f79a
	crcHelpGetDeepLinkInfo                                               = 0x3fedc75f
	crcHelpGetAppConfig                                                  = 0x98914110
	crcHelpSaveAppLog                                                    = 0x6f02f748
	crcHelpGetPassportConfig                                             = 0xc661ad08
	crcHelpGetSupportName                                                = 0xd360e72c
	crcHelpGetUserInfo                                                   = 0x038a08d3
	crcHelpEditUserInfo                                                  = 0x66b91b70
	crcHelpGetPromoData                                                  = 0xc0977421
	crcHelpHidePromoData                                                 = 0x1e251c95
	crcHelpDismissSuggestion                                             = 0x077fa99f
	crcChannelsReadHistory                                               = 0xcc104937
	crcChannelsDeleteMessages                                            = 0x84c1fd4e
	crcChannelsDeleteUserHistory                                         = 0xd10dd71b
	crcChannelsReportSpam                                                = 0xfe087810
	crcChannelsGetMessages                                               = 0xad8c9a23
	crcChannelsGetParticipants                                           = 0x123e05e9
	crcChannelsGetParticipant                                            = 0x546dd7a6
	crcChannelsGetChannels                                               = 0x0a7f6bbb
	crcChannelsGetFullChannel                                            = 0x08736a09
	crcChannelsCreateChannel                                             = 0x3d5fb10f
	crcChannelsEditAdmin                                                 = 0xd33c8902
	crcChannelsEditTitle                                                 = 0x566decd0
	crcChannelsEditPhoto                                                 = 0xf12e57c9
	crcChannelsCheckUsername                                             = 0x10e6bd2c
	crcChannelsUpdateUsername                                            = 0x3514b3de
	crcChannelsJoinChannel                                               = 0x24b524c5
	crcChannelsLeaveChannel                                              = 0xf836aa95
	crcChannelsInviteToChannel                                           = 0x199f3a6c
	crcChannelsDeleteChannel                                             = 0xc0111fe3
	crcChannelsExportMessageLink                                         = 0xceb77163
	crcChannelsToggleSignatures                                          = 0x1f69b606
	crcChannelsGetAdminedPublicChannels                                  = 0xf8b036af
	crcChannelsEditBanned                                                = 0x72796912
	crcChannelsGetAdminLog                                               = 0x33ddf480
	crcChannelsSetStickers                                               = 0xea8ca4f9
	crcChannelsReadMessageContents                                       = 0xeab5dc38
	crcChannelsDeleteHistory                                             = 0xaf369d42
	crcChannelsTogglePreHistoryHidden                                    = 0xeabbb94c
	crcChannelsGetLeftChannels                                           = 0x8341ecc0
	crcChannelsGetGroupsForDiscussion                                    = 0xf5dad378
	crcChannelsSetDiscussionGroup                                        = 0x40582bb2
	crcChannelsEditCreator                                               = 0x8f38cd1f
	crcChannelsEditLocation                                              = 0x58e63f6d
	crcChannelsToggleSlowMode                                            = 0xedd49ef0
	crcChannelsGetInactiveChannels                                       = 0x11e831ee
	crcBotsSendCustomRequest                                             = 0xaa2769ed
	crcBotsAnswerWebhookJSONQuery                                        = 0xe6213f4d
	crcBotsSetBotCommands                                                = 0x805d46f6
	crcPaymentsGetPaymentForm                                            = 0x99f09745
	crcPaymentsGetPaymentReceipt                                         = 0xa092a980
	crcPaymentsValidateRequestedInfo                                     = 0x770a8e74
	crcPaymentsSendPaymentForm                                           = 0x2b8879b3
	crcPaymentsGetSavedInfo                                              = 0x227d824b
	crcPaymentsClearSavedInfo                                            = 0xd83d70c1
	crcPaymentsGetBankCardData                                           = 0x2e79d779
	crcStickersCreateStickerSet                                          = 0xf1036780
	crcStickersRemoveStickerFromSet                                      = 0xf7760f51
	crcStickersChangeStickerPosition                                     = 0xffb6d4ca
	crcStickersAddStickerToSet                                           = 0x8653febe
	crcStickersSetStickerSetThumb                                        = 0x9a364e30
	crcPhoneGetCallConfig                                                = 0x55451fa9
	crcPhoneRequestCall                                                  = 0x42ff96ed
	crcPhoneAcceptCall                                                   = 0x3bd2b4a0
	crcPhoneConfirmCall                                                  = 0x2efe1722
	crcPhoneReceivedCall                                                 = 0x17d54f61
	crcPhoneDiscardCall                                                  = 0xb2cbc1c0
	crcPhoneSetCallRating                                                = 0x59ead627
	crcPhoneSaveCallDebug                                                = 0x277add7e
	crcPhoneSendSignalingData                                            = 0xff7a9383
	crcLangpackGetLangPack                                               = 0xf2f2330a
	crcLangpackGetStrings                                                = 0xefea3803
	crcLangpackGetDifference                                             = 0xcd984aa5
	crcLangpackGetLanguages                                              = 0x42c6978f
	crcLangpackGetLanguage                                               = 0x6a596502
	crcFoldersEditPeerFolders                                            = 0x6847d0ab
	crcFoldersDeleteFolder                                               = 0x1c295881
	crcStatsGetBroadcastStats                                            = 0xab42441a
	crcStatsLoadAsyncGraph                                               = 0x621d5fa0
	crcStatsGetMegagroupStats                                            = 0xdcdf8607
)

type ResPQ struct {
	Nonce                       []byte
	ServerNonce                 []byte
	Pq                          string
	ServerPublicKeyFingerprints []int64
}

type PQInnerDataDc struct {
	Pq          string
	P           string
	Q           string
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
	Dc          int32
}

type PQInnerDataTempDc struct {
	Pq          string
	P           string
	Q           string
	Nonce       []byte
	ServerNonce []byte
	NewNonce    []byte
	Dc          int32
	ExpiresIn   int32
}

type ServerDHParamsFail struct {
	Nonce        []byte
	ServerNonce  []byte
	NewNonceHash []byte
}

type ServerDHParamsOk struct {
	Nonce           []byte
	ServerNonce     []byte
	EncryptedAnswer string
}

type ServerDHInnerData struct {
	Nonce       []byte
	ServerNonce []byte
	G           int32
	DhPrime     string
	GA          string
	ServerTime  int32
}

type ClientDHInnerData struct {
	Nonce       []byte
	ServerNonce []byte
	RetryID     int64
	GB          string
}

type DhGenOk struct {
	Nonce         []byte
	ServerNonce   []byte
	NewNonceHash1 []byte
}

type DhGenRetry struct {
	Nonce         []byte
	ServerNonce   []byte
	NewNonceHash2 []byte
}

type DhGenFail struct {
	Nonce         []byte
	ServerNonce   []byte
	NewNonceHash3 []byte
}

type BindAuthKeyInner struct {
	Nonce         int64
	TempAuthKeyID int64
	PermAuthKeyID int64
	TempSessionID int64
	ExpiresAt     int32
}

type RpcError struct {
	ErrorCode    int32
	ErrorMessage string
}

type RpcAnswerUnknown struct {
}

type RpcAnswerDroppedRunning struct {
}

type RpcAnswerDropped struct {
	MsgID int64
	SeqNo int32
	Bytes int32
}

type FutureSalt struct {
	ValidSince int32
	ValidUntil int32
	Salt       int64
}

type FutureSalts struct {
	ReqMsgID int64
	Now      int32
	Salts    []TL // Future_salt
}

type Pong struct {
	MsgID  int64
	PingID int64
}

type DestroySessionOk struct {
	SessionID int64
}

type DestroySessionNone struct {
	SessionID int64
}

type NewSessionCreated struct {
	FirstMsgID int64
	UniqueID   int64
	ServerSalt int64
}

type MsgsAck struct {
	MsgIds []int64
}

type BadMsgNotification struct {
	BadMsgID    int64
	BadMsgSeqno int32
	ErrorCode   int32
}

type BadServerSalt struct {
	BadMsgID      int64
	BadMsgSeqno   int32
	ErrorCode     int32
	NewServerSalt int64
}

type MsgResendReq struct {
	MsgIds []int64
}

type MsgsStateReq struct {
	MsgIds []int64
}

type MsgsStateInfo struct {
	ReqMsgID int64
	Info     string
}

type MsgsAllInfo struct {
	MsgIds []int64
	Info   string
}

type MsgDetailedInfo struct {
	MsgID       int64
	AnswerMsgID int64
	Bytes       int32
	Status      int32
}

type MsgNewDetailedInfo struct {
	AnswerMsgID int64
	Bytes       int32
	Status      int32
}

type DestroyAuthKeyOk struct {
}

type DestroyAuthKeyNone struct {
}

type DestroyAuthKeyFail struct {
}

type ReqPqMulti struct {
	Nonce []byte
}

type ReqDHParams struct {
	Nonce                []byte
	ServerNonce          []byte
	P                    string
	Q                    string
	PublicKeyFingerprint int64
	EncryptedData        string
}

type SetClientDHParams struct {
	Nonce         []byte
	ServerNonce   []byte
	EncryptedData string
}

type RpcDropAnswer struct {
	ReqMsgID int64
}

type GetFutureSalts struct {
	Num int32
}

type Ping struct {
	PingID int64
}

type PingDelayDisconnect struct {
	PingID          int64
	DisconnectDelay int32
}

type DestroySession struct {
	SessionID int64
}

type HttpWait struct {
	MaxDelay  int32
	WaitAfter int32
	MaxWait   int32
}

type DestroyAuthKey struct {
}

type True struct {
}

type BoolFalse struct {
}

type BoolTrue struct {
}

type Error struct {
	Code int32
	Text string
}

type IpPort struct {
	Ipv4 int32
	Port int32
}

type InputPeerEmpty struct {
}

type InputPeerSelf struct {
}

type InputPeerChat struct {
	ChatID int32
}

type InputPeerUser struct {
	UserID     int32
	AccessHash int64
}

type InputPeerChannel struct {
	ChannelID  int32
	AccessHash int64
}

type InputPeerUserFromMessage struct {
	Peer   TL // InputPeer
	MsgID  int32
	UserID int32
}

type InputPeerChannelFromMessage struct {
	Peer      TL // InputPeer
	MsgID     int32
	ChannelID int32
}

type InputUserEmpty struct {
}

type InputUserSelf struct {
}

type InputUser struct {
	UserID     int32
	AccessHash int64
}

type InputUserFromMessage struct {
	Peer   TL // InputPeer
	MsgID  int32
	UserID int32
}

type InputPhoneContact struct {
	ClientID  int64
	Phone     string
	FirstName string
	LastName  string
}

type InputFile struct {
	ID          int64
	Parts       int32
	Name        string
	Md5Checksum string
}

type InputFileBig struct {
	ID    int64
	Parts int32
	Name  string
}

type InputMediaEmpty struct {
}

type InputMediaUploadedPhoto struct {
	Flags      int32
	File       TL    // InputFile
	Stickers   []TL  // InputDocument // flag
	TtlSeconds int32 // flag
}

type InputMediaPhoto struct {
	Flags      int32
	ID         TL    // InputPhoto
	TtlSeconds int32 // flag
}

type InputMediaGeoPoint struct {
	GeoPoint TL // InputGeoPoint
}

type InputMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
}

type InputMediaUploadedDocument struct {
	Flags        int32
	NosoundVideo bool // flag
	ForceFile    bool // flag
	File         TL   // InputFile
	Thumb        TL   // InputFile // flag
	MimeType     string
	Attributes   []TL  // DocumentAttribute
	Stickers     []TL  // InputDocument // flag
	TtlSeconds   int32 // flag
}

type InputMediaDocument struct {
	Flags      int32
	ID         TL    // InputDocument
	TtlSeconds int32 // flag
}

type InputMediaVenue struct {
	GeoPoint  TL // InputGeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

type InputMediaPhotoExternal struct {
	Flags      int32
	Url        string
	TtlSeconds int32 // flag
}

type InputMediaDocumentExternal struct {
	Flags      int32
	Url        string
	TtlSeconds int32 // flag
}

type InputMediaGame struct {
	ID TL // InputGame
}

type InputMediaInvoice struct {
	Flags        int32
	Title        string
	Description  string
	Photo        TL // InputWebDocument // flag
	Invoice      TL // Invoice
	Payload      []byte
	Provider     string
	ProviderData TL // DataJSON
	StartParam   string
}

type InputMediaGeoLive struct {
	Flags    int32
	Stopped  bool  // flag
	GeoPoint TL    // InputGeoPoint
	Period   int32 // flag
}

type InputMediaPoll struct {
	Flags            int32
	Poll             TL     // Poll
	CorrectAnswers   []TL   // Bytes // flag
	Solution         string // flag
	SolutionEntities []TL   // MessageEntity // flag
}

type InputMediaDice struct {
	Emoticon string
}

type InputChatPhotoEmpty struct {
}

type InputChatUploadedPhoto struct {
	Flags        int32
	File         TL      // InputFile // flag
	Video        TL      // InputFile // flag
	VideoStartTs float64 // flag
}

type InputChatPhoto struct {
	ID TL // InputPhoto
}

type InputGeoPointEmpty struct {
}

type InputGeoPoint struct {
	Lat  float64
	Long float64
}

type InputPhotoEmpty struct {
}

type InputPhoto struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

type InputFileLocation struct {
	VolumeID      int64
	LocalID       int32
	Secret        int64
	FileReference []byte
}

type InputEncryptedFileLocation struct {
	ID         int64
	AccessHash int64
}

type InputDocumentFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

type InputSecureFileLocation struct {
	ID         int64
	AccessHash int64
}

type InputTakeoutFileLocation struct {
}

type InputPhotoFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	ThumbSize     string
}

type InputPhotoLegacyFileLocation struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
	VolumeID      int64
	LocalID       int32
	Secret        int64
}

type InputPeerPhotoFileLocation struct {
	Flags    int32
	Big      bool // flag
	Peer     TL   // InputPeer
	VolumeID int64
	LocalID  int32
}

type InputStickerSetThumb struct {
	Stickerset TL // InputStickerSet
	VolumeID   int64
	LocalID    int32
}

type PeerUser struct {
	UserID int32
}

type PeerChat struct {
	ChatID int32
}

type PeerChannel struct {
	ChannelID int32
}

type StorageFileUnknown struct {
}

type StorageFilePartial struct {
}

type StorageFileJpeg struct {
}

type StorageFileGif struct {
}

type StorageFilePng struct {
}

type StorageFilePdf struct {
}

type StorageFileMp3 struct {
}

type StorageFileMov struct {
}

type StorageFileMp4 struct {
}

type StorageFileWebp struct {
}

type UserEmpty struct {
	ID int32
}

type User struct {
	Flags                int32
	Self                 bool // flag
	Contact              bool // flag
	MutualContact        bool // flag
	Deleted              bool // flag
	Bot                  bool // flag
	BotChatHistory       bool // flag
	BotNochats           bool // flag
	Verified             bool // flag
	Restricted           bool // flag
	Min                  bool // flag
	BotInlineGeo         bool // flag
	Support              bool // flag
	Scam                 bool // flag
	ID                   int32
	AccessHash           int64  // flag
	FirstName            string // flag
	LastName             string // flag
	Username             string // flag
	Phone                string // flag
	Photo                TL     // UserProfilePhoto // flag
	Status               TL     // UserStatus // flag
	BotInfoVersion       int32  // flag
	RestrictionReason    []TL   // RestrictionReason // flag
	BotInlinePlaceholder string // flag
	LangCode             string // flag
}

type UserProfilePhotoEmpty struct {
}

type UserProfilePhoto struct {
	Flags      int32
	HasVideo   bool // flag
	PhotoID    int64
	PhotoSmall TL // FileLocation
	PhotoBig   TL // FileLocation
	DcID       int32
}

type UserStatusEmpty struct {
}

type UserStatusOnline struct {
	Expires int32
}

type UserStatusOffline struct {
	WasOnline int32
}

type UserStatusRecently struct {
}

type UserStatusLastWeek struct {
}

type UserStatusLastMonth struct {
}

type ChatEmpty struct {
	ID int32
}

type Chat struct {
	Flags               int32
	Creator             bool // flag
	Kicked              bool // flag
	Left                bool // flag
	Deactivated         bool // flag
	ID                  int32
	Title               string
	Photo               TL // ChatPhoto
	ParticipantsCount   int32
	Date                int32
	Version             int32
	MigratedTo          TL // InputChannel // flag
	AdminRights         TL // ChatAdminRights // flag
	DefaultBannedRights TL // ChatBannedRights // flag
}

type ChatForbidden struct {
	ID    int32
	Title string
}

type Channel struct {
	Flags               int32
	Creator             bool // flag
	Left                bool // flag
	Broadcast           bool // flag
	Verified            bool // flag
	Megagroup           bool // flag
	Restricted          bool // flag
	Signatures          bool // flag
	Min                 bool // flag
	Scam                bool // flag
	HasLink             bool // flag
	HasGeo              bool // flag
	SlowmodeEnabled     bool // flag
	ID                  int32
	AccessHash          int64 // flag
	Title               string
	Username            string // flag
	Photo               TL     // ChatPhoto
	Date                int32
	Version             int32
	RestrictionReason   []TL  // RestrictionReason // flag
	AdminRights         TL    // ChatAdminRights // flag
	BannedRights        TL    // ChatBannedRights // flag
	DefaultBannedRights TL    // ChatBannedRights // flag
	ParticipantsCount   int32 // flag
}

type ChannelForbidden struct {
	Flags      int32
	Broadcast  bool // flag
	Megagroup  bool // flag
	ID         int32
	AccessHash int64
	Title      string
	UntilDate  int32 // flag
}

type ChatFull struct {
	Flags          int32
	CanSetUsername bool // flag
	HasScheduled   bool // flag
	ID             int32
	About          string
	Participants   TL    // ChatParticipants
	ChatPhoto      TL    // Photo // flag
	NotifySettings TL    // PeerNotifySettings
	ExportedInvite TL    // ExportedChatInvite
	BotInfo        []TL  // BotInfo // flag
	PinnedMsgID    int32 // flag
	FolderID       int32 // flag
}

type ChannelFull struct {
	Flags                int32
	CanViewParticipants  bool // flag
	CanSetUsername       bool // flag
	CanSetStickers       bool // flag
	HiddenPrehistory     bool // flag
	CanViewStats         bool // flag
	CanSetLocation       bool // flag
	HasScheduled         bool // flag
	ID                   int32
	About                string
	ParticipantsCount    int32 // flag
	AdminsCount          int32 // flag
	KickedCount          int32 // flag
	BannedCount          int32 // flag
	OnlineCount          int32 // flag
	ReadInboxMaxID       int32
	ReadOutboxMaxID      int32
	UnreadCount          int32
	ChatPhoto            TL    // Photo
	NotifySettings       TL    // PeerNotifySettings
	ExportedInvite       TL    // ExportedChatInvite
	BotInfo              []TL  // BotInfo
	MigratedFromChatID   int32 // flag
	MigratedFromMaxID    int32 // flag
	PinnedMsgID          int32 // flag
	Stickerset           TL    // StickerSet // flag
	AvailableMinID       int32 // flag
	FolderID             int32 // flag
	LinkedChatID         int32 // flag
	Location             TL    // ChannelLocation // flag
	SlowmodeSeconds      int32 // flag
	SlowmodeNextSendDate int32 // flag
	StatsDc              int32 // flag
	Pts                  int32
}

type ChatParticipant struct {
	UserID    int32
	InviterID int32
	Date      int32
}

type ChatParticipantCreator struct {
	UserID int32
}

type ChatParticipantAdmin struct {
	UserID    int32
	InviterID int32
	Date      int32
}

type ChatParticipantsForbidden struct {
	Flags           int32
	ChatID          int32
	SelfParticipant TL // ChatParticipant // flag
}

type ChatParticipants struct {
	ChatID       int32
	Participants []TL // ChatParticipant
	Version      int32
}

type ChatPhotoEmpty struct {
}

type ChatPhoto struct {
	Flags      int32
	HasVideo   bool // flag
	PhotoSmall TL   // FileLocation
	PhotoBig   TL   // FileLocation
	DcID       int32
}

type MessageEmpty struct {
	ID int32
}

type Message struct {
	Flags             int32
	Out               bool // flag
	Mentioned         bool // flag
	MediaUnread       bool // flag
	Silent            bool // flag
	Post              bool // flag
	FromScheduled     bool // flag
	Legacy            bool // flag
	EditHide          bool // flag
	ID                int32
	FromID            int32 // flag
	ToID              TL    // Peer
	FwdFrom           TL    // MessageFwdHeader // flag
	ViaBotID          int32 // flag
	ReplyToMsgID      int32 // flag
	Date              int32
	Message           string
	Media             TL     // MessageMedia // flag
	ReplyMarkup       TL     // ReplyMarkup // flag
	Entities          []TL   // MessageEntity // flag
	Views             int32  // flag
	EditDate          int32  // flag
	PostAuthor        string // flag
	GroupedID         int64  // flag
	RestrictionReason []TL   // RestrictionReason // flag
}

type MessageService struct {
	Flags        int32
	Out          bool // flag
	Mentioned    bool // flag
	MediaUnread  bool // flag
	Silent       bool // flag
	Post         bool // flag
	Legacy       bool // flag
	ID           int32
	FromID       int32 // flag
	ToID         TL    // Peer
	ReplyToMsgID int32 // flag
	Date         int32
	Action       TL // MessageAction
}

type MessageMediaEmpty struct {
}

type MessageMediaPhoto struct {
	Flags      int32
	Photo      TL    // Photo // flag
	TtlSeconds int32 // flag
}

type MessageMediaGeo struct {
	Geo TL // GeoPoint
}

type MessageMediaContact struct {
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	UserID      int32
}

type MessageMediaUnsupported struct {
}

type MessageMediaDocument struct {
	Flags      int32
	Document   TL    // Document // flag
	TtlSeconds int32 // flag
}

type MessageMediaWebPage struct {
	Webpage TL // WebPage
}

type MessageMediaVenue struct {
	Geo       TL // GeoPoint
	Title     string
	Address   string
	Provider  string
	VenueID   string
	VenueType string
}

type MessageMediaGame struct {
	Game TL // Game
}

type MessageMediaInvoice struct {
	Flags                    int32
	ShippingAddressRequested bool // flag
	Test                     bool // flag
	Title                    string
	Description              string
	Photo                    TL    // WebDocument // flag
	ReceiptMsgID             int32 // flag
	Currency                 string
	TotalAmount              int64
	StartParam               string
}

type MessageMediaGeoLive struct {
	Geo    TL // GeoPoint
	Period int32
}

type MessageMediaPoll struct {
	Poll    TL // Poll
	Results TL // PollResults
}

type MessageMediaDice struct {
	Value    int32
	Emoticon string
}

type MessageActionEmpty struct {
}

type MessageActionChatCreate struct {
	Title string
	Users []int32
}

type MessageActionChatEditTitle struct {
	Title string
}

type MessageActionChatEditPhoto struct {
	Photo TL // Photo
}

type MessageActionChatDeletePhoto struct {
}

type MessageActionChatAddUser struct {
	Users []int32
}

type MessageActionChatDeleteUser struct {
	UserID int32
}

type MessageActionChatJoinedByLink struct {
	InviterID int32
}

type MessageActionChannelCreate struct {
	Title string
}

type MessageActionChatMigrateTo struct {
	ChannelID int32
}

type MessageActionChannelMigrateFrom struct {
	Title  string
	ChatID int32
}

type MessageActionPinMessage struct {
}

type MessageActionHistoryClear struct {
}

type MessageActionGameScore struct {
	GameID int64
	Score  int32
}

type MessageActionPaymentSentMe struct {
	Flags            int32
	Currency         string
	TotalAmount      int64
	Payload          []byte
	Info             TL     // PaymentRequestedInfo // flag
	ShippingOptionID string // flag
	Charge           TL     // PaymentCharge
}

type MessageActionPaymentSent struct {
	Currency    string
	TotalAmount int64
}

type MessageActionPhoneCall struct {
	Flags    int32
	Video    bool // flag
	CallID   int64
	Reason   TL    // PhoneCallDiscardReason // flag
	Duration int32 // flag
}

type MessageActionScreenshotTaken struct {
}

type MessageActionCustomAction struct {
	Message string
}

type MessageActionBotAllowed struct {
	Domain string
}

type MessageActionSecureValuesSentMe struct {
	Values      []TL // SecureValue
	Credentials TL   // SecureCredentialsEncrypted
}

type MessageActionSecureValuesSent struct {
	Types []TL // SecureValueType
}

type MessageActionContactSignUp struct {
}

type Dialog struct {
	Flags               int32
	Pinned              bool // flag
	UnreadMark          bool // flag
	Peer                TL   // Peer
	TopMessage          int32
	ReadInboxMaxID      int32
	ReadOutboxMaxID     int32
	UnreadCount         int32
	UnreadMentionsCount int32
	NotifySettings      TL    // PeerNotifySettings
	Pts                 int32 // flag
	Draft               TL    // DraftMessage // flag
	FolderID            int32 // flag
}

type DialogFolder struct {
	Flags                      int32
	Pinned                     bool // flag
	Folder                     TL   // Folder
	Peer                       TL   // Peer
	TopMessage                 int32
	UnreadMutedPeersCount      int32
	UnreadUnmutedPeersCount    int32
	UnreadMutedMessagesCount   int32
	UnreadUnmutedMessagesCount int32
}

type PhotoEmpty struct {
	ID int64
}

type Photo struct {
	Flags         int32
	HasStickers   bool // flag
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	Sizes         []TL // PhotoSize
	VideoSizes    []TL // VideoSize // flag
	DcID          int32
}

type PhotoSizeEmpty struct {
	Type string
}

type PhotoSize struct {
	Type     string
	Location TL // FileLocation
	W        int32
	H        int32
	Size     int32
}

type PhotoCachedSize struct {
	Type     string
	Location TL // FileLocation
	W        int32
	H        int32
	Bytes    []byte
}

type PhotoStrippedSize struct {
	Type  string
	Bytes []byte
}

type GeoPointEmpty struct {
}

type GeoPoint struct {
	Long       float64
	Lat        float64
	AccessHash int64
}

type AuthSentCode struct {
	Flags         int32
	Type          TL // auth_SentCodeType
	PhoneCodeHash string
	NextType      TL    // auth_CodeType // flag
	Timeout       int32 // flag
}

type AuthAuthorization struct {
	Flags       int32
	TmpSessions int32 // flag
	User        TL    // User
}

type AuthAuthorizationSignUpRequired struct {
	Flags          int32
	TermsOfService TL // help_TermsOfService // flag
}

type AuthExportedAuthorization struct {
	ID    int32
	Bytes []byte
}

type InputNotifyPeer struct {
	Peer TL // InputPeer
}

type InputNotifyUsers struct {
}

type InputNotifyChats struct {
}

type InputNotifyBroadcasts struct {
}

type InputPeerNotifySettings struct {
	Flags        int32
	ShowPreviews TL     // Bool // flag
	Silent       TL     // Bool // flag
	MuteUntil    int32  // flag
	Sound        string // flag
}

type PeerNotifySettings struct {
	Flags        int32
	ShowPreviews TL     // Bool // flag
	Silent       TL     // Bool // flag
	MuteUntil    int32  // flag
	Sound        string // flag
}

type PeerSettings struct {
	Flags                 int32
	ReportSpam            bool  // flag
	AddContact            bool  // flag
	BlockContact          bool  // flag
	ShareContact          bool  // flag
	NeedContactsException bool  // flag
	ReportGeo             bool  // flag
	Autoarchived          bool  // flag
	GeoDistance           int32 // flag
}

type WallPaper struct {
	ID         int64
	Flags      int32
	Creator    bool // flag
	Default    bool // flag
	Pattern    bool // flag
	Dark       bool // flag
	AccessHash int64
	Slug       string
	Document   TL // Document
	Settings   TL // WallPaperSettings // flag
}

type WallPaperNoFile struct {
	Flags    int32
	Default  bool // flag
	Dark     bool // flag
	Settings TL   // WallPaperSettings // flag
}

type InputReportReasonSpam struct {
}

type InputReportReasonViolence struct {
}

type InputReportReasonPornography struct {
}

type InputReportReasonChildAbuse struct {
}

type InputReportReasonOther struct {
	Text string
}

type InputReportReasonCopyright struct {
}

type InputReportReasonGeoIrrelevant struct {
}

type UserFull struct {
	Flags               int32
	Blocked             bool   // flag
	PhoneCallsAvailable bool   // flag
	PhoneCallsPrivate   bool   // flag
	CanPinMessage       bool   // flag
	HasScheduled        bool   // flag
	User                TL     // User
	About               string // flag
	Settings            TL     // PeerSettings
	ProfilePhoto        TL     // Photo // flag
	NotifySettings      TL     // PeerNotifySettings
	BotInfo             TL     // BotInfo // flag
	PinnedMsgID         int32  // flag
	CommonChatsCount    int32
	FolderID            int32 // flag
}

type Contact struct {
	UserID int32
	Mutual TL // Bool
}

type ImportedContact struct {
	UserID   int32
	ClientID int64
}

type ContactBlocked struct {
	UserID int32
	Date   int32
}

type ContactStatus struct {
	UserID int32
	Status TL // UserStatus
}

type ContactsContactsNotModified struct {
}

type ContactsContacts struct {
	Contacts   []TL // Contact
	SavedCount int32
	Users      []TL // User
}

type ContactsImportedContacts struct {
	Imported       []TL // ImportedContact
	PopularInvites []TL // PopularContact
	RetryContacts  []int64
	Users          []TL // User
}

type ContactsBlocked struct {
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type ContactsBlockedSlice struct {
	Count   int32
	Blocked []TL // ContactBlocked
	Users   []TL // User
}

type MessagesDialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type MessagesDialogsSlice struct {
	Count    int32
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type MessagesDialogsNotModified struct {
	Count int32
}

type MessagesMessages struct {
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type MessagesMessagesSlice struct {
	Flags    int32
	Inexact  bool // flag
	Count    int32
	NextRate int32 // flag
	Messages []TL  // Message
	Chats    []TL  // Chat
	Users    []TL  // User
}

type MessagesChannelMessages struct {
	Flags    int32
	Inexact  bool // flag
	Pts      int32
	Count    int32
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
}

type MessagesMessagesNotModified struct {
	Count int32
}

type MessagesChats struct {
	Chats []TL // Chat
}

type MessagesChatsSlice struct {
	Count int32
	Chats []TL // Chat
}

type MessagesChatFull struct {
	FullChat TL   // ChatFull
	Chats    []TL // Chat
	Users    []TL // User
}

type MessagesAffectedHistory struct {
	Pts      int32
	PtsCount int32
	Offset   int32
}

type InputMessagesFilterEmpty struct {
}

type InputMessagesFilterPhotos struct {
}

type InputMessagesFilterVideo struct {
}

type InputMessagesFilterPhotoVideo struct {
}

type InputMessagesFilterDocument struct {
}

type InputMessagesFilterUrl struct {
}

type InputMessagesFilterGif struct {
}

type InputMessagesFilterVoice struct {
}

type InputMessagesFilterMusic struct {
}

type InputMessagesFilterChatPhotos struct {
}

type InputMessagesFilterPhoneCalls struct {
	Flags  int32
	Missed bool // flag
}

type InputMessagesFilterRoundVoice struct {
}

type InputMessagesFilterRoundVideo struct {
}

type InputMessagesFilterMyMentions struct {
}

type InputMessagesFilterGeo struct {
}

type InputMessagesFilterContacts struct {
}

type UpdateNewMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type UpdateMessageID struct {
	ID       int32
	RandomID int64
}

type UpdateDeleteMessages struct {
	Messages []int32
	Pts      int32
	PtsCount int32
}

type UpdateUserTyping struct {
	UserID int32
	Action TL // SendMessageAction
}

type UpdateChatUserTyping struct {
	ChatID int32
	UserID int32
	Action TL // SendMessageAction
}

type UpdateChatParticipants struct {
	Participants TL // ChatParticipants
}

type UpdateUserStatus struct {
	UserID int32
	Status TL // UserStatus
}

type UpdateUserName struct {
	UserID    int32
	FirstName string
	LastName  string
	Username  string
}

type UpdateUserPhoto struct {
	UserID   int32
	Date     int32
	Photo    TL // UserProfilePhoto
	Previous TL // Bool
}

type UpdateNewEncryptedMessage struct {
	Message TL // EncryptedMessage
	Qts     int32
}

type UpdateEncryptedChatTyping struct {
	ChatID int32
}

type UpdateEncryption struct {
	Chat TL // EncryptedChat
	Date int32
}

type UpdateEncryptedMessagesRead struct {
	ChatID  int32
	MaxDate int32
	Date    int32
}

type UpdateChatParticipantAdd struct {
	ChatID    int32
	UserID    int32
	InviterID int32
	Date      int32
	Version   int32
}

type UpdateChatParticipantDelete struct {
	ChatID  int32
	UserID  int32
	Version int32
}

type UpdateDcOptions struct {
	DcOptions []TL // DcOption
}

type UpdateUserBlocked struct {
	UserID  int32
	Blocked TL // Bool
}

type UpdateNotifySettings struct {
	Peer           TL // NotifyPeer
	NotifySettings TL // PeerNotifySettings
}

type UpdateServiceNotification struct {
	Flags     int32
	Popup     bool  // flag
	InboxDate int32 // flag
	Type      string
	Message   string
	Media     TL   // MessageMedia
	Entities  []TL // MessageEntity
}

type UpdatePrivacy struct {
	Key   TL   // PrivacyKey
	Rules []TL // PrivacyRule
}

type UpdateUserPhone struct {
	UserID int32
	Phone  string
}

type UpdateReadHistoryInbox struct {
	Flags            int32
	FolderID         int32 // flag
	Peer             TL    // Peer
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
	PtsCount         int32
}

type UpdateReadHistoryOutbox struct {
	Peer     TL // Peer
	MaxID    int32
	Pts      int32
	PtsCount int32
}

type UpdateWebPage struct {
	Webpage  TL // WebPage
	Pts      int32
	PtsCount int32
}

type UpdateReadMessagesContents struct {
	Messages []int32
	Pts      int32
	PtsCount int32
}

type UpdateChannelTooLong struct {
	Flags     int32
	ChannelID int32
	Pts       int32 // flag
}

type UpdateChannel struct {
	ChannelID int32
}

type UpdateNewChannelMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type UpdateReadChannelInbox struct {
	Flags            int32
	FolderID         int32 // flag
	ChannelID        int32
	MaxID            int32
	StillUnreadCount int32
	Pts              int32
}

type UpdateDeleteChannelMessages struct {
	ChannelID int32
	Messages  []int32
	Pts       int32
	PtsCount  int32
}

type UpdateChannelMessageViews struct {
	ChannelID int32
	ID        int32
	Views     int32
}

type UpdateChatParticipantAdmin struct {
	ChatID  int32
	UserID  int32
	IsAdmin TL // Bool
	Version int32
}

type UpdateNewStickerSet struct {
	Stickerset TL // messages_StickerSet
}

type UpdateStickerSetsOrder struct {
	Flags int32
	Masks bool // flag
	Order []int64
}

type UpdateStickerSets struct {
}

type UpdateSavedGifs struct {
}

type UpdateBotInlineQuery struct {
	Flags   int32
	QueryID int64
	UserID  int32
	Query   string
	Geo     TL // GeoPoint // flag
	Offset  string
}

type UpdateBotInlineSend struct {
	Flags  int32
	UserID int32
	Query  string
	Geo    TL // GeoPoint // flag
	ID     string
	MsgID  TL // InputBotInlineMessageID // flag
}

type UpdateEditChannelMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type UpdateChannelPinnedMessage struct {
	ChannelID int32
	ID        int32
}

type UpdateBotCallbackQuery struct {
	Flags         int32
	QueryID       int64
	UserID        int32
	Peer          TL // Peer
	MsgID         int32
	ChatInstance  int64
	Data          []byte // flag
	GameShortName string // flag
}

type UpdateEditMessage struct {
	Message  TL // Message
	Pts      int32
	PtsCount int32
}

type UpdateInlineBotCallbackQuery struct {
	Flags         int32
	QueryID       int64
	UserID        int32
	MsgID         TL // InputBotInlineMessageID
	ChatInstance  int64
	Data          []byte // flag
	GameShortName string // flag
}

type UpdateReadChannelOutbox struct {
	ChannelID int32
	MaxID     int32
}

type UpdateDraftMessage struct {
	Peer  TL // Peer
	Draft TL // DraftMessage
}

type UpdateReadFeaturedStickers struct {
}

type UpdateRecentStickers struct {
}

type UpdateConfig struct {
}

type UpdatePtsChanged struct {
}

type UpdateChannelWebPage struct {
	ChannelID int32
	Webpage   TL // WebPage
	Pts       int32
	PtsCount  int32
}

type UpdateDialogPinned struct {
	Flags    int32
	Pinned   bool  // flag
	FolderID int32 // flag
	Peer     TL    // DialogPeer
}

type UpdatePinnedDialogs struct {
	Flags    int32
	FolderID int32 // flag
	Order    []TL  // DialogPeer // flag
}

type UpdateBotWebhookJSON struct {
	Data TL // DataJSON
}

type UpdateBotWebhookJSONQuery struct {
	QueryID int64
	Data    TL // DataJSON
	Timeout int32
}

type UpdateBotShippingQuery struct {
	QueryID         int64
	UserID          int32
	Payload         []byte
	ShippingAddress TL // PostAddress
}

type UpdateBotPrecheckoutQuery struct {
	Flags            int32
	QueryID          int64
	UserID           int32
	Payload          []byte
	Info             TL     // PaymentRequestedInfo // flag
	ShippingOptionID string // flag
	Currency         string
	TotalAmount      int64
}

type UpdatePhoneCall struct {
	PhoneCall TL // PhoneCall
}

type UpdateLangPackTooLong struct {
	LangCode string
}

type UpdateLangPack struct {
	Difference TL // LangPackDifference
}

type UpdateFavedStickers struct {
}

type UpdateChannelReadMessagesContents struct {
	ChannelID int32
	Messages  []int32
}

type UpdateContactsReset struct {
}

type UpdateChannelAvailableMessages struct {
	ChannelID      int32
	AvailableMinID int32
}

type UpdateDialogUnreadMark struct {
	Flags  int32
	Unread bool // flag
	Peer   TL   // DialogPeer
}

type UpdateUserPinnedMessage struct {
	UserID int32
	ID     int32
}

type UpdateChatPinnedMessage struct {
	ChatID  int32
	ID      int32
	Version int32
}

type UpdateMessagePoll struct {
	Flags   int32
	PollID  int64
	Poll    TL // Poll // flag
	Results TL // PollResults
}

type UpdateChatDefaultBannedRights struct {
	Peer                TL // Peer
	DefaultBannedRights TL // ChatBannedRights
	Version             int32
}

type UpdateFolderPeers struct {
	FolderPeers []TL // FolderPeer
	Pts         int32
	PtsCount    int32
}

type UpdatePeerSettings struct {
	Peer     TL // Peer
	Settings TL // PeerSettings
}

type UpdatePeerLocated struct {
	Peers []TL // PeerLocated
}

type UpdateNewScheduledMessage struct {
	Message TL // Message
}

type UpdateDeleteScheduledMessages struct {
	Peer     TL // Peer
	Messages []int32
}

type UpdateTheme struct {
	Theme TL // Theme
}

type UpdateGeoLiveViewed struct {
	Peer  TL // Peer
	MsgID int32
}

type UpdateLoginToken struct {
}

type UpdateMessagePollVote struct {
	PollID  int64
	UserID  int32
	Options []TL // Bytes
}

type UpdateDialogFilter struct {
	Flags  int32
	ID     int32
	Filter TL // DialogFilter // flag
}

type UpdateDialogFilterOrder struct {
	Order []int32
}

type UpdateDialogFilters struct {
}

type UpdatePhoneCallSignalingData struct {
	PhoneCallID int64
	Data        []byte
}

type UpdateChannelParticipant struct {
	Flags           int32
	ChannelID       int32
	Date            int32
	UserID          int32
	PrevParticipant TL // ChannelParticipant // flag
	NewParticipant  TL // ChannelParticipant // flag
	Qts             int32
}

type UpdatesState struct {
	Pts         int32
	Qts         int32
	Date        int32
	Seq         int32
	UnreadCount int32
}

type UpdatesDifferenceEmpty struct {
	Date int32
	Seq  int32
}

type UpdatesDifference struct {
	NewMessages          []TL // Message
	NewEncryptedMessages []TL // EncryptedMessage
	OtherUpdates         []TL // Update
	Chats                []TL // Chat
	Users                []TL // User
	State                TL   // updates_State
}

type UpdatesDifferenceSlice struct {
	NewMessages          []TL // Message
	NewEncryptedMessages []TL // EncryptedMessage
	OtherUpdates         []TL // Update
	Chats                []TL // Chat
	Users                []TL // User
	IntermediateState    TL   // updates_State
}

type UpdatesDifferenceTooLong struct {
	Pts int32
}

type UpdatesTooLong struct {
}

type UpdateShortMessage struct {
	Flags        int32
	Out          bool // flag
	Mentioned    bool // flag
	MediaUnread  bool // flag
	Silent       bool // flag
	ID           int32
	UserID       int32
	Message      string
	Pts          int32
	PtsCount     int32
	Date         int32
	FwdFrom      TL    // MessageFwdHeader // flag
	ViaBotID     int32 // flag
	ReplyToMsgID int32 // flag
	Entities     []TL  // MessageEntity // flag
}

type UpdateShortChatMessage struct {
	Flags        int32
	Out          bool // flag
	Mentioned    bool // flag
	MediaUnread  bool // flag
	Silent       bool // flag
	ID           int32
	FromID       int32
	ChatID       int32
	Message      string
	Pts          int32
	PtsCount     int32
	Date         int32
	FwdFrom      TL    // MessageFwdHeader // flag
	ViaBotID     int32 // flag
	ReplyToMsgID int32 // flag
	Entities     []TL  // MessageEntity // flag
}

type UpdateShort struct {
	Update TL // Update
	Date   int32
}

type UpdatesCombined struct {
	Updates  []TL // Update
	Users    []TL // User
	Chats    []TL // Chat
	Date     int32
	SeqStart int32
	Seq      int32
}

type Updates struct {
	Updates []TL // Update
	Users   []TL // User
	Chats   []TL // Chat
	Date    int32
	Seq     int32
}

type UpdateShortSentMessage struct {
	Flags    int32
	Out      bool // flag
	ID       int32
	Pts      int32
	PtsCount int32
	Date     int32
	Media    TL   // MessageMedia // flag
	Entities []TL // MessageEntity // flag
}

type PhotosPhotos struct {
	Photos []TL // Photo
	Users  []TL // User
}

type PhotosPhotosSlice struct {
	Count  int32
	Photos []TL // Photo
	Users  []TL // User
}

type PhotosPhoto struct {
	Photo TL   // Photo
	Users []TL // User
}

type UploadFile struct {
	Type  TL // storage_FileType
	Mtime int32
	Bytes []byte
}

type UploadFileCdnRedirect struct {
	DcID          int32
	FileToken     []byte
	EncryptionKey []byte
	EncryptionIv  []byte
	FileHashes    []TL // FileHash
}

type DcOption struct {
	Flags     int32
	Ipv6      bool // flag
	MediaOnly bool // flag
	TcpoOnly  bool // flag
	Cdn       bool // flag
	Static    bool // flag
	ID        int32
	IpAddress string
	Port      int32
	Secret    []byte // flag
}

type Config struct {
	Flags                   int32
	PhonecallsEnabled       bool // flag
	DefaultP2pContacts      bool // flag
	PreloadFeaturedStickers bool // flag
	IgnorePhoneEntities     bool // flag
	RevokePmInbox           bool // flag
	BlockedMode             bool // flag
	PfsEnabled              bool // flag
	Date                    int32
	Expires                 int32
	TestMode                TL // Bool
	ThisDc                  int32
	DcOptions               []TL // DcOption
	DcTxtDomainName         string
	ChatSizeMax             int32
	MegagroupSizeMax        int32
	ForwardedCountMax       int32
	OnlineUpdatePeriodMs    int32
	OfflineBlurTimeoutMs    int32
	OfflineIdleTimeoutMs    int32
	OnlineCloudTimeoutMs    int32
	NotifyCloudDelayMs      int32
	NotifyDefaultDelayMs    int32
	PushChatPeriodMs        int32
	PushChatLimit           int32
	SavedGifsLimit          int32
	EditTimeLimit           int32
	RevokeTimeLimit         int32
	RevokePmTimeLimit       int32
	RatingEDecay            int32
	StickersRecentLimit     int32
	StickersFavedLimit      int32
	ChannelsReadMediaPeriod int32
	TmpSessions             int32 // flag
	PinnedDialogsCountMax   int32
	PinnedInfolderCountMax  int32
	CallReceiveTimeoutMs    int32
	CallRingTimeoutMs       int32
	CallConnectTimeoutMs    int32
	CallPacketTimeoutMs     int32
	MeUrlPrefix             string
	AutoupdateUrlPrefix     string // flag
	GifSearchUsername       string // flag
	VenueSearchUsername     string // flag
	ImgSearchUsername       string // flag
	StaticMapsProvider      string // flag
	CaptionLengthMax        int32
	MessageLengthMax        int32
	WebfileDcID             int32
	SuggestedLangCode       string // flag
	LangPackVersion         int32  // flag
	BaseLangPackVersion     int32  // flag
}

type NearestDc struct {
	Country   string
	ThisDc    int32
	NearestDc int32
}

type HelpAppUpdate struct {
	Flags      int32
	CanNotSkip bool // flag
	ID         int32
	Version    string
	Text       string
	Entities   []TL   // MessageEntity
	Document   TL     // Document // flag
	Url        string // flag
}

type HelpNoAppUpdate struct {
}

type HelpInviteText struct {
	Message string
}

type EncryptedChatEmpty struct {
	ID int32
}

type EncryptedChatWaiting struct {
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
}

type EncryptedChatRequested struct {
	Flags         int32
	FolderID      int32 // flag
	ID            int32
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	GA            []byte
}

type EncryptedChat struct {
	ID             int32
	AccessHash     int64
	Date           int32
	AdminID        int32
	ParticipantID  int32
	GAOrB          []byte
	KeyFingerprint int64
}

type EncryptedChatDiscarded struct {
	ID int32
}

type InputEncryptedChat struct {
	ChatID     int32
	AccessHash int64
}

type EncryptedFileEmpty struct {
}

type EncryptedFile struct {
	ID             int64
	AccessHash     int64
	Size           int32
	DcID           int32
	KeyFingerprint int32
}

type InputEncryptedFileEmpty struct {
}

type InputEncryptedFileUploaded struct {
	ID             int64
	Parts          int32
	Md5Checksum    string
	KeyFingerprint int32
}

type InputEncryptedFile struct {
	ID         int64
	AccessHash int64
}

type InputEncryptedFileBigUploaded struct {
	ID             int64
	Parts          int32
	KeyFingerprint int32
}

type EncryptedMessage struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
	File     TL // EncryptedFile
}

type EncryptedMessageService struct {
	RandomID int64
	ChatID   int32
	Date     int32
	Bytes    []byte
}

type MessagesDhConfigNotModified struct {
	Random []byte
}

type MessagesDhConfig struct {
	G       int32
	P       []byte
	Version int32
	Random  []byte
}

type MessagesSentEncryptedMessage struct {
	Date int32
}

type MessagesSentEncryptedFile struct {
	Date int32
	File TL // EncryptedFile
}

type InputDocumentEmpty struct {
}

type InputDocument struct {
	ID            int64
	AccessHash    int64
	FileReference []byte
}

type DocumentEmpty struct {
	ID int64
}

type Document struct {
	Flags         int32
	ID            int64
	AccessHash    int64
	FileReference []byte
	Date          int32
	MimeType      string
	Size          int32
	Thumbs        []TL // PhotoSize // flag
	VideoThumbs   []TL // VideoSize // flag
	DcID          int32
	Attributes    []TL // DocumentAttribute
}

type HelpSupport struct {
	PhoneNumber string
	User        TL // User
}

type NotifyPeer struct {
	Peer TL // Peer
}

type NotifyUsers struct {
}

type NotifyChats struct {
}

type NotifyBroadcasts struct {
}

type SendMessageTypingAction struct {
}

type SendMessageCancelAction struct {
}

type SendMessageRecordVideoAction struct {
}

type SendMessageUploadVideoAction struct {
	Progress int32
}

type SendMessageRecordAudioAction struct {
}

type SendMessageUploadAudioAction struct {
	Progress int32
}

type SendMessageUploadPhotoAction struct {
	Progress int32
}

type SendMessageUploadDocumentAction struct {
	Progress int32
}

type SendMessageGeoLocationAction struct {
}

type SendMessageChooseContactAction struct {
}

type SendMessageGamePlayAction struct {
}

type SendMessageRecordRoundAction struct {
}

type SendMessageUploadRoundAction struct {
	Progress int32
}

type ContactsFound struct {
	MyResults []TL // Peer
	Results   []TL // Peer
	Chats     []TL // Chat
	Users     []TL // User
}

type InputPrivacyKeyStatusTimestamp struct {
}

type InputPrivacyKeyChatInvite struct {
}

type InputPrivacyKeyPhoneCall struct {
}

type InputPrivacyKeyPhoneP2P struct {
}

type InputPrivacyKeyForwards struct {
}

type InputPrivacyKeyProfilePhoto struct {
}

type InputPrivacyKeyPhoneNumber struct {
}

type InputPrivacyKeyAddedByPhone struct {
}

type PrivacyKeyStatusTimestamp struct {
}

type PrivacyKeyChatInvite struct {
}

type PrivacyKeyPhoneCall struct {
}

type PrivacyKeyPhoneP2P struct {
}

type PrivacyKeyForwards struct {
}

type PrivacyKeyProfilePhoto struct {
}

type PrivacyKeyPhoneNumber struct {
}

type PrivacyKeyAddedByPhone struct {
}

type InputPrivacyValueAllowContacts struct {
}

type InputPrivacyValueAllowAll struct {
}

type InputPrivacyValueAllowUsers struct {
	Users []TL // InputUser
}

type InputPrivacyValueDisallowContacts struct {
}

type InputPrivacyValueDisallowAll struct {
}

type InputPrivacyValueDisallowUsers struct {
	Users []TL // InputUser
}

type InputPrivacyValueAllowChatParticipants struct {
	Chats []int32
}

type InputPrivacyValueDisallowChatParticipants struct {
	Chats []int32
}

type PrivacyValueAllowContacts struct {
}

type PrivacyValueAllowAll struct {
}

type PrivacyValueAllowUsers struct {
	Users []int32
}

type PrivacyValueDisallowContacts struct {
}

type PrivacyValueDisallowAll struct {
}

type PrivacyValueDisallowUsers struct {
	Users []int32
}

type PrivacyValueAllowChatParticipants struct {
	Chats []int32
}

type PrivacyValueDisallowChatParticipants struct {
	Chats []int32
}

type AccountPrivacyRules struct {
	Rules []TL // PrivacyRule
	Chats []TL // Chat
	Users []TL // User
}

type AccountDaysTTL struct {
	Days int32
}

type DocumentAttributeImageSize struct {
	W int32
	H int32
}

type DocumentAttributeAnimated struct {
}

type DocumentAttributeSticker struct {
	Flags      int32
	Mask       bool // flag
	Alt        string
	Stickerset TL // InputStickerSet
	MaskCoords TL // MaskCoords // flag
}

type DocumentAttributeVideo struct {
	Flags             int32
	RoundMessage      bool // flag
	SupportsStreaming bool // flag
	Duration          int32
	W                 int32
	H                 int32
}

type DocumentAttributeAudio struct {
	Flags     int32
	Voice     bool // flag
	Duration  int32
	Title     string // flag
	Performer string // flag
	Waveform  []byte // flag
}

type DocumentAttributeFilename struct {
	FileName string
}

type DocumentAttributeHasStickers struct {
}

type MessagesStickersNotModified struct {
}

type MessagesStickers struct {
	Hash     int32
	Stickers []TL // Document
}

type StickerPack struct {
	Emoticon  string
	Documents []int64
}

type MessagesAllStickersNotModified struct {
}

type MessagesAllStickers struct {
	Hash int32
	Sets []TL // StickerSet
}

type MessagesAffectedMessages struct {
	Pts      int32
	PtsCount int32
}

type WebPageEmpty struct {
	ID int64
}

type WebPagePending struct {
	ID   int64
	Date int32
}

type WebPage struct {
	Flags       int32
	ID          int64
	Url         string
	DisplayUrl  string
	Hash        int32
	Type        string // flag
	SiteName    string // flag
	Title       string // flag
	Description string // flag
	Photo       TL     // Photo // flag
	EmbedUrl    string // flag
	EmbedType   string // flag
	EmbedWidth  int32  // flag
	EmbedHeight int32  // flag
	Duration    int32  // flag
	Author      string // flag
	Document    TL     // Document // flag
	CachedPage  TL     // Page // flag
	Attributes  []TL   // WebPageAttribute // flag
}

type WebPageNotModified struct {
	Flags           int32
	CachedPageViews int32 // flag
}

type Authorization struct {
	Flags           int32
	Current         bool // flag
	OfficialApp     bool // flag
	PasswordPending bool // flag
	Hash            int64
	DeviceModel     string
	Platform        string
	SystemVersion   string
	ApiID           int32
	AppName         string
	AppVersion      string
	DateCreated     int32
	DateActive      int32
	Ip              string
	Country         string
	Region          string
}

type AccountAuthorizations struct {
	Authorizations []TL // Authorization
}

type AccountPassword struct {
	Flags                   int32
	HasRecovery             bool   // flag
	HasSecureValues         bool   // flag
	HasPassword             bool   // flag
	CurrentAlgo             TL     // PasswordKdfAlgo // flag
	SrpB                    []byte // flag
	SrpID                   int64  // flag
	Hint                    string // flag
	EmailUnconfirmedPattern string // flag
	NewAlgo                 TL     // PasswordKdfAlgo
	NewSecureAlgo           TL     // SecurePasswordKdfAlgo
	SecureRandom            []byte
}

type AccountPasswordSettings struct {
	Flags          int32
	Email          string // flag
	SecureSettings TL     // SecureSecretSettings // flag
}

type AccountPasswordInputSettings struct {
	Flags             int32
	NewAlgo           TL     // PasswordKdfAlgo // flag
	NewPasswordHash   []byte // flag
	Hint              string // flag
	Email             string // flag
	NewSecureSettings TL     // SecureSecretSettings // flag
}

type AuthPasswordRecovery struct {
	EmailPattern string
}

type ReceivedNotifyMessage struct {
	ID    int32
	Flags int32
}

type ChatInviteEmpty struct {
}

type ChatInviteExported struct {
	Link string
}

type ChatInviteAlready struct {
	Chat TL // Chat
}

type ChatInvite struct {
	Flags             int32
	Channel           bool // flag
	Broadcast         bool // flag
	Public            bool // flag
	Megagroup         bool // flag
	Title             string
	Photo             TL // Photo
	ParticipantsCount int32
	Participants      []TL // User // flag
}

type ChatInvitePeek struct {
	Chat    TL // Chat
	Expires int32
}

type InputStickerSetEmpty struct {
}

type InputStickerSetID struct {
	ID         int64
	AccessHash int64
}

type InputStickerSetShortName struct {
	ShortName string
}

type InputStickerSetAnimatedEmoji struct {
}

type InputStickerSetDice struct {
	Emoticon string
}

type StickerSet struct {
	Flags         int32
	Archived      bool  // flag
	Official      bool  // flag
	Masks         bool  // flag
	Animated      bool  // flag
	InstalledDate int32 // flag
	ID            int64
	AccessHash    int64
	Title         string
	ShortName     string
	Thumb         TL    // PhotoSize // flag
	ThumbDcID     int32 // flag
	Count         int32
	Hash          int32
}

type MessagesStickerSet struct {
	Set       TL   // StickerSet
	Packs     []TL // StickerPack
	Documents []TL // Document
}

type BotCommand struct {
	Command     string
	Description string
}

type BotInfo struct {
	UserID      int32
	Description string
	Commands    []TL // BotCommand
}

type KeyboardButton struct {
	Text string
}

type KeyboardButtonUrl struct {
	Text string
	Url  string
}

type KeyboardButtonCallback struct {
	Text string
	Data []byte
}

type KeyboardButtonRequestPhone struct {
	Text string
}

type KeyboardButtonRequestGeoLocation struct {
	Text string
}

type KeyboardButtonSwitchInline struct {
	Flags    int32
	SamePeer bool // flag
	Text     string
	Query    string
}

type KeyboardButtonGame struct {
	Text string
}

type KeyboardButtonBuy struct {
	Text string
}

type KeyboardButtonUrlAuth struct {
	Flags    int32
	Text     string
	FwdText  string // flag
	Url      string
	ButtonID int32
}

type InputKeyboardButtonUrlAuth struct {
	Flags              int32
	RequestWriteAccess bool // flag
	Text               string
	FwdText            string // flag
	Url                string
	Bot                TL // InputUser
}

type KeyboardButtonRequestPoll struct {
	Flags int32
	Quiz  TL // Bool // flag
	Text  string
}

type KeyboardButtonRow struct {
	Buttons []TL // KeyboardButton
}

type ReplyKeyboardHide struct {
	Flags     int32
	Selective bool // flag
}

type ReplyKeyboardForceReply struct {
	Flags     int32
	SingleUse bool // flag
	Selective bool // flag
}

type ReplyKeyboardMarkup struct {
	Flags     int32
	Resize    bool // flag
	SingleUse bool // flag
	Selective bool // flag
	Rows      []TL // KeyboardButtonRow
}

type ReplyInlineMarkup struct {
	Rows []TL // KeyboardButtonRow
}

type MessageEntityUnknown struct {
	Offset int32
	Length int32
}

type MessageEntityMention struct {
	Offset int32
	Length int32
}

type MessageEntityHashtag struct {
	Offset int32
	Length int32
}

type MessageEntityBotCommand struct {
	Offset int32
	Length int32
}

type MessageEntityUrl struct {
	Offset int32
	Length int32
}

type MessageEntityEmail struct {
	Offset int32
	Length int32
}

type MessageEntityBold struct {
	Offset int32
	Length int32
}

type MessageEntityItalic struct {
	Offset int32
	Length int32
}

type MessageEntityCode struct {
	Offset int32
	Length int32
}

type MessageEntityPre struct {
	Offset   int32
	Length   int32
	Language string
}

type MessageEntityTextUrl struct {
	Offset int32
	Length int32
	Url    string
}

type MessageEntityMentionName struct {
	Offset int32
	Length int32
	UserID int32
}

type InputMessageEntityMentionName struct {
	Offset int32
	Length int32
	UserID TL // InputUser
}

type MessageEntityPhone struct {
	Offset int32
	Length int32
}

type MessageEntityCashtag struct {
	Offset int32
	Length int32
}

type MessageEntityUnderline struct {
	Offset int32
	Length int32
}

type MessageEntityStrike struct {
	Offset int32
	Length int32
}

type MessageEntityBlockquote struct {
	Offset int32
	Length int32
}

type MessageEntityBankCard struct {
	Offset int32
	Length int32
}

type InputChannelEmpty struct {
}

type InputChannel struct {
	ChannelID  int32
	AccessHash int64
}

type InputChannelFromMessage struct {
	Peer      TL // InputPeer
	MsgID     int32
	ChannelID int32
}

type ContactsResolvedPeer struct {
	Peer  TL   // Peer
	Chats []TL // Chat
	Users []TL // User
}

type MessageRange struct {
	MinID int32
	MaxID int32
}

type UpdatesChannelDifferenceEmpty struct {
	Flags   int32
	Final   bool // flag
	Pts     int32
	Timeout int32 // flag
}

type UpdatesChannelDifferenceTooLong struct {
	Flags    int32
	Final    bool  // flag
	Timeout  int32 // flag
	Dialog   TL    // Dialog
	Messages []TL  // Message
	Chats    []TL  // Chat
	Users    []TL  // User
}

type UpdatesChannelDifference struct {
	Flags        int32
	Final        bool // flag
	Pts          int32
	Timeout      int32 // flag
	NewMessages  []TL  // Message
	OtherUpdates []TL  // Update
	Chats        []TL  // Chat
	Users        []TL  // User
}

type ChannelMessagesFilterEmpty struct {
}

type ChannelMessagesFilter struct {
	Flags              int32
	ExcludeNewMessages bool // flag
	Ranges             []TL // MessageRange
}

type ChannelParticipant struct {
	UserID int32
	Date   int32
}

type ChannelParticipantSelf struct {
	UserID    int32
	InviterID int32
	Date      int32
}

type ChannelParticipantCreator struct {
	Flags  int32
	UserID int32
	Rank   string // flag
}

type ChannelParticipantAdmin struct {
	Flags       int32
	CanEdit     bool // flag
	Self        bool // flag
	UserID      int32
	InviterID   int32 // flag
	PromotedBy  int32
	Date        int32
	AdminRights TL     // ChatAdminRights
	Rank        string // flag
}

type ChannelParticipantBanned struct {
	Flags        int32
	Left         bool // flag
	UserID       int32
	KickedBy     int32
	Date         int32
	BannedRights TL // ChatBannedRights
}

type ChannelParticipantsRecent struct {
}

type ChannelParticipantsAdmins struct {
}

type ChannelParticipantsKicked struct {
	Q string
}

type ChannelParticipantsBots struct {
}

type ChannelParticipantsBanned struct {
	Q string
}

type ChannelParticipantsSearch struct {
	Q string
}

type ChannelParticipantsContacts struct {
	Q string
}

type ChannelsChannelParticipants struct {
	Count        int32
	Participants []TL // ChannelParticipant
	Users        []TL // User
}

type ChannelsChannelParticipantsNotModified struct {
}

type ChannelsChannelParticipant struct {
	Participant TL   // ChannelParticipant
	Users       []TL // User
}

type HelpTermsOfService struct {
	Flags         int32
	Popup         bool // flag
	ID            TL   // DataJSON
	Text          string
	Entities      []TL  // MessageEntity
	MinAgeConfirm int32 // flag
}

type MessagesSavedGifsNotModified struct {
}

type MessagesSavedGifs struct {
	Hash int32
	Gifs []TL // Document
}

type InputBotInlineMessageMediaAuto struct {
	Flags       int32
	Message     string
	Entities    []TL // MessageEntity // flag
	ReplyMarkup TL   // ReplyMarkup // flag
}

type InputBotInlineMessageText struct {
	Flags       int32
	NoWebpage   bool // flag
	Message     string
	Entities    []TL // MessageEntity // flag
	ReplyMarkup TL   // ReplyMarkup // flag
}

type InputBotInlineMessageMediaGeo struct {
	Flags       int32
	GeoPoint    TL // InputGeoPoint
	Period      int32
	ReplyMarkup TL // ReplyMarkup // flag
}

type InputBotInlineMessageMediaVenue struct {
	Flags       int32
	GeoPoint    TL // InputGeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup TL // ReplyMarkup // flag
}

type InputBotInlineMessageMediaContact struct {
	Flags       int32
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup TL // ReplyMarkup // flag
}

type InputBotInlineMessageGame struct {
	Flags       int32
	ReplyMarkup TL // ReplyMarkup // flag
}

type InputBotInlineResult struct {
	Flags       int32
	ID          string
	Type        string
	Title       string // flag
	Description string // flag
	Url         string // flag
	Thumb       TL     // InputWebDocument // flag
	Content     TL     // InputWebDocument // flag
	SendMessage TL     // InputBotInlineMessage
}

type InputBotInlineResultPhoto struct {
	ID          string
	Type        string
	Photo       TL // InputPhoto
	SendMessage TL // InputBotInlineMessage
}

type InputBotInlineResultDocument struct {
	Flags       int32
	ID          string
	Type        string
	Title       string // flag
	Description string // flag
	Document    TL     // InputDocument
	SendMessage TL     // InputBotInlineMessage
}

type InputBotInlineResultGame struct {
	ID          string
	ShortName   string
	SendMessage TL // InputBotInlineMessage
}

type BotInlineMessageMediaAuto struct {
	Flags       int32
	Message     string
	Entities    []TL // MessageEntity // flag
	ReplyMarkup TL   // ReplyMarkup // flag
}

type BotInlineMessageText struct {
	Flags       int32
	NoWebpage   bool // flag
	Message     string
	Entities    []TL // MessageEntity // flag
	ReplyMarkup TL   // ReplyMarkup // flag
}

type BotInlineMessageMediaGeo struct {
	Flags       int32
	Geo         TL // GeoPoint
	Period      int32
	ReplyMarkup TL // ReplyMarkup // flag
}

type BotInlineMessageMediaVenue struct {
	Flags       int32
	Geo         TL // GeoPoint
	Title       string
	Address     string
	Provider    string
	VenueID     string
	VenueType   string
	ReplyMarkup TL // ReplyMarkup // flag
}

type BotInlineMessageMediaContact struct {
	Flags       int32
	PhoneNumber string
	FirstName   string
	LastName    string
	Vcard       string
	ReplyMarkup TL // ReplyMarkup // flag
}

type BotInlineResult struct {
	Flags       int32
	ID          string
	Type        string
	Title       string // flag
	Description string // flag
	Url         string // flag
	Thumb       TL     // WebDocument // flag
	Content     TL     // WebDocument // flag
	SendMessage TL     // BotInlineMessage
}

type BotInlineMediaResult struct {
	Flags       int32
	ID          string
	Type        string
	Photo       TL     // Photo // flag
	Document    TL     // Document // flag
	Title       string // flag
	Description string // flag
	SendMessage TL     // BotInlineMessage
}

type MessagesBotResults struct {
	Flags      int32
	Gallery    bool // flag
	QueryID    int64
	NextOffset string // flag
	SwitchPm   TL     // InlineBotSwitchPM // flag
	Results    []TL   // BotInlineResult
	CacheTime  int32
	Users      []TL // User
}

type ExportedMessageLink struct {
	Link string
	Html string
}

type MessageFwdHeader struct {
	Flags          int32
	FromID         int32  // flag
	FromName       string // flag
	Date           int32
	ChannelID      int32  // flag
	ChannelPost    int32  // flag
	PostAuthor     string // flag
	SavedFromPeer  TL     // Peer // flag
	SavedFromMsgID int32  // flag
	PsaType        string // flag
}

type AuthCodeTypeSms struct {
}

type AuthCodeTypeCall struct {
}

type AuthCodeTypeFlashCall struct {
}

type AuthSentCodeTypeApp struct {
	Length int32
}

type AuthSentCodeTypeSms struct {
	Length int32
}

type AuthSentCodeTypeCall struct {
	Length int32
}

type AuthSentCodeTypeFlashCall struct {
	Pattern string
}

type MessagesBotCallbackAnswer struct {
	Flags     int32
	Alert     bool   // flag
	HasUrl    bool   // flag
	NativeUi  bool   // flag
	Message   string // flag
	Url       string // flag
	CacheTime int32
}

type MessagesMessageEditData struct {
	Flags   int32
	Caption bool // flag
}

type InputBotInlineMessageID struct {
	DcID       int32
	ID         int64
	AccessHash int64
}

type InlineBotSwitchPM struct {
	Text       string
	StartParam string
}

type MessagesPeerDialogs struct {
	Dialogs  []TL // Dialog
	Messages []TL // Message
	Chats    []TL // Chat
	Users    []TL // User
	State    TL   // updates_State
}

type TopPeer struct {
	Peer   TL // Peer
	Rating float64
}

type TopPeerCategoryBotsPM struct {
}

type TopPeerCategoryBotsInline struct {
}

type TopPeerCategoryCorrespondents struct {
}

type TopPeerCategoryGroups struct {
}

type TopPeerCategoryChannels struct {
}

type TopPeerCategoryPhoneCalls struct {
}

type TopPeerCategoryForwardUsers struct {
}

type TopPeerCategoryForwardChats struct {
}

type TopPeerCategoryPeers struct {
	Category TL // TopPeerCategory
	Count    int32
	Peers    []TL // TopPeer
}

type ContactsTopPeersNotModified struct {
}

type ContactsTopPeers struct {
	Categories []TL // TopPeerCategoryPeers
	Chats      []TL // Chat
	Users      []TL // User
}

type ContactsTopPeersDisabled struct {
}

type DraftMessageEmpty struct {
	Flags int32
	Date  int32 // flag
}

type DraftMessage struct {
	Flags        int32
	NoWebpage    bool  // flag
	ReplyToMsgID int32 // flag
	Message      string
	Entities     []TL // MessageEntity // flag
	Date         int32
}

type MessagesFeaturedStickersNotModified struct {
	Count int32
}

type MessagesFeaturedStickers struct {
	Hash   int32
	Count  int32
	Sets   []TL // StickerSetCovered
	Unread []int64
}

type MessagesRecentStickersNotModified struct {
}

type MessagesRecentStickers struct {
	Hash     int32
	Packs    []TL // StickerPack
	Stickers []TL // Document
	Dates    []int32
}

type MessagesArchivedStickers struct {
	Count int32
	Sets  []TL // StickerSetCovered
}

type MessagesStickerSetInstallResultSuccess struct {
}

type MessagesStickerSetInstallResultArchive struct {
	Sets []TL // StickerSetCovered
}

type StickerSetCovered struct {
	Set   TL // StickerSet
	Cover TL // Document
}

type StickerSetMultiCovered struct {
	Set    TL   // StickerSet
	Covers []TL // Document
}

type MaskCoords struct {
	N    int32
	X    float64
	Y    float64
	Zoom float64
}

type InputStickeredMediaPhoto struct {
	ID TL // InputPhoto
}

type InputStickeredMediaDocument struct {
	ID TL // InputDocument
}

type Game struct {
	Flags       int32
	ID          int64
	AccessHash  int64
	ShortName   string
	Title       string
	Description string
	Photo       TL // Photo
	Document    TL // Document // flag
}

type InputGameID struct {
	ID         int64
	AccessHash int64
}

type InputGameShortName struct {
	BotID     TL // InputUser
	ShortName string
}

type HighScore struct {
	Pos    int32
	UserID int32
	Score  int32
}

type MessagesHighScores struct {
	Scores []TL // HighScore
	Users  []TL // User
}

type TextEmpty struct {
}

type TextPlain struct {
	Text string
}

type TextBold struct {
	Text TL // RichText
}

type TextItalic struct {
	Text TL // RichText
}

type TextUnderline struct {
	Text TL // RichText
}

type TextStrike struct {
	Text TL // RichText
}

type TextFixed struct {
	Text TL // RichText
}

type TextUrl struct {
	Text      TL // RichText
	Url       string
	WebpageID int64
}

type TextEmail struct {
	Text  TL // RichText
	Email string
}

type TextConcat struct {
	Texts []TL // RichText
}

type TextSubscript struct {
	Text TL // RichText
}

type TextSuperscript struct {
	Text TL // RichText
}

type TextMarked struct {
	Text TL // RichText
}

type TextPhone struct {
	Text  TL // RichText
	Phone string
}

type TextImage struct {
	DocumentID int64
	W          int32
	H          int32
}

type TextAnchor struct {
	Text TL // RichText
	Name string
}

type PageBlockUnsupported struct {
}

type PageBlockTitle struct {
	Text TL // RichText
}

type PageBlockSubtitle struct {
	Text TL // RichText
}

type PageBlockAuthorDate struct {
	Author        TL // RichText
	PublishedDate int32
}

type PageBlockHeader struct {
	Text TL // RichText
}

type PageBlockSubheader struct {
	Text TL // RichText
}

type PageBlockParagraph struct {
	Text TL // RichText
}

type PageBlockPreformatted struct {
	Text     TL // RichText
	Language string
}

type PageBlockFooter struct {
	Text TL // RichText
}

type PageBlockDivider struct {
}

type PageBlockAnchor struct {
	Name string
}

type PageBlockList struct {
	Items []TL // PageListItem
}

type PageBlockBlockquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type PageBlockPullquote struct {
	Text    TL // RichText
	Caption TL // RichText
}

type PageBlockPhoto struct {
	Flags     int32
	PhotoID   int64
	Caption   TL     // PageCaption
	Url       string // flag
	WebpageID int64  // flag
}

type PageBlockVideo struct {
	Flags    int32
	Autoplay bool // flag
	Loop     bool // flag
	VideoID  int64
	Caption  TL // PageCaption
}

type PageBlockCover struct {
	Cover TL // PageBlock
}

type PageBlockEmbed struct {
	Flags          int32
	FullWidth      bool   // flag
	AllowScrolling bool   // flag
	Url            string // flag
	Html           string // flag
	PosterPhotoID  int64  // flag
	W              int32  // flag
	H              int32  // flag
	Caption        TL     // PageCaption
}

type PageBlockEmbedPost struct {
	Url           string
	WebpageID     int64
	AuthorPhotoID int64
	Author        string
	Date          int32
	Blocks        []TL // PageBlock
	Caption       TL   // PageCaption
}

type PageBlockCollage struct {
	Items   []TL // PageBlock
	Caption TL   // PageCaption
}

type PageBlockSlideshow struct {
	Items   []TL // PageBlock
	Caption TL   // PageCaption
}

type PageBlockChannel struct {
	Channel TL // Chat
}

type PageBlockAudio struct {
	AudioID int64
	Caption TL // PageCaption
}

type PageBlockKicker struct {
	Text TL // RichText
}

type PageBlockTable struct {
	Flags    int32
	Bordered bool // flag
	Striped  bool // flag
	Title    TL   // RichText
	Rows     []TL // PageTableRow
}

type PageBlockOrderedList struct {
	Items []TL // PageListOrderedItem
}

type PageBlockDetails struct {
	Flags  int32
	Open   bool // flag
	Blocks []TL // PageBlock
	Title  TL   // RichText
}

type PageBlockRelatedArticles struct {
	Title    TL   // RichText
	Articles []TL // PageRelatedArticle
}

type PageBlockMap struct {
	Geo     TL // GeoPoint
	Zoom    int32
	W       int32
	H       int32
	Caption TL // PageCaption
}

type PhoneCallDiscardReasonMissed struct {
}

type PhoneCallDiscardReasonDisconnect struct {
}

type PhoneCallDiscardReasonHangup struct {
}

type PhoneCallDiscardReasonBusy struct {
}

type DataJSON struct {
	Data string
}

type LabeledPrice struct {
	Label  string
	Amount int64
}

type Invoice struct {
	Flags                    int32
	Test                     bool // flag
	NameRequested            bool // flag
	PhoneRequested           bool // flag
	EmailRequested           bool // flag
	ShippingAddressRequested bool // flag
	Flexible                 bool // flag
	PhoneToProvider          bool // flag
	EmailToProvider          bool // flag
	Currency                 string
	Prices                   []TL // LabeledPrice
}

type PaymentCharge struct {
	ID               string
	ProviderChargeID string
}

type PostAddress struct {
	StreetLine1 string
	StreetLine2 string
	City        string
	State       string
	CountryIso2 string
	PostCode    string
}

type PaymentRequestedInfo struct {
	Flags           int32
	Name            string // flag
	Phone           string // flag
	Email           string // flag
	ShippingAddress TL     // PostAddress // flag
}

type PaymentSavedCredentialsCard struct {
	ID    string
	Title string
}

type WebDocument struct {
	Url        string
	AccessHash int64
	Size       int32
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type WebDocumentNoProxy struct {
	Url        string
	Size       int32
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type InputWebDocument struct {
	Url        string
	Size       int32
	MimeType   string
	Attributes []TL // DocumentAttribute
}

type InputWebFileLocation struct {
	Url        string
	AccessHash int64
}

type InputWebFileGeoPointLocation struct {
	GeoPoint   TL // InputGeoPoint
	AccessHash int64
	W          int32
	H          int32
	Zoom       int32
	Scale      int32
}

type UploadWebFile struct {
	Size     int32
	MimeType string
	FileType TL // storage_FileType
	Mtime    int32
	Bytes    []byte
}

type PaymentsPaymentForm struct {
	Flags              int32
	CanSaveCredentials bool // flag
	PasswordMissing    bool // flag
	BotID              int32
	Invoice            TL // Invoice
	ProviderID         int32
	Url                string
	NativeProvider     string // flag
	NativeParams       TL     // DataJSON // flag
	SavedInfo          TL     // PaymentRequestedInfo // flag
	SavedCredentials   TL     // PaymentSavedCredentials // flag
	Users              []TL   // User
}

type PaymentsValidatedRequestedInfo struct {
	Flags           int32
	ID              string // flag
	ShippingOptions []TL   // ShippingOption // flag
}

type PaymentsPaymentResult struct {
	Updates TL // Updates
}

type PaymentsPaymentVerificationNeeded struct {
	Url string
}

type PaymentsPaymentReceipt struct {
	Flags            int32
	Date             int32
	BotID            int32
	Invoice          TL // Invoice
	ProviderID       int32
	Info             TL // PaymentRequestedInfo // flag
	Shipping         TL // ShippingOption // flag
	Currency         string
	TotalAmount      int64
	CredentialsTitle string
	Users            []TL // User
}

type PaymentsSavedInfo struct {
	Flags               int32
	HasSavedCredentials bool // flag
	SavedInfo           TL   // PaymentRequestedInfo // flag
}

type InputPaymentCredentialsSaved struct {
	ID          string
	TmpPassword []byte
}

type InputPaymentCredentials struct {
	Flags int32
	Save  bool // flag
	Data  TL   // DataJSON
}

type InputPaymentCredentialsApplePay struct {
	PaymentData TL // DataJSON
}

type InputPaymentCredentialsAndroidPay struct {
	PaymentToken        TL // DataJSON
	GoogleTransactionID string
}

type AccountTmpPassword struct {
	TmpPassword []byte
	ValidUntil  int32
}

type ShippingOption struct {
	ID     string
	Title  string
	Prices []TL // LabeledPrice
}

type InputStickerSetItem struct {
	Flags      int32
	Document   TL // InputDocument
	Emoji      string
	MaskCoords TL // MaskCoords // flag
}

type InputPhoneCall struct {
	ID         int64
	AccessHash int64
}

type PhoneCallEmpty struct {
	ID int64
}

type PhoneCallWaiting struct {
	Flags         int32
	Video         bool // flag
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	Protocol      TL    // PhoneCallProtocol
	ReceiveDate   int32 // flag
}

type PhoneCallRequested struct {
	Flags         int32
	Video         bool // flag
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	GAHash        []byte
	Protocol      TL // PhoneCallProtocol
}

type PhoneCallAccepted struct {
	Flags         int32
	Video         bool // flag
	ID            int64
	AccessHash    int64
	Date          int32
	AdminID       int32
	ParticipantID int32
	GB            []byte
	Protocol      TL // PhoneCallProtocol
}

type PhoneCall struct {
	Flags          int32
	P2pAllowed     bool // flag
	ID             int64
	AccessHash     int64
	Date           int32
	AdminID        int32
	ParticipantID  int32
	GAOrB          []byte
	KeyFingerprint int64
	Protocol       TL   // PhoneCallProtocol
	Connections    []TL // PhoneConnection
	StartDate      int32
}

type PhoneCallDiscarded struct {
	Flags      int32
	NeedRating bool // flag
	NeedDebug  bool // flag
	Video      bool // flag
	ID         int64
	Reason     TL    // PhoneCallDiscardReason // flag
	Duration   int32 // flag
}

type PhoneConnection struct {
	ID      int64
	Ip      string
	Ipv6    string
	Port    int32
	PeerTag []byte
}

type PhoneCallProtocol struct {
	Flags           int32
	UdpP2p          bool // flag
	UdpReflector    bool // flag
	MinLayer        int32
	MaxLayer        int32
	LibraryVersions []string
}

type PhonePhoneCall struct {
	PhoneCall TL   // PhoneCall
	Users     []TL // User
}

type UploadCdnFileReuploadNeeded struct {
	RequestToken []byte
}

type UploadCdnFile struct {
	Bytes []byte
}

type CdnPublicKey struct {
	DcID      int32
	PublicKey string
}

type CdnConfig struct {
	PublicKeys []TL // CdnPublicKey
}

type LangPackString struct {
	Key   string
	Value string
}

type LangPackStringPluralized struct {
	Flags      int32
	Key        string
	ZeroValue  string // flag
	OneValue   string // flag
	TwoValue   string // flag
	FewValue   string // flag
	ManyValue  string // flag
	OtherValue string
}

type LangPackStringDeleted struct {
	Key string
}

type LangPackDifference struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Strings     []TL // LangPackString
}

type LangPackLanguage struct {
	Flags           int32
	Official        bool // flag
	Rtl             bool // flag
	Beta            bool // flag
	Name            string
	NativeName      string
	LangCode        string
	BaseLangCode    string // flag
	PluralCode      string
	StringsCount    int32
	TranslatedCount int32
	TranslationsUrl string
}

type ChannelAdminLogEventActionChangeTitle struct {
	PrevValue string
	NewValue  string
}

type ChannelAdminLogEventActionChangeAbout struct {
	PrevValue string
	NewValue  string
}

type ChannelAdminLogEventActionChangeUsername struct {
	PrevValue string
	NewValue  string
}

type ChannelAdminLogEventActionChangePhoto struct {
	PrevPhoto TL // Photo
	NewPhoto  TL // Photo
}

type ChannelAdminLogEventActionToggleInvites struct {
	NewValue TL // Bool
}

type ChannelAdminLogEventActionToggleSignatures struct {
	NewValue TL // Bool
}

type ChannelAdminLogEventActionUpdatePinned struct {
	Message TL // Message
}

type ChannelAdminLogEventActionEditMessage struct {
	PrevMessage TL // Message
	NewMessage  TL // Message
}

type ChannelAdminLogEventActionDeleteMessage struct {
	Message TL // Message
}

type ChannelAdminLogEventActionParticipantJoin struct {
}

type ChannelAdminLogEventActionParticipantLeave struct {
}

type ChannelAdminLogEventActionParticipantInvite struct {
	Participant TL // ChannelParticipant
}

type ChannelAdminLogEventActionParticipantToggleBan struct {
	PrevParticipant TL // ChannelParticipant
	NewParticipant  TL // ChannelParticipant
}

type ChannelAdminLogEventActionParticipantToggleAdmin struct {
	PrevParticipant TL // ChannelParticipant
	NewParticipant  TL // ChannelParticipant
}

type ChannelAdminLogEventActionChangeStickerSet struct {
	PrevStickerset TL // InputStickerSet
	NewStickerset  TL // InputStickerSet
}

type ChannelAdminLogEventActionTogglePreHistoryHidden struct {
	NewValue TL // Bool
}

type ChannelAdminLogEventActionDefaultBannedRights struct {
	PrevBannedRights TL // ChatBannedRights
	NewBannedRights  TL // ChatBannedRights
}

type ChannelAdminLogEventActionStopPoll struct {
	Message TL // Message
}

type ChannelAdminLogEventActionChangeLinkedChat struct {
	PrevValue int32
	NewValue  int32
}

type ChannelAdminLogEventActionChangeLocation struct {
	PrevValue TL // ChannelLocation
	NewValue  TL // ChannelLocation
}

type ChannelAdminLogEventActionToggleSlowMode struct {
	PrevValue int32
	NewValue  int32
}

type ChannelAdminLogEvent struct {
	ID     int64
	Date   int32
	UserID int32
	Action TL // ChannelAdminLogEventAction
}

type ChannelsAdminLogResults struct {
	Events []TL // ChannelAdminLogEvent
	Chats  []TL // Chat
	Users  []TL // User
}

type ChannelAdminLogEventsFilter struct {
	Flags    int32
	Join     bool // flag
	Leave    bool // flag
	Invite   bool // flag
	Ban      bool // flag
	Unban    bool // flag
	Kick     bool // flag
	Unkick   bool // flag
	Promote  bool // flag
	Demote   bool // flag
	Info     bool // flag
	Settings bool // flag
	Pinned   bool // flag
	Edit     bool // flag
	Delete   bool // flag
}

type PopularContact struct {
	ClientID  int64
	Importers int32
}

type MessagesFavedStickersNotModified struct {
}

type MessagesFavedStickers struct {
	Hash     int32
	Packs    []TL // StickerPack
	Stickers []TL // Document
}

type RecentMeUrlUnknown struct {
	Url string
}

type RecentMeUrlUser struct {
	Url    string
	UserID int32
}

type RecentMeUrlChat struct {
	Url    string
	ChatID int32
}

type RecentMeUrlChatInvite struct {
	Url        string
	ChatInvite TL // ChatInvite
}

type RecentMeUrlStickerSet struct {
	Url string
	Set TL // StickerSetCovered
}

type HelpRecentMeUrls struct {
	Urls  []TL // RecentMeUrl
	Chats []TL // Chat
	Users []TL // User
}

type InputSingleMedia struct {
	Flags    int32
	Media    TL // InputMedia
	RandomID int64
	Message  string
	Entities []TL // MessageEntity // flag
}

type WebAuthorization struct {
	Hash        int64
	BotID       int32
	Domain      string
	Browser     string
	Platform    string
	DateCreated int32
	DateActive  int32
	Ip          string
	Region      string
}

type AccountWebAuthorizations struct {
	Authorizations []TL // WebAuthorization
	Users          []TL // User
}

type InputMessageID struct {
	ID int32
}

type InputMessageReplyTo struct {
	ID int32
}

type InputMessagePinned struct {
}

type InputDialogPeer struct {
	Peer TL // InputPeer
}

type InputDialogPeerFolder struct {
	FolderID int32
}

type DialogPeer struct {
	Peer TL // Peer
}

type DialogPeerFolder struct {
	FolderID int32
}

type MessagesFoundStickerSetsNotModified struct {
}

type MessagesFoundStickerSets struct {
	Hash int32
	Sets []TL // StickerSetCovered
}

type FileHash struct {
	Offset int32
	Limit  int32
	Hash   []byte
}

type InputClientProxy struct {
	Address string
	Port    int32
}

type HelpTermsOfServiceUpdateEmpty struct {
	Expires int32
}

type HelpTermsOfServiceUpdate struct {
	Expires        int32
	TermsOfService TL // help_TermsOfService
}

type InputSecureFileUploaded struct {
	ID          int64
	Parts       int32
	Md5Checksum string
	FileHash    []byte
	Secret      []byte
}

type InputSecureFile struct {
	ID         int64
	AccessHash int64
}

type SecureFileEmpty struct {
}

type SecureFile struct {
	ID         int64
	AccessHash int64
	Size       int32
	DcID       int32
	Date       int32
	FileHash   []byte
	Secret     []byte
}

type SecureData struct {
	Data     []byte
	DataHash []byte
	Secret   []byte
}

type SecurePlainPhone struct {
	Phone string
}

type SecurePlainEmail struct {
	Email string
}

type SecureValueTypePersonalDetails struct {
}

type SecureValueTypePassport struct {
}

type SecureValueTypeDriverLicense struct {
}

type SecureValueTypeIdentityCard struct {
}

type SecureValueTypeInternalPassport struct {
}

type SecureValueTypeAddress struct {
}

type SecureValueTypeUtilityBill struct {
}

type SecureValueTypeBankStatement struct {
}

type SecureValueTypeRentalAgreement struct {
}

type SecureValueTypePassportRegistration struct {
}

type SecureValueTypeTemporaryRegistration struct {
}

type SecureValueTypePhone struct {
}

type SecureValueTypeEmail struct {
}

type SecureValue struct {
	Flags       int32
	Type        TL   // SecureValueType
	Data        TL   // SecureData // flag
	FrontSide   TL   // SecureFile // flag
	ReverseSide TL   // SecureFile // flag
	Selfie      TL   // SecureFile // flag
	Translation []TL // SecureFile // flag
	Files       []TL // SecureFile // flag
	PlainData   TL   // SecurePlainData // flag
	Hash        []byte
}

type InputSecureValue struct {
	Flags       int32
	Type        TL   // SecureValueType
	Data        TL   // SecureData // flag
	FrontSide   TL   // InputSecureFile // flag
	ReverseSide TL   // InputSecureFile // flag
	Selfie      TL   // InputSecureFile // flag
	Translation []TL // InputSecureFile // flag
	Files       []TL // InputSecureFile // flag
	PlainData   TL   // SecurePlainData // flag
}

type SecureValueHash struct {
	Type TL // SecureValueType
	Hash []byte
}

type SecureValueErrorData struct {
	Type     TL // SecureValueType
	DataHash []byte
	Field    string
	Text     string
}

type SecureValueErrorFrontSide struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type SecureValueErrorReverseSide struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type SecureValueErrorSelfie struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type SecureValueErrorFile struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type SecureValueErrorFiles struct {
	Type     TL   // SecureValueType
	FileHash []TL // Bytes
	Text     string
}

type SecureValueError struct {
	Type TL // SecureValueType
	Hash []byte
	Text string
}

type SecureValueErrorTranslationFile struct {
	Type     TL // SecureValueType
	FileHash []byte
	Text     string
}

type SecureValueErrorTranslationFiles struct {
	Type     TL   // SecureValueType
	FileHash []TL // Bytes
	Text     string
}

type SecureCredentialsEncrypted struct {
	Data   []byte
	Hash   []byte
	Secret []byte
}

type AccountAuthorizationForm struct {
	Flags            int32
	RequiredTypes    []TL   // SecureRequiredType
	Values           []TL   // SecureValue
	Errors           []TL   // SecureValueError
	Users            []TL   // User
	PrivacyPolicyUrl string // flag
}

type AccountSentEmailCode struct {
	EmailPattern string
	Length       int32
}

type HelpDeepLinkInfoEmpty struct {
}

type HelpDeepLinkInfo struct {
	Flags     int32
	UpdateApp bool // flag
	Message   string
	Entities  []TL // MessageEntity // flag
}

type SavedPhoneContact struct {
	Phone     string
	FirstName string
	LastName  string
	Date      int32
}

type AccountTakeout struct {
	ID int64
}

type PasswordKdfAlgoUnknown struct {
}

type PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow struct {
	Salt1 []byte
	Salt2 []byte
	G     int32
	P     []byte
}

type SecurePasswordKdfAlgoUnknown struct {
}

type SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000 struct {
	Salt []byte
}

type SecurePasswordKdfAlgoSHA512 struct {
	Salt []byte
}

type SecureSecretSettings struct {
	SecureAlgo     TL // SecurePasswordKdfAlgo
	SecureSecret   []byte
	SecureSecretID int64
}

type InputCheckPasswordEmpty struct {
}

type InputCheckPasswordSRP struct {
	SrpID int64
	A     []byte
	M1    []byte
}

type SecureRequiredType struct {
	Flags               int32
	NativeNames         bool // flag
	SelfieRequired      bool // flag
	TranslationRequired bool // flag
	Type                TL   // SecureValueType
}

type SecureRequiredTypeOneOf struct {
	Types []TL // SecureRequiredType
}

type HelpPassportConfigNotModified struct {
}

type HelpPassportConfig struct {
	Hash           int32
	CountriesLangs TL // DataJSON
}

type InputAppEvent struct {
	Time float64
	Type string
	Peer int64
	Data TL // JSONValue
}

type JsonObjectValue struct {
	Key   string
	Value TL // JSONValue
}

type JsonNull struct {
}

type JsonBool struct {
	Value TL // Bool
}

type JsonNumber struct {
	Value float64
}

type JsonString struct {
	Value string
}

type JsonArray struct {
	Value []TL // JSONValue
}

type JsonObject struct {
	Value []TL // JSONObjectValue
}

type PageTableCell struct {
	Flags        int32
	Header       bool  // flag
	AlignCenter  bool  // flag
	AlignRight   bool  // flag
	ValignMiddle bool  // flag
	ValignBottom bool  // flag
	Text         TL    // RichText // flag
	Colspan      int32 // flag
	Rowspan      int32 // flag
}

type PageTableRow struct {
	Cells []TL // PageTableCell
}

type PageCaption struct {
	Text   TL // RichText
	Credit TL // RichText
}

type PageListItemText struct {
	Text TL // RichText
}

type PageListItemBlocks struct {
	Blocks []TL // PageBlock
}

type PageListOrderedItemText struct {
	Num  string
	Text TL // RichText
}

type PageListOrderedItemBlocks struct {
	Num    string
	Blocks []TL // PageBlock
}

type PageRelatedArticle struct {
	Flags         int32
	Url           string
	WebpageID     int64
	Title         string // flag
	Description   string // flag
	PhotoID       int64  // flag
	Author        string // flag
	PublishedDate int32  // flag
}

type Page struct {
	Flags     int32
	Part      bool // flag
	Rtl       bool // flag
	V2        bool // flag
	Url       string
	Blocks    []TL  // PageBlock
	Photos    []TL  // Photo
	Documents []TL  // Document
	Views     int32 // flag
}

type HelpSupportName struct {
	Name string
}

type HelpUserInfoEmpty struct {
}

type HelpUserInfo struct {
	Message  string
	Entities []TL // MessageEntity
	Author   string
	Date     int32
}

type PollAnswer struct {
	Text   string
	Option []byte
}

type Poll struct {
	ID             int64
	Flags          int32
	Closed         bool // flag
	PublicVoters   bool // flag
	MultipleChoice bool // flag
	Quiz           bool // flag
	Question       string
	Answers        []TL  // PollAnswer
	ClosePeriod    int32 // flag
	CloseDate      int32 // flag
}

type PollAnswerVoters struct {
	Flags   int32
	Chosen  bool // flag
	Correct bool // flag
	Option  []byte
	Voters  int32
}

type PollResults struct {
	Flags            int32
	Min              bool    // flag
	Results          []TL    // PollAnswerVoters // flag
	TotalVoters      int32   // flag
	RecentVoters     []int32 // flag
	Solution         string  // flag
	SolutionEntities []TL    // MessageEntity // flag
}

type ChatOnlines struct {
	Onlines int32
}

type StatsURL struct {
	Url string
}

type ChatAdminRights struct {
	Flags          int32
	ChangeInfo     bool // flag
	PostMessages   bool // flag
	EditMessages   bool // flag
	DeleteMessages bool // flag
	BanUsers       bool // flag
	InviteUsers    bool // flag
	PinMessages    bool // flag
	AddAdmins      bool // flag
}

type ChatBannedRights struct {
	Flags        int32
	ViewMessages bool // flag
	SendMessages bool // flag
	SendMedia    bool // flag
	SendStickers bool // flag
	SendGifs     bool // flag
	SendGames    bool // flag
	SendInline   bool // flag
	EmbedLinks   bool // flag
	SendPolls    bool // flag
	ChangeInfo   bool // flag
	InviteUsers  bool // flag
	PinMessages  bool // flag
	UntilDate    int32
}

type InputWallPaper struct {
	ID         int64
	AccessHash int64
}

type InputWallPaperSlug struct {
	Slug string
}

type InputWallPaperNoFile struct {
}

type AccountWallPapersNotModified struct {
}

type AccountWallPapers struct {
	Hash       int32
	Wallpapers []TL // WallPaper
}

type CodeSettings struct {
	Flags          int32
	AllowFlashcall bool // flag
	CurrentNumber  bool // flag
	AllowAppHash   bool // flag
}

type WallPaperSettings struct {
	Flags                 int32
	Blur                  bool  // flag
	Motion                bool  // flag
	BackgroundColor       int32 // flag
	SecondBackgroundColor int32 // flag
	Intensity             int32 // flag
	Rotation              int32 // flag
}

type AutoDownloadSettings struct {
	Flags                 int32
	Disabled              bool // flag
	VideoPreloadLarge     bool // flag
	AudioPreloadNext      bool // flag
	PhonecallsLessData    bool // flag
	PhotoSizeMax          int32
	VideoSizeMax          int32
	FileSizeMax           int32
	VideoUploadMaxbitrate int32
}

type AccountAutoDownloadSettings struct {
	Low    TL // AutoDownloadSettings
	Medium TL // AutoDownloadSettings
	High   TL // AutoDownloadSettings
}

type EmojiKeyword struct {
	Keyword   string
	Emoticons []string
}

type EmojiKeywordDeleted struct {
	Keyword   string
	Emoticons []string
}

type EmojiKeywordsDifference struct {
	LangCode    string
	FromVersion int32
	Version     int32
	Keywords    []TL // EmojiKeyword
}

type EmojiURL struct {
	Url string
}

type EmojiLanguage struct {
	LangCode string
}

type FileLocationToBeDeprecated struct {
	VolumeID int64
	LocalID  int32
}

type Folder struct {
	Flags                     int32
	AutofillNewBroadcasts     bool // flag
	AutofillPublicGroups      bool // flag
	AutofillNewCorrespondents bool // flag
	ID                        int32
	Title                     string
	Photo                     TL // ChatPhoto // flag
}

type InputFolderPeer struct {
	Peer     TL // InputPeer
	FolderID int32
}

type FolderPeer struct {
	Peer     TL // Peer
	FolderID int32
}

type MessagesSearchCounter struct {
	Flags   int32
	Inexact bool // flag
	Filter  TL   // MessagesFilter
	Count   int32
}

type UrlAuthResultRequest struct {
	Flags              int32
	RequestWriteAccess bool // flag
	Bot                TL   // User
	Domain             string
}

type UrlAuthResultAccepted struct {
	Url string
}

type UrlAuthResultDefault struct {
}

type ChannelLocationEmpty struct {
}

type ChannelLocation struct {
	GeoPoint TL // GeoPoint
	Address  string
}

type PeerLocated struct {
	Peer     TL // Peer
	Expires  int32
	Distance int32
}

type PeerSelfLocated struct {
	Expires int32
}

type RestrictionReason struct {
	Platform string
	Reason   string
	Text     string
}

type InputTheme struct {
	ID         int64
	AccessHash int64
}

type InputThemeSlug struct {
	Slug string
}

type Theme struct {
	Flags         int32
	Creator       bool // flag
	Default       bool // flag
	ID            int64
	AccessHash    int64
	Slug          string
	Title         string
	Document      TL // Document // flag
	Settings      TL // ThemeSettings // flag
	InstallsCount int32
}

type AccountThemesNotModified struct {
}

type AccountThemes struct {
	Hash   int32
	Themes []TL // Theme
}

type AuthLoginToken struct {
	Expires int32
	Token   []byte
}

type AuthLoginTokenMigrateTo struct {
	DcID  int32
	Token []byte
}

type AuthLoginTokenSuccess struct {
	Authorization TL // auth_Authorization
}

type AccountContentSettings struct {
	Flags              int32
	SensitiveEnabled   bool // flag
	SensitiveCanChange bool // flag
}

type MessagesInactiveChats struct {
	Dates []int32
	Chats []TL // Chat
	Users []TL // User
}

type BaseThemeClassic struct {
}

type BaseThemeDay struct {
}

type BaseThemeNight struct {
}

type BaseThemeTinted struct {
}

type BaseThemeArctic struct {
}

type InputThemeSettings struct {
	Flags              int32
	BaseTheme          TL // BaseTheme
	AccentColor        int32
	MessageTopColor    int32 // flag
	MessageBottomColor int32 // flag
	Wallpaper          TL    // InputWallPaper // flag
	WallpaperSettings  TL    // WallPaperSettings // flag
}

type ThemeSettings struct {
	Flags              int32
	BaseTheme          TL // BaseTheme
	AccentColor        int32
	MessageTopColor    int32 // flag
	MessageBottomColor int32 // flag
	Wallpaper          TL    // WallPaper // flag
}

type WebPageAttributeTheme struct {
	Flags     int32
	Documents []TL // Document // flag
	Settings  TL   // ThemeSettings // flag
}

type MessageUserVote struct {
	UserID int32
	Option []byte
	Date   int32
}

type MessageUserVoteInputOption struct {
	UserID int32
	Date   int32
}

type MessageUserVoteMultiple struct {
	UserID  int32
	Options []TL // Bytes
	Date    int32
}

type MessagesVotesList struct {
	Flags      int32
	Count      int32
	Votes      []TL   // MessageUserVote
	Users      []TL   // User
	NextOffset string // flag
}

type BankCardOpenUrl struct {
	Url  string
	Name string
}

type PaymentsBankCardData struct {
	Title    string
	OpenUrls []TL // BankCardOpenUrl
}

type DialogFilter struct {
	Flags           int32
	Contacts        bool // flag
	NonContacts     bool // flag
	Groups          bool // flag
	Broadcasts      bool // flag
	Bots            bool // flag
	ExcludeMuted    bool // flag
	ExcludeRead     bool // flag
	ExcludeArchived bool // flag
	ID              int32
	Title           string
	Emoticon        string // flag
	PinnedPeers     []TL   // InputPeer
	IncludePeers    []TL   // InputPeer
	ExcludePeers    []TL   // InputPeer
}

type DialogFilterSuggested struct {
	Filter      TL // DialogFilter
	Description string
}

type StatsDateRangeDays struct {
	MinDate int32
	MaxDate int32
}

type StatsAbsValueAndPrev struct {
	Current  float64
	Previous float64
}

type StatsPercentValue struct {
	Part  float64
	Total float64
}

type StatsGraphAsync struct {
	Token string
}

type StatsGraphError struct {
	Error string
}

type StatsGraph struct {
	Flags     int32
	Json      TL     // DataJSON
	ZoomToken string // flag
}

type MessageInteractionCounters struct {
	MsgID    int32
	Views    int32
	Forwards int32
}

type StatsBroadcastStats struct {
	Period                    TL   // StatsDateRangeDays
	Followers                 TL   // StatsAbsValueAndPrev
	ViewsPerPost              TL   // StatsAbsValueAndPrev
	SharesPerPost             TL   // StatsAbsValueAndPrev
	EnabledNotifications      TL   // StatsPercentValue
	GrowthGraph               TL   // StatsGraph
	FollowersGraph            TL   // StatsGraph
	MuteGraph                 TL   // StatsGraph
	TopHoursGraph             TL   // StatsGraph
	InteractionsGraph         TL   // StatsGraph
	IvInteractionsGraph       TL   // StatsGraph
	ViewsBySourceGraph        TL   // StatsGraph
	NewFollowersBySourceGraph TL   // StatsGraph
	LanguagesGraph            TL   // StatsGraph
	RecentMessageInteractions []TL // MessageInteractionCounters
}

type HelpPromoDataEmpty struct {
	Expires int32
}

type HelpPromoData struct {
	Flags      int32
	Proxy      bool // flag
	Expires    int32
	Peer       TL     // Peer
	Chats      []TL   // Chat
	Users      []TL   // User
	PsaType    string // flag
	PsaMessage string // flag
}

type VideoSize struct {
	Flags        int32
	Type         string
	Location     TL // FileLocation
	W            int32
	H            int32
	Size         int32
	VideoStartTs float64 // flag
}

type StatsGroupTopPoster struct {
	UserID   int32
	Messages int32
	AvgChars int32
}

type StatsGroupTopAdmin struct {
	UserID  int32
	Deleted int32
	Kicked  int32
	Banned  int32
}

type StatsGroupTopInviter struct {
	UserID      int32
	Invitations int32
}

type StatsMegagroupStats struct {
	Period                  TL   // StatsDateRangeDays
	Members                 TL   // StatsAbsValueAndPrev
	Messages                TL   // StatsAbsValueAndPrev
	Viewers                 TL   // StatsAbsValueAndPrev
	Posters                 TL   // StatsAbsValueAndPrev
	GrowthGraph             TL   // StatsGraph
	MembersGraph            TL   // StatsGraph
	NewMembersBySourceGraph TL   // StatsGraph
	LanguagesGraph          TL   // StatsGraph
	MessagesGraph           TL   // StatsGraph
	ActionsGraph            TL   // StatsGraph
	TopHoursGraph           TL   // StatsGraph
	WeekdaysGraph           TL   // StatsGraph
	TopPosters              []TL // StatsGroupTopPoster
	TopAdmins               []TL // StatsGroupTopAdmin
	TopInviters             []TL // StatsGroupTopInviter
	Users                   []TL // User
}

type GlobalPrivacySettings struct {
	Flags                            int32
	ArchiveAndMuteNewNoncontactPeers TL // Bool // flag
}

type InvokeAfterMsg struct {
	MsgID int64
	Query TL
}

type InvokeAfterMsgs struct {
	MsgIds []int64
	Query  TL
}

type InitConnection struct {
	Flags          int32
	ApiID          int32
	DeviceModel    string
	SystemVersion  string
	AppVersion     string
	SystemLangCode string
	LangPack       string
	LangCode       string
	Proxy          TL // InputClientProxy // flag
	Params         TL // JSONValue // flag
	Query          TL
}

type InvokeWithLayer struct {
	Layer int32
	Query TL
}

type InvokeWithoutUpdates struct {
	Query TL
}

type InvokeWithMessagesRange struct {
	Range TL // MessageRange
	Query TL
}

type InvokeWithTakeout struct {
	TakeoutID int64
	Query     TL
}

type AuthSendCode struct {
	PhoneNumber string
	ApiID       int32
	ApiHash     string
	Settings    TL // CodeSettings
}

type AuthSignUp struct {
	PhoneNumber   string
	PhoneCodeHash string
	FirstName     string
	LastName      string
}

type AuthSignIn struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type AuthLogOut struct {
}

type AuthResetAuthorizations struct {
}

type AuthExportAuthorization struct {
	DcID int32
}

type AuthImportAuthorization struct {
	ID    int32
	Bytes []byte
}

type AuthBindTempAuthKey struct {
	PermAuthKeyID    int64
	Nonce            int64
	ExpiresAt        int32
	EncryptedMessage []byte
}

type AuthImportBotAuthorization struct {
	Flags        int32
	ApiID        int32
	ApiHash      string
	BotAuthToken string
}

type AuthCheckPassword struct {
	Password TL // InputCheckPasswordSRP
}

type AuthRequestPasswordRecovery struct {
}

type AuthRecoverPassword struct {
	Code string
}

type AuthResendCode struct {
	PhoneNumber   string
	PhoneCodeHash string
}

type AuthCancelCode struct {
	PhoneNumber   string
	PhoneCodeHash string
}

type AuthDropTempAuthKeys struct {
	ExceptAuthKeys []int64
}

type AuthExportLoginToken struct {
	ApiID     int32
	ApiHash   string
	ExceptIds []int32
}

type AuthImportLoginToken struct {
	Token []byte
}

type AuthAcceptLoginToken struct {
	Token []byte
}

type AccountRegisterDevice struct {
	Flags      int32
	NoMuted    bool // flag
	TokenType  int32
	Token      string
	AppSandbox TL // Bool
	Secret     []byte
	OtherUids  []int32
}

type AccountUnregisterDevice struct {
	TokenType int32
	Token     string
	OtherUids []int32
}

type AccountUpdateNotifySettings struct {
	Peer     TL // InputNotifyPeer
	Settings TL // InputPeerNotifySettings
}

type AccountGetNotifySettings struct {
	Peer TL // InputNotifyPeer
}

type AccountResetNotifySettings struct {
}

type AccountUpdateProfile struct {
	Flags     int32
	FirstName string // flag
	LastName  string // flag
	About     string // flag
}

type AccountUpdateStatus struct {
	Offline TL // Bool
}

type AccountGetWallPapers struct {
	Hash int32
}

type AccountReportPeer struct {
	Peer   TL // InputPeer
	Reason TL // ReportReason
}

type AccountCheckUsername struct {
	Username string
}

type AccountUpdateUsername struct {
	Username string
}

type AccountGetPrivacy struct {
	Key TL // InputPrivacyKey
}

type AccountSetPrivacy struct {
	Key   TL   // InputPrivacyKey
	Rules []TL // InputPrivacyRule
}

type AccountDeleteAccount struct {
	Reason string
}

type AccountGetAccountTTL struct {
}

type AccountSetAccountTTL struct {
	Ttl TL // AccountDaysTTL
}

type AccountSendChangePhoneCode struct {
	PhoneNumber string
	Settings    TL // CodeSettings
}

type AccountChangePhone struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type AccountUpdateDeviceLocked struct {
	Period int32
}

type AccountGetAuthorizations struct {
}

type AccountResetAuthorization struct {
	Hash int64
}

type AccountGetPassword struct {
}

type AccountGetPasswordSettings struct {
	Password TL // InputCheckPasswordSRP
}

type AccountUpdatePasswordSettings struct {
	Password    TL // InputCheckPasswordSRP
	NewSettings TL // account_PasswordInputSettings
}

type AccountSendConfirmPhoneCode struct {
	Hash     string
	Settings TL // CodeSettings
}

type AccountConfirmPhone struct {
	PhoneCodeHash string
	PhoneCode     string
}

type AccountGetTmpPassword struct {
	Password TL // InputCheckPasswordSRP
	Period   int32
}

type AccountGetWebAuthorizations struct {
}

type AccountResetWebAuthorization struct {
	Hash int64
}

type AccountResetWebAuthorizations struct {
}

type AccountGetAllSecureValues struct {
}

type AccountGetSecureValue struct {
	Types []TL // SecureValueType
}

type AccountSaveSecureValue struct {
	Value          TL // InputSecureValue
	SecureSecretID int64
}

type AccountDeleteSecureValue struct {
	Types []TL // SecureValueType
}

type AccountGetAuthorizationForm struct {
	BotID     int32
	Scope     string
	PublicKey string
}

type AccountAcceptAuthorization struct {
	BotID       int32
	Scope       string
	PublicKey   string
	ValueHashes []TL // SecureValueHash
	Credentials TL   // SecureCredentialsEncrypted
}

type AccountSendVerifyPhoneCode struct {
	PhoneNumber string
	Settings    TL // CodeSettings
}

type AccountVerifyPhone struct {
	PhoneNumber   string
	PhoneCodeHash string
	PhoneCode     string
}

type AccountSendVerifyEmailCode struct {
	Email string
}

type AccountVerifyEmail struct {
	Email string
	Code  string
}

type AccountInitTakeoutSession struct {
	Flags             int32
	Contacts          bool  // flag
	MessageUsers      bool  // flag
	MessageChats      bool  // flag
	MessageMegagroups bool  // flag
	MessageChannels   bool  // flag
	Files             bool  // flag
	FileMaxSize       int32 // flag
}

type AccountFinishTakeoutSession struct {
	Flags   int32
	Success bool // flag
}

type AccountConfirmPasswordEmail struct {
	Code string
}

type AccountResendPasswordEmail struct {
}

type AccountCancelPasswordEmail struct {
}

type AccountGetContactSignUpNotification struct {
}

type AccountSetContactSignUpNotification struct {
	Silent TL // Bool
}

type AccountGetNotifyExceptions struct {
	Flags        int32
	CompareSound bool // flag
	Peer         TL   // InputNotifyPeer // flag
}

type AccountGetWallPaper struct {
	Wallpaper TL // InputWallPaper
}

type AccountUploadWallPaper struct {
	File     TL // InputFile
	MimeType string
	Settings TL // WallPaperSettings
}

type AccountSaveWallPaper struct {
	Wallpaper TL // InputWallPaper
	Unsave    TL // Bool
	Settings  TL // WallPaperSettings
}

type AccountInstallWallPaper struct {
	Wallpaper TL // InputWallPaper
	Settings  TL // WallPaperSettings
}

type AccountResetWallPapers struct {
}

type AccountGetAutoDownloadSettings struct {
}

type AccountSaveAutoDownloadSettings struct {
	Flags    int32
	Low      bool // flag
	High     bool // flag
	Settings TL   // AutoDownloadSettings
}

type AccountUploadTheme struct {
	Flags    int32
	File     TL // InputFile
	Thumb    TL // InputFile // flag
	FileName string
	MimeType string
}

type AccountCreateTheme struct {
	Flags    int32
	Slug     string
	Title    string
	Document TL // InputDocument // flag
	Settings TL // InputThemeSettings // flag
}

type AccountUpdateTheme struct {
	Flags    int32
	Format   string
	Theme    TL     // InputTheme
	Slug     string // flag
	Title    string // flag
	Document TL     // InputDocument // flag
	Settings TL     // InputThemeSettings // flag
}

type AccountSaveTheme struct {
	Theme  TL // InputTheme
	Unsave TL // Bool
}

type AccountInstallTheme struct {
	Flags  int32
	Dark   bool   // flag
	Format string // flag
	Theme  TL     // InputTheme // flag
}

type AccountGetTheme struct {
	Format     string
	Theme      TL // InputTheme
	DocumentID int64
}

type AccountGetThemes struct {
	Format string
	Hash   int32
}

type AccountSetContentSettings struct {
	Flags            int32
	SensitiveEnabled bool // flag
}

type AccountGetContentSettings struct {
}

type AccountGetMultiWallPapers struct {
	Wallpapers []TL // InputWallPaper
}

type AccountGetGlobalPrivacySettings struct {
}

type AccountSetGlobalPrivacySettings struct {
	Settings TL // GlobalPrivacySettings
}

type UsersGetUsers struct {
	ID []TL // InputUser
}

type UsersGetFullUser struct {
	ID TL // InputUser
}

type UsersSetSecureValueErrors struct {
	ID     TL   // InputUser
	Errors []TL // SecureValueError
}

type ContactsGetContactIDs struct {
	Hash int32
}

type ContactsGetStatuses struct {
}

type ContactsGetContacts struct {
	Hash int32
}

type ContactsImportContacts struct {
	Contacts []TL // InputContact
}

type ContactsDeleteContacts struct {
	ID []TL // InputUser
}

type ContactsDeleteByPhones struct {
	Phones []string
}

type ContactsBlock struct {
	ID TL // InputUser
}

type ContactsUnblock struct {
	ID TL // InputUser
}

type ContactsGetBlocked struct {
	Offset int32
	Limit  int32
}

type ContactsSearch struct {
	Q     string
	Limit int32
}

type ContactsResolveUsername struct {
	Username string
}

type ContactsGetTopPeers struct {
	Flags          int32
	Correspondents bool // flag
	BotsPm         bool // flag
	BotsInline     bool // flag
	PhoneCalls     bool // flag
	ForwardUsers   bool // flag
	ForwardChats   bool // flag
	Groups         bool // flag
	Channels       bool // flag
	Offset         int32
	Limit          int32
	Hash           int32
}

type ContactsResetTopPeerRating struct {
	Category TL // TopPeerCategory
	Peer     TL // InputPeer
}

type ContactsResetSaved struct {
}

type ContactsGetSaved struct {
}

type ContactsToggleTopPeers struct {
	Enabled TL // Bool
}

type ContactsAddContact struct {
	Flags                    int32
	AddPhonePrivacyException bool // flag
	ID                       TL   // InputUser
	FirstName                string
	LastName                 string
	Phone                    string
}

type ContactsAcceptContact struct {
	ID TL // InputUser
}

type ContactsGetLocated struct {
	Flags       int32
	Background  bool  // flag
	GeoPoint    TL    // InputGeoPoint
	SelfExpires int32 // flag
}

type MessagesGetMessages struct {
	ID []TL // InputMessage
}

type MessagesGetDialogs struct {
	Flags         int32
	ExcludePinned bool  // flag
	FolderID      int32 // flag
	OffsetDate    int32
	OffsetID      int32
	OffsetPeer    TL // InputPeer
	Limit         int32
	Hash          int32
}

type MessagesGetHistory struct {
	Peer       TL // InputPeer
	OffsetID   int32
	OffsetDate int32
	AddOffset  int32
	Limit      int32
	MaxID      int32
	MinID      int32
	Hash       int32
}

type MessagesSearch struct {
	Flags     int32
	Peer      TL // InputPeer
	Q         string
	FromID    TL // InputUser // flag
	Filter    TL // MessagesFilter
	MinDate   int32
	MaxDate   int32
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
	Hash      int32
}

type MessagesReadHistory struct {
	Peer  TL // InputPeer
	MaxID int32
}

type MessagesDeleteHistory struct {
	Flags     int32
	JustClear bool // flag
	Revoke    bool // flag
	Peer      TL   // InputPeer
	MaxID     int32
}

type MessagesDeleteMessages struct {
	Flags  int32
	Revoke bool // flag
	ID     []int32
}

type MessagesReceivedMessages struct {
	MaxID int32
}

type MessagesSetTyping struct {
	Peer   TL // InputPeer
	Action TL // SendMessageAction
}

type MessagesSendMessage struct {
	Flags        int32
	NoWebpage    bool  // flag
	Silent       bool  // flag
	Background   bool  // flag
	ClearDraft   bool  // flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 // flag
	Message      string
	RandomID     int64
	ReplyMarkup  TL    // ReplyMarkup // flag
	Entities     []TL  // MessageEntity // flag
	ScheduleDate int32 // flag
}

type MessagesSendMedia struct {
	Flags        int32
	Silent       bool  // flag
	Background   bool  // flag
	ClearDraft   bool  // flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 // flag
	Media        TL    // InputMedia
	Message      string
	RandomID     int64
	ReplyMarkup  TL    // ReplyMarkup // flag
	Entities     []TL  // MessageEntity // flag
	ScheduleDate int32 // flag
}

type MessagesForwardMessages struct {
	Flags        int32
	Silent       bool // flag
	Background   bool // flag
	WithMyScore  bool // flag
	Grouped      bool // flag
	FromPeer     TL   // InputPeer
	ID           []int32
	RandomID     []int64
	ToPeer       TL    // InputPeer
	ScheduleDate int32 // flag
}

type MessagesReportSpam struct {
	Peer TL // InputPeer
}

type MessagesGetPeerSettings struct {
	Peer TL // InputPeer
}

type MessagesReport struct {
	Peer   TL // InputPeer
	ID     []int32
	Reason TL // ReportReason
}

type MessagesGetChats struct {
	ID []int32
}

type MessagesGetFullChat struct {
	ChatID int32
}

type MessagesEditChatTitle struct {
	ChatID int32
	Title  string
}

type MessagesEditChatPhoto struct {
	ChatID int32
	Photo  TL // InputChatPhoto
}

type MessagesAddChatUser struct {
	ChatID   int32
	UserID   TL // InputUser
	FwdLimit int32
}

type MessagesDeleteChatUser struct {
	ChatID int32
	UserID TL // InputUser
}

type MessagesCreateChat struct {
	Users []TL // InputUser
	Title string
}

type MessagesGetDhConfig struct {
	Version      int32
	RandomLength int32
}

type MessagesRequestEncryption struct {
	UserID   TL // InputUser
	RandomID int32
	GA       []byte
}

type MessagesAcceptEncryption struct {
	Peer           TL // InputEncryptedChat
	GB             []byte
	KeyFingerprint int64
}

type MessagesDiscardEncryption struct {
	ChatID int32
}

type MessagesSetEncryptedTyping struct {
	Peer   TL // InputEncryptedChat
	Typing TL // Bool
}

type MessagesReadEncryptedHistory struct {
	Peer    TL // InputEncryptedChat
	MaxDate int32
}

type MessagesSendEncrypted struct {
	Peer     TL // InputEncryptedChat
	RandomID int64
	Data     []byte
}

type MessagesSendEncryptedFile struct {
	Peer     TL // InputEncryptedChat
	RandomID int64
	Data     []byte
	File     TL // InputEncryptedFile
}

type MessagesSendEncryptedService struct {
	Peer     TL // InputEncryptedChat
	RandomID int64
	Data     []byte
}

type MessagesReceivedQueue struct {
	MaxQts int32
}

type MessagesReportEncryptedSpam struct {
	Peer TL // InputEncryptedChat
}

type MessagesReadMessageContents struct {
	ID []int32
}

type MessagesGetStickers struct {
	Emoticon string
	Hash     int32
}

type MessagesGetAllStickers struct {
	Hash int32
}

type MessagesGetWebPagePreview struct {
	Flags    int32
	Message  string
	Entities []TL // MessageEntity // flag
}

type MessagesExportChatInvite struct {
	Peer TL // InputPeer
}

type MessagesCheckChatInvite struct {
	Hash string
}

type MessagesImportChatInvite struct {
	Hash string
}

type MessagesGetStickerSet struct {
	Stickerset TL // InputStickerSet
}

type MessagesInstallStickerSet struct {
	Stickerset TL // InputStickerSet
	Archived   TL // Bool
}

type MessagesUninstallStickerSet struct {
	Stickerset TL // InputStickerSet
}

type MessagesStartBot struct {
	Bot        TL // InputUser
	Peer       TL // InputPeer
	RandomID   int64
	StartParam string
}

type MessagesGetMessagesViews struct {
	Peer      TL // InputPeer
	ID        []int32
	Increment TL // Bool
}

type MessagesEditChatAdmin struct {
	ChatID  int32
	UserID  TL // InputUser
	IsAdmin TL // Bool
}

type MessagesMigrateChat struct {
	ChatID int32
}

type MessagesSearchGlobal struct {
	Flags      int32
	FolderID   int32 // flag
	Q          string
	OffsetRate int32
	OffsetPeer TL // InputPeer
	OffsetID   int32
	Limit      int32
}

type MessagesReorderStickerSets struct {
	Flags int32
	Masks bool // flag
	Order []int64
}

type MessagesGetDocumentByHash struct {
	Sha256   []byte
	Size     int32
	MimeType string
}

type MessagesGetSavedGifs struct {
	Hash int32
}

type MessagesSaveGif struct {
	ID     TL // InputDocument
	Unsave TL // Bool
}

type MessagesGetInlineBotResults struct {
	Flags    int32
	Bot      TL // InputUser
	Peer     TL // InputPeer
	GeoPoint TL // InputGeoPoint // flag
	Query    string
	Offset   string
}

type MessagesSetInlineBotResults struct {
	Flags      int32
	Gallery    bool // flag
	Private    bool // flag
	QueryID    int64
	Results    []TL // InputBotInlineResult
	CacheTime  int32
	NextOffset string // flag
	SwitchPm   TL     // InlineBotSwitchPM // flag
}

type MessagesSendInlineBotResult struct {
	Flags        int32
	Silent       bool  // flag
	Background   bool  // flag
	ClearDraft   bool  // flag
	HideVia      bool  // flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 // flag
	RandomID     int64
	QueryID      int64
	ID           string
	ScheduleDate int32 // flag
}

type MessagesGetMessageEditData struct {
	Peer TL // InputPeer
	ID   int32
}

type MessagesEditMessage struct {
	Flags        int32
	NoWebpage    bool // flag
	Peer         TL   // InputPeer
	ID           int32
	Message      string // flag
	Media        TL     // InputMedia // flag
	ReplyMarkup  TL     // ReplyMarkup // flag
	Entities     []TL   // MessageEntity // flag
	ScheduleDate int32  // flag
}

type MessagesEditInlineBotMessage struct {
	Flags       int32
	NoWebpage   bool   // flag
	ID          TL     // InputBotInlineMessageID
	Message     string // flag
	Media       TL     // InputMedia // flag
	ReplyMarkup TL     // ReplyMarkup // flag
	Entities    []TL   // MessageEntity // flag
}

type MessagesGetBotCallbackAnswer struct {
	Flags int32
	Game  bool // flag
	Peer  TL   // InputPeer
	MsgID int32
	Data  []byte // flag
}

type MessagesSetBotCallbackAnswer struct {
	Flags     int32
	Alert     bool // flag
	QueryID   int64
	Message   string // flag
	Url       string // flag
	CacheTime int32
}

type MessagesGetPeerDialogs struct {
	Peers []TL // InputDialogPeer
}

type MessagesSaveDraft struct {
	Flags        int32
	NoWebpage    bool  // flag
	ReplyToMsgID int32 // flag
	Peer         TL    // InputPeer
	Message      string
	Entities     []TL // MessageEntity // flag
}

type MessagesGetAllDrafts struct {
}

type MessagesGetFeaturedStickers struct {
	Hash int32
}

type MessagesReadFeaturedStickers struct {
	ID []int64
}

type MessagesGetRecentStickers struct {
	Flags    int32
	Attached bool // flag
	Hash     int32
}

type MessagesSaveRecentSticker struct {
	Flags    int32
	Attached bool // flag
	ID       TL   // InputDocument
	Unsave   TL   // Bool
}

type MessagesClearRecentStickers struct {
	Flags    int32
	Attached bool // flag
}

type MessagesGetArchivedStickers struct {
	Flags    int32
	Masks    bool // flag
	OffsetID int64
	Limit    int32
}

type MessagesGetMaskStickers struct {
	Hash int32
}

type MessagesGetAttachedStickers struct {
	Media TL // InputStickeredMedia
}

type MessagesSetGameScore struct {
	Flags       int32
	EditMessage bool // flag
	Force       bool // flag
	Peer        TL   // InputPeer
	ID          int32
	UserID      TL // InputUser
	Score       int32
}

type MessagesSetInlineGameScore struct {
	Flags       int32
	EditMessage bool // flag
	Force       bool // flag
	ID          TL   // InputBotInlineMessageID
	UserID      TL   // InputUser
	Score       int32
}

type MessagesGetGameHighScores struct {
	Peer   TL // InputPeer
	ID     int32
	UserID TL // InputUser
}

type MessagesGetInlineGameHighScores struct {
	ID     TL // InputBotInlineMessageID
	UserID TL // InputUser
}

type MessagesGetCommonChats struct {
	UserID TL // InputUser
	MaxID  int32
	Limit  int32
}

type MessagesGetAllChats struct {
	ExceptIds []int32
}

type MessagesGetWebPage struct {
	Url  string
	Hash int32
}

type MessagesToggleDialogPin struct {
	Flags  int32
	Pinned bool // flag
	Peer   TL   // InputDialogPeer
}

type MessagesReorderPinnedDialogs struct {
	Flags    int32
	Force    bool // flag
	FolderID int32
	Order    []TL // InputDialogPeer
}

type MessagesGetPinnedDialogs struct {
	FolderID int32
}

type MessagesSetBotShippingResults struct {
	Flags           int32
	QueryID         int64
	Error           string // flag
	ShippingOptions []TL   // ShippingOption // flag
}

type MessagesSetBotPrecheckoutResults struct {
	Flags   int32
	Success bool // flag
	QueryID int64
	Error   string // flag
}

type MessagesUploadMedia struct {
	Peer  TL // InputPeer
	Media TL // InputMedia
}

type MessagesSendScreenshotNotification struct {
	Peer         TL // InputPeer
	ReplyToMsgID int32
	RandomID     int64
}

type MessagesGetFavedStickers struct {
	Hash int32
}

type MessagesFaveSticker struct {
	ID     TL // InputDocument
	Unfave TL // Bool
}

type MessagesGetUnreadMentions struct {
	Peer      TL // InputPeer
	OffsetID  int32
	AddOffset int32
	Limit     int32
	MaxID     int32
	MinID     int32
}

type MessagesReadMentions struct {
	Peer TL // InputPeer
}

type MessagesGetRecentLocations struct {
	Peer  TL // InputPeer
	Limit int32
	Hash  int32
}

type MessagesSendMultiMedia struct {
	Flags        int32
	Silent       bool  // flag
	Background   bool  // flag
	ClearDraft   bool  // flag
	Peer         TL    // InputPeer
	ReplyToMsgID int32 // flag
	MultiMedia   []TL  // InputSingleMedia
	ScheduleDate int32 // flag
}

type MessagesUploadEncryptedFile struct {
	Peer TL // InputEncryptedChat
	File TL // InputEncryptedFile
}

type MessagesSearchStickerSets struct {
	Flags           int32
	ExcludeFeatured bool // flag
	Q               string
	Hash            int32
}

type MessagesGetSplitRanges struct {
}

type MessagesMarkDialogUnread struct {
	Flags  int32
	Unread bool // flag
	Peer   TL   // InputDialogPeer
}

type MessagesGetDialogUnreadMarks struct {
}

type MessagesClearAllDrafts struct {
}

type MessagesUpdatePinnedMessage struct {
	Flags  int32
	Silent bool // flag
	Peer   TL   // InputPeer
	ID     int32
}

type MessagesSendVote struct {
	Peer    TL // InputPeer
	MsgID   int32
	Options []TL // Bytes
}

type MessagesGetPollResults struct {
	Peer  TL // InputPeer
	MsgID int32
}

type MessagesGetOnlines struct {
	Peer TL // InputPeer
}

type MessagesGetStatsURL struct {
	Flags  int32
	Dark   bool // flag
	Peer   TL   // InputPeer
	Params string
}

type MessagesEditChatAbout struct {
	Peer  TL // InputPeer
	About string
}

type MessagesEditChatDefaultBannedRights struct {
	Peer         TL // InputPeer
	BannedRights TL // ChatBannedRights
}

type MessagesGetEmojiKeywords struct {
	LangCode string
}

type MessagesGetEmojiKeywordsDifference struct {
	LangCode    string
	FromVersion int32
}

type MessagesGetEmojiKeywordsLanguages struct {
	LangCodes []string
}

type MessagesGetEmojiURL struct {
	LangCode string
}

type MessagesGetSearchCounters struct {
	Peer    TL   // InputPeer
	Filters []TL // MessagesFilter
}

type MessagesRequestUrlAuth struct {
	Peer     TL // InputPeer
	MsgID    int32
	ButtonID int32
}

type MessagesAcceptUrlAuth struct {
	Flags        int32
	WriteAllowed bool // flag
	Peer         TL   // InputPeer
	MsgID        int32
	ButtonID     int32
}

type MessagesHidePeerSettingsBar struct {
	Peer TL // InputPeer
}

type MessagesGetScheduledHistory struct {
	Peer TL // InputPeer
	Hash int32
}

type MessagesGetScheduledMessages struct {
	Peer TL // InputPeer
	ID   []int32
}

type MessagesSendScheduledMessages struct {
	Peer TL // InputPeer
	ID   []int32
}

type MessagesDeleteScheduledMessages struct {
	Peer TL // InputPeer
	ID   []int32
}

type MessagesGetPollVotes struct {
	Flags  int32
	Peer   TL // InputPeer
	ID     int32
	Option []byte // flag
	Offset string // flag
	Limit  int32
}

type MessagesToggleStickerSets struct {
	Flags       int32
	Uninstall   bool // flag
	Archive     bool // flag
	Unarchive   bool // flag
	Stickersets []TL // InputStickerSet
}

type MessagesGetDialogFilters struct {
}

type MessagesGetSuggestedDialogFilters struct {
}

type MessagesUpdateDialogFilter struct {
	Flags  int32
	ID     int32
	Filter TL // DialogFilter // flag
}

type MessagesUpdateDialogFiltersOrder struct {
	Order []int32
}

type MessagesGetOldFeaturedStickers struct {
	Offset int32
	Limit  int32
	Hash   int32
}

type UpdatesGetState struct {
}

type UpdatesGetDifference struct {
	Flags         int32
	Pts           int32
	PtsTotalLimit int32 // flag
	Date          int32
	Qts           int32
}

type UpdatesGetChannelDifference struct {
	Flags   int32
	Force   bool // flag
	Channel TL   // InputChannel
	Filter  TL   // ChannelMessagesFilter
	Pts     int32
	Limit   int32
}

type PhotosUpdateProfilePhoto struct {
	ID TL // InputPhoto
}

type PhotosUploadProfilePhoto struct {
	Flags        int32
	File         TL      // InputFile // flag
	Video        TL      // InputFile // flag
	VideoStartTs float64 // flag
}

type PhotosDeletePhotos struct {
	ID []TL // InputPhoto
}

type PhotosGetUserPhotos struct {
	UserID TL // InputUser
	Offset int32
	MaxID  int64
	Limit  int32
}

type UploadSaveFilePart struct {
	FileID   int64
	FilePart int32
	Bytes    []byte
}

type UploadGetFile struct {
	Flags        int32
	Precise      bool // flag
	CdnSupported bool // flag
	Location     TL   // InputFileLocation
	Offset       int32
	Limit        int32
}

type UploadSaveBigFilePart struct {
	FileID         int64
	FilePart       int32
	FileTotalParts int32
	Bytes          []byte
}

type UploadGetWebFile struct {
	Location TL // InputWebFileLocation
	Offset   int32
	Limit    int32
}

type UploadGetCdnFile struct {
	FileToken []byte
	Offset    int32
	Limit     int32
}

type UploadReuploadCdnFile struct {
	FileToken    []byte
	RequestToken []byte
}

type UploadGetCdnFileHashes struct {
	FileToken []byte
	Offset    int32
}

type UploadGetFileHashes struct {
	Location TL // InputFileLocation
	Offset   int32
}

type HelpGetConfig struct {
}

type HelpGetNearestDc struct {
}

type HelpGetAppUpdate struct {
	Source string
}

type HelpGetInviteText struct {
}

type HelpGetSupport struct {
}

type HelpGetAppChangelog struct {
	PrevAppVersion string
}

type HelpSetBotUpdatesStatus struct {
	PendingUpdatesCount int32
	Message             string
}

type HelpGetCdnConfig struct {
}

type HelpGetRecentMeUrls struct {
	Referer string
}

type HelpGetTermsOfServiceUpdate struct {
}

type HelpAcceptTermsOfService struct {
	ID TL // DataJSON
}

type HelpGetDeepLinkInfo struct {
	Path string
}

type HelpGetAppConfig struct {
}

type HelpSaveAppLog struct {
	Events []TL // InputAppEvent
}

type HelpGetPassportConfig struct {
	Hash int32
}

type HelpGetSupportName struct {
}

type HelpGetUserInfo struct {
	UserID TL // InputUser
}

type HelpEditUserInfo struct {
	UserID   TL // InputUser
	Message  string
	Entities []TL // MessageEntity
}

type HelpGetPromoData struct {
}

type HelpHidePromoData struct {
	Peer TL // InputPeer
}

type HelpDismissSuggestion struct {
	Suggestion string
}

type ChannelsReadHistory struct {
	Channel TL // InputChannel
	MaxID   int32
}

type ChannelsDeleteMessages struct {
	Channel TL // InputChannel
	ID      []int32
}

type ChannelsDeleteUserHistory struct {
	Channel TL // InputChannel
	UserID  TL // InputUser
}

type ChannelsReportSpam struct {
	Channel TL // InputChannel
	UserID  TL // InputUser
	ID      []int32
}

type ChannelsGetMessages struct {
	Channel TL   // InputChannel
	ID      []TL // InputMessage
}

type ChannelsGetParticipants struct {
	Channel TL // InputChannel
	Filter  TL // ChannelParticipantsFilter
	Offset  int32
	Limit   int32
	Hash    int32
}

type ChannelsGetParticipant struct {
	Channel TL // InputChannel
	UserID  TL // InputUser
}

type ChannelsGetChannels struct {
	ID []TL // InputChannel
}

type ChannelsGetFullChannel struct {
	Channel TL // InputChannel
}

type ChannelsCreateChannel struct {
	Flags     int32
	Broadcast bool // flag
	Megagroup bool // flag
	Title     string
	About     string
	GeoPoint  TL     // InputGeoPoint // flag
	Address   string // flag
}

type ChannelsEditAdmin struct {
	Channel     TL // InputChannel
	UserID      TL // InputUser
	AdminRights TL // ChatAdminRights
	Rank        string
}

type ChannelsEditTitle struct {
	Channel TL // InputChannel
	Title   string
}

type ChannelsEditPhoto struct {
	Channel TL // InputChannel
	Photo   TL // InputChatPhoto
}

type ChannelsCheckUsername struct {
	Channel  TL // InputChannel
	Username string
}

type ChannelsUpdateUsername struct {
	Channel  TL // InputChannel
	Username string
}

type ChannelsJoinChannel struct {
	Channel TL // InputChannel
}

type ChannelsLeaveChannel struct {
	Channel TL // InputChannel
}

type ChannelsInviteToChannel struct {
	Channel TL   // InputChannel
	Users   []TL // InputUser
}

type ChannelsDeleteChannel struct {
	Channel TL // InputChannel
}

type ChannelsExportMessageLink struct {
	Channel TL // InputChannel
	ID      int32
	Grouped TL // Bool
}

type ChannelsToggleSignatures struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type ChannelsGetAdminedPublicChannels struct {
	Flags      int32
	ByLocation bool // flag
	CheckLimit bool // flag
}

type ChannelsEditBanned struct {
	Channel      TL // InputChannel
	UserID       TL // InputUser
	BannedRights TL // ChatBannedRights
}

type ChannelsGetAdminLog struct {
	Flags        int32
	Channel      TL // InputChannel
	Q            string
	EventsFilter TL   // ChannelAdminLogEventsFilter // flag
	Admins       []TL // InputUser // flag
	MaxID        int64
	MinID        int64
	Limit        int32
}

type ChannelsSetStickers struct {
	Channel    TL // InputChannel
	Stickerset TL // InputStickerSet
}

type ChannelsReadMessageContents struct {
	Channel TL // InputChannel
	ID      []int32
}

type ChannelsDeleteHistory struct {
	Channel TL // InputChannel
	MaxID   int32
}

type ChannelsTogglePreHistoryHidden struct {
	Channel TL // InputChannel
	Enabled TL // Bool
}

type ChannelsGetLeftChannels struct {
	Offset int32
}

type ChannelsGetGroupsForDiscussion struct {
}

type ChannelsSetDiscussionGroup struct {
	Broadcast TL // InputChannel
	Group     TL // InputChannel
}

type ChannelsEditCreator struct {
	Channel  TL // InputChannel
	UserID   TL // InputUser
	Password TL // InputCheckPasswordSRP
}

type ChannelsEditLocation struct {
	Channel  TL // InputChannel
	GeoPoint TL // InputGeoPoint
	Address  string
}

type ChannelsToggleSlowMode struct {
	Channel TL // InputChannel
	Seconds int32
}

type ChannelsGetInactiveChannels struct {
}

type BotsSendCustomRequest struct {
	CustomMethod string
	Params       TL // DataJSON
}

type BotsAnswerWebhookJSONQuery struct {
	QueryID int64
	Data    TL // DataJSON
}

type BotsSetBotCommands struct {
	Commands []TL // BotCommand
}

type PaymentsGetPaymentForm struct {
	MsgID int32
}

type PaymentsGetPaymentReceipt struct {
	MsgID int32
}

type PaymentsValidateRequestedInfo struct {
	Flags int32
	Save  bool // flag
	MsgID int32
	Info  TL // PaymentRequestedInfo
}

type PaymentsSendPaymentForm struct {
	Flags            int32
	MsgID            int32
	RequestedInfoID  string // flag
	ShippingOptionID string // flag
	Credentials      TL     // InputPaymentCredentials
}

type PaymentsGetSavedInfo struct {
}

type PaymentsClearSavedInfo struct {
	Flags       int32
	Credentials bool // flag
	Info        bool // flag
}

type PaymentsGetBankCardData struct {
	Number string
}

type StickersCreateStickerSet struct {
	Flags     int32
	Masks     bool // flag
	Animated  bool // flag
	UserID    TL   // InputUser
	Title     string
	ShortName string
	Thumb     TL   // InputDocument // flag
	Stickers  []TL // InputStickerSetItem
}

type StickersRemoveStickerFromSet struct {
	Sticker TL // InputDocument
}

type StickersChangeStickerPosition struct {
	Sticker  TL // InputDocument
	Position int32
}

type StickersAddStickerToSet struct {
	Stickerset TL // InputStickerSet
	Sticker    TL // InputStickerSetItem
}

type StickersSetStickerSetThumb struct {
	Stickerset TL // InputStickerSet
	Thumb      TL // InputDocument
}

type PhoneGetCallConfig struct {
}

type PhoneRequestCall struct {
	Flags    int32
	Video    bool // flag
	UserID   TL   // InputUser
	RandomID int32
	GAHash   []byte
	Protocol TL // PhoneCallProtocol
}

type PhoneAcceptCall struct {
	Peer     TL // InputPhoneCall
	GB       []byte
	Protocol TL // PhoneCallProtocol
}

type PhoneConfirmCall struct {
	Peer           TL // InputPhoneCall
	GA             []byte
	KeyFingerprint int64
	Protocol       TL // PhoneCallProtocol
}

type PhoneReceivedCall struct {
	Peer TL // InputPhoneCall
}

type PhoneDiscardCall struct {
	Flags        int32
	Video        bool // flag
	Peer         TL   // InputPhoneCall
	Duration     int32
	Reason       TL // PhoneCallDiscardReason
	ConnectionID int64
}

type PhoneSetCallRating struct {
	Flags          int32
	UserInitiative bool // flag
	Peer           TL   // InputPhoneCall
	Rating         int32
	Comment        string
}

type PhoneSaveCallDebug struct {
	Peer  TL // InputPhoneCall
	Debug TL // DataJSON
}

type PhoneSendSignalingData struct {
	Peer TL // InputPhoneCall
	Data []byte
}

type LangpackGetLangPack struct {
	LangPack string
	LangCode string
}

type LangpackGetStrings struct {
	LangPack string
	LangCode string
	Keys     []string
}

type LangpackGetDifference struct {
	LangPack    string
	LangCode    string
	FromVersion int32
}

type LangpackGetLanguages struct {
	LangPack string
}

type LangpackGetLanguage struct {
	LangPack string
	LangCode string
}

type FoldersEditPeerFolders struct {
	FolderPeers []TL // InputFolderPeer
}

type FoldersDeleteFolder struct {
	FolderID int32
}

type StatsGetBroadcastStats struct {
	Flags   int32
	Dark    bool // flag
	Channel TL   // InputChannel
}

type StatsLoadAsyncGraph struct {
	Flags int32
	Token string
	X     int64 // flag
}

type StatsGetMegagroupStats struct {
	Flags   int32
	Dark    bool // flag
	Channel TL   // InputChannel
}

func (e ResPQ) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcResPQ)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.Pq)
	x.VectorLong(e.ServerPublicKeyFingerprints)
	return x.buf
}

func (e PQInnerDataDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPQInnerDataDc)
	x.String(e.Pq)
	x.String(e.P)
	x.String(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonce)
	x.Int(e.Dc)
	return x.buf
}

func (e PQInnerDataTempDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPQInnerDataTempDc)
	x.String(e.Pq)
	x.String(e.P)
	x.String(e.Q)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonce)
	x.Int(e.Dc)
	x.Int(e.ExpiresIn)
	return x.buf
}

func (e ServerDHParamsFail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcServerDHParamsFail)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash)
	return x.buf
}

func (e ServerDHParamsOk) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcServerDHParamsOk)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.EncryptedAnswer)
	return x.buf
}

func (e ServerDHInnerData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcServerDHInnerData)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Int(e.G)
	x.String(e.DhPrime)
	x.String(e.GA)
	x.Int(e.ServerTime)
	return x.buf
}

func (e ClientDHInnerData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcClientDHInnerData)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Long(e.RetryID)
	x.String(e.GB)
	return x.buf
}

func (e DhGenOk) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDhGenOk)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash1)
	return x.buf
}

func (e DhGenRetry) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDhGenRetry)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash2)
	return x.buf
}

func (e DhGenFail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDhGenFail)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.Bytes(e.NewNonceHash3)
	return x.buf
}

func (e BindAuthKeyInner) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBindAuthKeyInner)
	x.Long(e.Nonce)
	x.Long(e.TempAuthKeyID)
	x.Long(e.PermAuthKeyID)
	x.Long(e.TempSessionID)
	x.Int(e.ExpiresAt)
	return x.buf
}

func (e RpcError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRpcError)
	x.Int(e.ErrorCode)
	x.String(e.ErrorMessage)
	return x.buf
}

func (e RpcAnswerUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRpcAnswerUnknown)
	return x.buf
}

func (e RpcAnswerDroppedRunning) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRpcAnswerDroppedRunning)
	return x.buf
}

func (e RpcAnswerDropped) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRpcAnswerDropped)
	x.Long(e.MsgID)
	x.Int(e.SeqNo)
	x.Int(e.Bytes)
	return x.buf
}

func (e FutureSalt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFutureSalt)
	x.Int(e.ValidSince)
	x.Int(e.ValidUntil)
	x.Long(e.Salt)
	return x.buf
}

func (e FutureSalts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFutureSalts)
	x.Long(e.ReqMsgID)
	x.Int(e.Now)
	x.Vector(e.Salts)
	return x.buf
}

func (e Pong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPong)
	x.Long(e.MsgID)
	x.Long(e.PingID)
	return x.buf
}

func (e DestroySessionOk) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroySessionOk)
	x.Long(e.SessionID)
	return x.buf
}

func (e DestroySessionNone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroySessionNone)
	x.Long(e.SessionID)
	return x.buf
}

func (e NewSessionCreated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcNewSessionCreated)
	x.Long(e.FirstMsgID)
	x.Long(e.UniqueID)
	x.Long(e.ServerSalt)
	return x.buf
}

func (e MsgsAck) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgsAck)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e BadMsgNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBadMsgNotification)
	x.Long(e.BadMsgID)
	x.Int(e.BadMsgSeqno)
	x.Int(e.ErrorCode)
	return x.buf
}

func (e BadServerSalt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBadServerSalt)
	x.Long(e.BadMsgID)
	x.Int(e.BadMsgSeqno)
	x.Int(e.ErrorCode)
	x.Long(e.NewServerSalt)
	return x.buf
}

func (e MsgResendReq) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgResendReq)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e MsgsStateReq) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgsStateReq)
	x.VectorLong(e.MsgIds)
	return x.buf
}

func (e MsgsStateInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgsStateInfo)
	x.Long(e.ReqMsgID)
	x.String(e.Info)
	return x.buf
}

func (e MsgsAllInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgsAllInfo)
	x.VectorLong(e.MsgIds)
	x.String(e.Info)
	return x.buf
}

func (e MsgDetailedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgDetailedInfo)
	x.Long(e.MsgID)
	x.Long(e.AnswerMsgID)
	x.Int(e.Bytes)
	x.Int(e.Status)
	return x.buf
}

func (e MsgNewDetailedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMsgNewDetailedInfo)
	x.Long(e.AnswerMsgID)
	x.Int(e.Bytes)
	x.Int(e.Status)
	return x.buf
}

func (e DestroyAuthKeyOk) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroyAuthKeyOk)
	return x.buf
}

func (e DestroyAuthKeyNone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroyAuthKeyNone)
	return x.buf
}

func (e DestroyAuthKeyFail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroyAuthKeyFail)
	return x.buf
}

func (e ReqPqMulti) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReqPqMulti)
	x.Bytes(e.Nonce)
	return x.buf
}

func (e ReqDHParams) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReqDHParams)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.P)
	x.String(e.Q)
	x.Long(e.PublicKeyFingerprint)
	x.String(e.EncryptedData)
	return x.buf
}

func (e SetClientDHParams) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSetClientDHParams)
	x.Bytes(e.Nonce)
	x.Bytes(e.ServerNonce)
	x.String(e.EncryptedData)
	return x.buf
}

func (e RpcDropAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRpcDropAnswer)
	x.Long(e.ReqMsgID)
	return x.buf
}

func (e GetFutureSalts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcGetFutureSalts)
	x.Int(e.Num)
	return x.buf
}

func (e Ping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPing)
	x.Long(e.PingID)
	return x.buf
}

func (e PingDelayDisconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPingDelayDisconnect)
	x.Long(e.PingID)
	x.Int(e.DisconnectDelay)
	return x.buf
}

func (e DestroySession) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroySession)
	x.Long(e.SessionID)
	return x.buf
}

func (e HttpWait) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHttpWait)
	x.Int(e.MaxDelay)
	x.Int(e.WaitAfter)
	x.Int(e.MaxWait)
	return x.buf
}

func (e DestroyAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDestroyAuthKey)
	return x.buf
}

func (e True) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTrue)
	return x.buf
}

func (e BoolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBoolFalse)
	return x.buf
}

func (e BoolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBoolTrue)
	return x.buf
}

func (e Error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcError)
	x.Int(e.Code)
	x.String(e.Text)
	return x.buf
}

func (e IpPort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcIpPort)
	x.Int(e.Ipv4)
	x.Int(e.Port)
	return x.buf
}

func (e InputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerEmpty)
	return x.buf
}

func (e InputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerSelf)
	return x.buf
}

func (e InputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e InputPeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerUser)
	x.Int(e.UserID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputPeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerChannel)
	x.Int(e.ChannelID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputPeerUserFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerUserFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.UserID)
	return x.buf
}

func (e InputPeerChannelFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerChannelFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ChannelID)
	return x.buf
}

func (e InputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputUserEmpty)
	return x.buf
}

func (e InputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputUserSelf)
	return x.buf
}

func (e InputUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputUser)
	x.Int(e.UserID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputUserFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputUserFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.UserID)
	return x.buf
}

func (e InputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPhoneContact)
	x.Long(e.ClientID)
	x.String(e.Phone)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e InputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputFile)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Name)
	x.String(e.Md5Checksum)
	return x.buf
}

func (e InputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputFileBig)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Name)
	return x.buf
}

func (e InputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaEmpty)
	return x.buf
}

func (e InputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaUploadedPhoto)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	if e.Flags&1 != 0 {
		x.Vector(e.Stickers)
	}
	if e.Flags&2 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e InputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaPhoto)
	x.Int(e.Flags)
	x.Bytes(e.ID.encode())
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e InputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaGeoPoint)
	x.Bytes(e.GeoPoint.encode())
	return x.buf
}

func (e InputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	return x.buf
}

func (e InputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaUploadedDocument)
	x.Int(e.Flags)
	//flag e.NosoundVideo
	//flag e.ForceFile
	x.Bytes(e.File.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	if e.Flags&1 != 0 {
		x.Vector(e.Stickers)
	}
	if e.Flags&2 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e InputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaDocument)
	x.Int(e.Flags)
	x.Bytes(e.ID.encode())
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e InputMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaVenue)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	return x.buf
}

func (e InputMediaPhotoExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaPhotoExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e InputMediaDocumentExternal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaDocumentExternal)
	x.Int(e.Flags)
	x.String(e.Url)
	if e.Flags&1 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e InputMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaGame)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e InputMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaInvoice)
	x.Int(e.Flags)
	x.String(e.Title)
	x.String(e.Description)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	x.Bytes(e.Invoice.encode())
	x.StringBytes(e.Payload)
	x.String(e.Provider)
	x.Bytes(e.ProviderData.encode())
	x.String(e.StartParam)
	return x.buf
}

func (e InputMediaGeoLive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaGeoLive)
	x.Int(e.Flags)
	//flag e.Stopped
	x.Bytes(e.GeoPoint.encode())
	if e.Flags&2 != 0 {
		x.Int(e.Period)
	}
	return x.buf
}

func (e InputMediaPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaPoll)
	x.Int(e.Flags)
	x.Bytes(e.Poll.encode())
	if e.Flags&1 != 0 {
		x.Vector(e.CorrectAnswers)
	}
	if e.Flags&2 != 0 {
		x.String(e.Solution)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.SolutionEntities)
	}
	return x.buf
}

func (e InputMediaDice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMediaDice)
	x.String(e.Emoticon)
	return x.buf
}

func (e InputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputChatPhotoEmpty)
	return x.buf
}

func (e InputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputChatUploadedPhoto)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.File.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Video.encode())
	}
	if e.Flags&4 != 0 {
		x.Double(e.VideoStartTs)
	}
	return x.buf
}

func (e InputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputChatPhoto)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e InputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputGeoPointEmpty)
	return x.buf
}

func (e InputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputGeoPoint)
	x.Double(e.Lat)
	x.Double(e.Long)
	return x.buf
}

func (e InputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPhotoEmpty)
	return x.buf
}

func (e InputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPhoto)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	return x.buf
}

func (e InputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputFileLocation)
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	x.Long(e.Secret)
	x.StringBytes(e.FileReference)
	return x.buf
}

func (e InputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputEncryptedFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputDocumentFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.String(e.ThumbSize)
	return x.buf
}

func (e InputSecureFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputSecureFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputTakeoutFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputTakeoutFileLocation)
	return x.buf
}

func (e InputPhotoFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPhotoFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.String(e.ThumbSize)
	return x.buf
}

func (e InputPhotoLegacyFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPhotoLegacyFileLocation)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	x.Long(e.Secret)
	return x.buf
}

func (e InputPeerPhotoFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerPhotoFileLocation)
	x.Int(e.Flags)
	//flag e.Big
	x.Bytes(e.Peer.encode())
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	return x.buf
}

func (e InputStickerSetThumb) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetThumb)
	x.Bytes(e.Stickerset.encode())
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	return x.buf
}

func (e PeerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerUser)
	x.Int(e.UserID)
	return x.buf
}

func (e PeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e PeerChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerChannel)
	x.Int(e.ChannelID)
	return x.buf
}

func (e StorageFileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileUnknown)
	return x.buf
}

func (e StorageFilePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFilePartial)
	return x.buf
}

func (e StorageFileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileJpeg)
	return x.buf
}

func (e StorageFileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileGif)
	return x.buf
}

func (e StorageFilePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFilePng)
	return x.buf
}

func (e StorageFilePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFilePdf)
	return x.buf
}

func (e StorageFileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileMp3)
	return x.buf
}

func (e StorageFileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileMov)
	return x.buf
}

func (e StorageFileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileMp4)
	return x.buf
}

func (e StorageFileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStorageFileWebp)
	return x.buf
}

func (e UserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e User) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUser)
	x.Int(e.Flags)
	//flag e.Self
	//flag e.Contact
	//flag e.MutualContact
	//flag e.Deleted
	//flag e.Bot
	//flag e.BotChatHistory
	//flag e.BotNochats
	//flag e.Verified
	//flag e.Restricted
	//flag e.Min
	//flag e.BotInlineGeo
	//flag e.Support
	//flag e.Scam
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.Long(e.AccessHash)
	}
	if e.Flags&2 != 0 {
		x.String(e.FirstName)
	}
	if e.Flags&4 != 0 {
		x.String(e.LastName)
	}
	if e.Flags&8 != 0 {
		x.String(e.Username)
	}
	if e.Flags&16 != 0 {
		x.String(e.Phone)
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&64 != 0 {
		x.Bytes(e.Status.encode())
	}
	if e.Flags&16384 != 0 {
		x.Int(e.BotInfoVersion)
	}
	if e.Flags&262144 != 0 {
		x.Vector(e.RestrictionReason)
	}
	if e.Flags&524288 != 0 {
		x.String(e.BotInlinePlaceholder)
	}
	if e.Flags&4194304 != 0 {
		x.String(e.LangCode)
	}
	return x.buf
}

func (e UserProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserProfilePhotoEmpty)
	return x.buf
}

func (e UserProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserProfilePhoto)
	x.Int(e.Flags)
	//flag e.HasVideo
	x.Long(e.PhotoID)
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	x.Int(e.DcID)
	return x.buf
}

func (e UserStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserStatusEmpty)
	return x.buf
}

func (e UserStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserStatusOnline)
	x.Int(e.Expires)
	return x.buf
}

func (e UserStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserStatusOffline)
	x.Int(e.WasOnline)
	return x.buf
}

func (e UserStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserStatusRecently)
	return x.buf
}

func (e UserStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserStatusLastWeek)
	return x.buf
}

func (e UserStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserStatusLastMonth)
	return x.buf
}

func (e ChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e Chat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChat)
	x.Int(e.Flags)
	//flag e.Creator
	//flag e.Kicked
	//flag e.Left
	//flag e.Deactivated
	x.Int(e.ID)
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	x.Int(e.Date)
	x.Int(e.Version)
	if e.Flags&64 != 0 {
		x.Bytes(e.MigratedTo.encode())
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.AdminRights.encode())
	}
	if e.Flags&262144 != 0 {
		x.Bytes(e.DefaultBannedRights.encode())
	}
	return x.buf
}

func (e ChatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatForbidden)
	x.Int(e.ID)
	x.String(e.Title)
	return x.buf
}

func (e Channel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannel)
	x.Int(e.Flags)
	//flag e.Creator
	//flag e.Left
	//flag e.Broadcast
	//flag e.Verified
	//flag e.Megagroup
	//flag e.Restricted
	//flag e.Signatures
	//flag e.Min
	//flag e.Scam
	//flag e.HasLink
	//flag e.HasGeo
	//flag e.SlowmodeEnabled
	x.Int(e.ID)
	if e.Flags&8192 != 0 {
		x.Long(e.AccessHash)
	}
	x.String(e.Title)
	if e.Flags&64 != 0 {
		x.String(e.Username)
	}
	x.Bytes(e.Photo.encode())
	x.Int(e.Date)
	x.Int(e.Version)
	if e.Flags&512 != 0 {
		x.Vector(e.RestrictionReason)
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.AdminRights.encode())
	}
	if e.Flags&32768 != 0 {
		x.Bytes(e.BannedRights.encode())
	}
	if e.Flags&262144 != 0 {
		x.Bytes(e.DefaultBannedRights.encode())
	}
	if e.Flags&131072 != 0 {
		x.Int(e.ParticipantsCount)
	}
	return x.buf
}

func (e ChannelForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelForbidden)
	x.Int(e.Flags)
	//flag e.Broadcast
	//flag e.Megagroup
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.String(e.Title)
	if e.Flags&65536 != 0 {
		x.Int(e.UntilDate)
	}
	return x.buf
}

func (e ChatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatFull)
	x.Int(e.Flags)
	//flag e.CanSetUsername
	//flag e.HasScheduled
	x.Int(e.ID)
	x.String(e.About)
	x.Bytes(e.Participants.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.ChatPhoto.encode())
	}
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.ExportedInvite.encode())
	if e.Flags&8 != 0 {
		x.Vector(e.BotInfo)
	}
	if e.Flags&64 != 0 {
		x.Int(e.PinnedMsgID)
	}
	if e.Flags&2048 != 0 {
		x.Int(e.FolderID)
	}
	return x.buf
}

func (e ChannelFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelFull)
	x.Int(e.Flags)
	//flag e.CanViewParticipants
	//flag e.CanSetUsername
	//flag e.CanSetStickers
	//flag e.HiddenPrehistory
	//flag e.CanViewStats
	//flag e.CanSetLocation
	//flag e.HasScheduled
	x.Int(e.ID)
	x.String(e.About)
	if e.Flags&1 != 0 {
		x.Int(e.ParticipantsCount)
	}
	if e.Flags&2 != 0 {
		x.Int(e.AdminsCount)
	}
	if e.Flags&4 != 0 {
		x.Int(e.KickedCount)
	}
	if e.Flags&4 != 0 {
		x.Int(e.BannedCount)
	}
	if e.Flags&8192 != 0 {
		x.Int(e.OnlineCount)
	}
	x.Int(e.ReadInboxMaxID)
	x.Int(e.ReadOutboxMaxID)
	x.Int(e.UnreadCount)
	x.Bytes(e.ChatPhoto.encode())
	x.Bytes(e.NotifySettings.encode())
	x.Bytes(e.ExportedInvite.encode())
	x.Vector(e.BotInfo)
	if e.Flags&16 != 0 {
		x.Int(e.MigratedFromChatID)
	}
	if e.Flags&16 != 0 {
		x.Int(e.MigratedFromMaxID)
	}
	if e.Flags&32 != 0 {
		x.Int(e.PinnedMsgID)
	}
	if e.Flags&256 != 0 {
		x.Bytes(e.Stickerset.encode())
	}
	if e.Flags&512 != 0 {
		x.Int(e.AvailableMinID)
	}
	if e.Flags&2048 != 0 {
		x.Int(e.FolderID)
	}
	if e.Flags&16384 != 0 {
		x.Int(e.LinkedChatID)
	}
	if e.Flags&32768 != 0 {
		x.Bytes(e.Location.encode())
	}
	if e.Flags&131072 != 0 {
		x.Int(e.SlowmodeSeconds)
	}
	if e.Flags&262144 != 0 {
		x.Int(e.SlowmodeNextSendDate)
	}
	if e.Flags&4096 != 0 {
		x.Int(e.StatsDc)
	}
	x.Int(e.Pts)
	return x.buf
}

func (e ChatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatParticipant)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	return x.buf
}

func (e ChatParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatParticipantCreator)
	x.Int(e.UserID)
	return x.buf
}

func (e ChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatParticipantAdmin)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	return x.buf
}

func (e ChatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatParticipantsForbidden)
	x.Int(e.Flags)
	x.Int(e.ChatID)
	if e.Flags&1 != 0 {
		x.Bytes(e.SelfParticipant.encode())
	}
	return x.buf
}

func (e ChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatParticipants)
	x.Int(e.ChatID)
	x.Vector(e.Participants)
	x.Int(e.Version)
	return x.buf
}

func (e ChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatPhotoEmpty)
	return x.buf
}

func (e ChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatPhoto)
	x.Int(e.Flags)
	//flag e.HasVideo
	x.Bytes(e.PhotoSmall.encode())
	x.Bytes(e.PhotoBig.encode())
	x.Int(e.DcID)
	return x.buf
}

func (e MessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e Message) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessage)
	x.Int(e.Flags)
	//flag e.Out
	//flag e.Mentioned
	//flag e.MediaUnread
	//flag e.Silent
	//flag e.Post
	//flag e.FromScheduled
	//flag e.Legacy
	//flag e.EditHide
	x.Int(e.ID)
	if e.Flags&256 != 0 {
		x.Int(e.FromID)
	}
	x.Bytes(e.ToID.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.FwdFrom.encode())
	}
	if e.Flags&2048 != 0 {
		x.Int(e.ViaBotID)
	}
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Int(e.Date)
	x.String(e.Message)
	if e.Flags&512 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&64 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&1024 != 0 {
		x.Int(e.Views)
	}
	if e.Flags&32768 != 0 {
		x.Int(e.EditDate)
	}
	if e.Flags&65536 != 0 {
		x.String(e.PostAuthor)
	}
	if e.Flags&131072 != 0 {
		x.Long(e.GroupedID)
	}
	if e.Flags&4194304 != 0 {
		x.Vector(e.RestrictionReason)
	}
	return x.buf
}

func (e MessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageService)
	x.Int(e.Flags)
	//flag e.Out
	//flag e.Mentioned
	//flag e.MediaUnread
	//flag e.Silent
	//flag e.Post
	//flag e.Legacy
	x.Int(e.ID)
	if e.Flags&256 != 0 {
		x.Int(e.FromID)
	}
	x.Bytes(e.ToID.encode())
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Int(e.Date)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e MessageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaEmpty)
	return x.buf
}

func (e MessageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaPhoto)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e MessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaGeo)
	x.Bytes(e.Geo.encode())
	return x.buf
}

func (e MessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaContact)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	x.Int(e.UserID)
	return x.buf
}

func (e MessageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaUnsupported)
	return x.buf
}

func (e MessageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaDocument)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.TtlSeconds)
	}
	return x.buf
}

func (e MessageMediaWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaWebPage)
	x.Bytes(e.Webpage.encode())
	return x.buf
}

func (e MessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaVenue)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	return x.buf
}

func (e MessageMediaGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaGame)
	x.Bytes(e.Game.encode())
	return x.buf
}

func (e MessageMediaInvoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaInvoice)
	x.Int(e.Flags)
	//flag e.ShippingAddressRequested
	//flag e.Test
	x.String(e.Title)
	x.String(e.Description)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.ReceiptMsgID)
	}
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.String(e.StartParam)
	return x.buf
}

func (e MessageMediaGeoLive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaGeoLive)
	x.Bytes(e.Geo.encode())
	x.Int(e.Period)
	return x.buf
}

func (e MessageMediaPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaPoll)
	x.Bytes(e.Poll.encode())
	x.Bytes(e.Results.encode())
	return x.buf
}

func (e MessageMediaDice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageMediaDice)
	x.Int(e.Value)
	x.String(e.Emoticon)
	return x.buf
}

func (e MessageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionEmpty)
	return x.buf
}

func (e MessageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatCreate)
	x.String(e.Title)
	x.VectorInt(e.Users)
	return x.buf
}

func (e MessageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatEditTitle)
	x.String(e.Title)
	return x.buf
}

func (e MessageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatEditPhoto)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e MessageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatDeletePhoto)
	return x.buf
}

func (e MessageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatAddUser)
	x.VectorInt(e.Users)
	return x.buf
}

func (e MessageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatDeleteUser)
	x.Int(e.UserID)
	return x.buf
}

func (e MessageActionChatJoinedByLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatJoinedByLink)
	x.Int(e.InviterID)
	return x.buf
}

func (e MessageActionChannelCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChannelCreate)
	x.String(e.Title)
	return x.buf
}

func (e MessageActionChatMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChatMigrateTo)
	x.Int(e.ChannelID)
	return x.buf
}

func (e MessageActionChannelMigrateFrom) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionChannelMigrateFrom)
	x.String(e.Title)
	x.Int(e.ChatID)
	return x.buf
}

func (e MessageActionPinMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionPinMessage)
	return x.buf
}

func (e MessageActionHistoryClear) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionHistoryClear)
	return x.buf
}

func (e MessageActionGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionGameScore)
	x.Long(e.GameID)
	x.Int(e.Score)
	return x.buf
}

func (e MessageActionPaymentSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionPaymentSentMe)
	x.Int(e.Flags)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.StringBytes(e.Payload)
	if e.Flags&1 != 0 {
		x.Bytes(e.Info.encode())
	}
	if e.Flags&2 != 0 {
		x.String(e.ShippingOptionID)
	}
	x.Bytes(e.Charge.encode())
	return x.buf
}

func (e MessageActionPaymentSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionPaymentSent)
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	return x.buf
}

func (e MessageActionPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionPhoneCall)
	x.Int(e.Flags)
	//flag e.Video
	x.Long(e.CallID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Reason.encode())
	}
	if e.Flags&2 != 0 {
		x.Int(e.Duration)
	}
	return x.buf
}

func (e MessageActionScreenshotTaken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionScreenshotTaken)
	return x.buf
}

func (e MessageActionCustomAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionCustomAction)
	x.String(e.Message)
	return x.buf
}

func (e MessageActionBotAllowed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionBotAllowed)
	x.String(e.Domain)
	return x.buf
}

func (e MessageActionSecureValuesSentMe) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionSecureValuesSentMe)
	x.Vector(e.Values)
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e MessageActionSecureValuesSent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionSecureValuesSent)
	x.Vector(e.Types)
	return x.buf
}

func (e MessageActionContactSignUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageActionContactSignUp)
	return x.buf
}

func (e Dialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDialog)
	x.Int(e.Flags)
	//flag e.Pinned
	//flag e.UnreadMark
	x.Bytes(e.Peer.encode())
	x.Int(e.TopMessage)
	x.Int(e.ReadInboxMaxID)
	x.Int(e.ReadOutboxMaxID)
	x.Int(e.UnreadCount)
	x.Int(e.UnreadMentionsCount)
	x.Bytes(e.NotifySettings.encode())
	if e.Flags&1 != 0 {
		x.Int(e.Pts)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Draft.encode())
	}
	if e.Flags&16 != 0 {
		x.Int(e.FolderID)
	}
	return x.buf
}

func (e DialogFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDialogFolder)
	x.Int(e.Flags)
	//flag e.Pinned
	x.Bytes(e.Folder.encode())
	x.Bytes(e.Peer.encode())
	x.Int(e.TopMessage)
	x.Int(e.UnreadMutedPeersCount)
	x.Int(e.UnreadUnmutedPeersCount)
	x.Int(e.UnreadMutedMessagesCount)
	x.Int(e.UnreadUnmutedMessagesCount)
	return x.buf
}

func (e PhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotoEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e Photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoto)
	x.Int(e.Flags)
	//flag e.HasStickers
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.Int(e.Date)
	x.Vector(e.Sizes)
	if e.Flags&2 != 0 {
		x.Vector(e.VideoSizes)
	}
	x.Int(e.DcID)
	return x.buf
}

func (e PhotoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotoSizeEmpty)
	x.String(e.Type)
	return x.buf
}

func (e PhotoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotoSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	return x.buf
}

func (e PhotoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotoCachedSize)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e PhotoStrippedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotoStrippedSize)
	x.String(e.Type)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e GeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcGeoPointEmpty)
	return x.buf
}

func (e GeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcGeoPoint)
	x.Double(e.Long)
	x.Double(e.Lat)
	x.Long(e.AccessHash)
	return x.buf
}

func (e AuthSentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSentCode)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	x.String(e.PhoneCodeHash)
	if e.Flags&2 != 0 {
		x.Bytes(e.NextType.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.Timeout)
	}
	return x.buf
}

func (e AuthAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthAuthorization)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.TmpSessions)
	}
	x.Bytes(e.User.encode())
	return x.buf
}

func (e AuthAuthorizationSignUpRequired) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthAuthorizationSignUpRequired)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.TermsOfService.encode())
	}
	return x.buf
}

func (e AuthExportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthExportedAuthorization)
	x.Int(e.ID)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e InputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e InputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputNotifyUsers)
	return x.buf
}

func (e InputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputNotifyChats)
	return x.buf
}

func (e InputNotifyBroadcasts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputNotifyBroadcasts)
	return x.buf
}

func (e InputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPeerNotifySettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.ShowPreviews.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Silent.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.MuteUntil)
	}
	if e.Flags&8 != 0 {
		x.String(e.Sound)
	}
	return x.buf
}

func (e PeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerNotifySettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.ShowPreviews.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Silent.encode())
	}
	if e.Flags&4 != 0 {
		x.Int(e.MuteUntil)
	}
	if e.Flags&8 != 0 {
		x.String(e.Sound)
	}
	return x.buf
}

func (e PeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerSettings)
	x.Int(e.Flags)
	//flag e.ReportSpam
	//flag e.AddContact
	//flag e.BlockContact
	//flag e.ShareContact
	//flag e.NeedContactsException
	//flag e.ReportGeo
	//flag e.Autoarchived
	if e.Flags&64 != 0 {
		x.Int(e.GeoDistance)
	}
	return x.buf
}

func (e WallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWallPaper)
	x.Long(e.ID)
	x.Int(e.Flags)
	//flag e.Creator
	//flag e.Default
	//flag e.Pattern
	//flag e.Dark
	x.Long(e.AccessHash)
	x.String(e.Slug)
	x.Bytes(e.Document.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e WallPaperNoFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWallPaperNoFile)
	x.Int(e.Flags)
	//flag e.Default
	//flag e.Dark
	if e.Flags&4 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e InputReportReasonSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonSpam)
	return x.buf
}

func (e InputReportReasonViolence) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonViolence)
	return x.buf
}

func (e InputReportReasonPornography) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonPornography)
	return x.buf
}

func (e InputReportReasonChildAbuse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonChildAbuse)
	return x.buf
}

func (e InputReportReasonOther) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonOther)
	x.String(e.Text)
	return x.buf
}

func (e InputReportReasonCopyright) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonCopyright)
	return x.buf
}

func (e InputReportReasonGeoIrrelevant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputReportReasonGeoIrrelevant)
	return x.buf
}

func (e UserFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUserFull)
	x.Int(e.Flags)
	//flag e.Blocked
	//flag e.PhoneCallsAvailable
	//flag e.PhoneCallsPrivate
	//flag e.CanPinMessage
	//flag e.HasScheduled
	x.Bytes(e.User.encode())
	if e.Flags&2 != 0 {
		x.String(e.About)
	}
	x.Bytes(e.Settings.encode())
	if e.Flags&4 != 0 {
		x.Bytes(e.ProfilePhoto.encode())
	}
	x.Bytes(e.NotifySettings.encode())
	if e.Flags&8 != 0 {
		x.Bytes(e.BotInfo.encode())
	}
	if e.Flags&64 != 0 {
		x.Int(e.PinnedMsgID)
	}
	x.Int(e.CommonChatsCount)
	if e.Flags&2048 != 0 {
		x.Int(e.FolderID)
	}
	return x.buf
}

func (e Contact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContact)
	x.Int(e.UserID)
	x.Bytes(e.Mutual.encode())
	return x.buf
}

func (e ImportedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcImportedContact)
	x.Int(e.UserID)
	x.Long(e.ClientID)
	return x.buf
}

func (e ContactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactBlocked)
	x.Int(e.UserID)
	x.Int(e.Date)
	return x.buf
}

func (e ContactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactStatus)
	x.Int(e.UserID)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e ContactsContactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsContactsNotModified)
	return x.buf
}

func (e ContactsContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsContacts)
	x.Vector(e.Contacts)
	x.Int(e.SavedCount)
	x.Vector(e.Users)
	return x.buf
}

func (e ContactsImportedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsImportedContacts)
	x.Vector(e.Imported)
	x.Vector(e.PopularInvites)
	x.VectorLong(e.RetryContacts)
	x.Vector(e.Users)
	return x.buf
}

func (e ContactsBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsBlocked)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e ContactsBlockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsBlockedSlice)
	x.Int(e.Count)
	x.Vector(e.Blocked)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesDialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDialogsSlice)
	x.Int(e.Count)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesDialogsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDialogsNotModified)
	x.Int(e.Count)
	return x.buf
}

func (e MessagesMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesMessages)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesMessagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesMessagesSlice)
	x.Int(e.Flags)
	//flag e.Inexact
	x.Int(e.Count)
	if e.Flags&1 != 0 {
		x.Int(e.NextRate)
	}
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesChannelMessages)
	x.Int(e.Flags)
	//flag e.Inexact
	x.Int(e.Pts)
	x.Int(e.Count)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesMessagesNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesMessagesNotModified)
	x.Int(e.Count)
	return x.buf
}

func (e MessagesChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesChats)
	x.Vector(e.Chats)
	return x.buf
}

func (e MessagesChatsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesChatsSlice)
	x.Int(e.Count)
	x.Vector(e.Chats)
	return x.buf
}

func (e MessagesChatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesChatFull)
	x.Bytes(e.FullChat.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessagesAffectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAffectedHistory)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Offset)
	return x.buf
}

func (e InputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterEmpty)
	return x.buf
}

func (e InputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterPhotos)
	return x.buf
}

func (e InputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterVideo)
	return x.buf
}

func (e InputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterPhotoVideo)
	return x.buf
}

func (e InputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterDocument)
	return x.buf
}

func (e InputMessagesFilterUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterUrl)
	return x.buf
}

func (e InputMessagesFilterGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterGif)
	return x.buf
}

func (e InputMessagesFilterVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterVoice)
	return x.buf
}

func (e InputMessagesFilterMusic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterMusic)
	return x.buf
}

func (e InputMessagesFilterChatPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterChatPhotos)
	return x.buf
}

func (e InputMessagesFilterPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterPhoneCalls)
	x.Int(e.Flags)
	//flag e.Missed
	return x.buf
}

func (e InputMessagesFilterRoundVoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterRoundVoice)
	return x.buf
}

func (e InputMessagesFilterRoundVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterRoundVideo)
	return x.buf
}

func (e InputMessagesFilterMyMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterMyMentions)
	return x.buf
}

func (e InputMessagesFilterGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterGeo)
	return x.buf
}

func (e InputMessagesFilterContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagesFilterContacts)
	return x.buf
}

func (e UpdateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateNewMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateMessageID)
	x.Int(e.ID)
	x.Long(e.RandomID)
	return x.buf
}

func (e UpdateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDeleteMessages)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserTyping)
	x.Int(e.UserID)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e UpdateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatUserTyping)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e UpdateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatParticipants)
	x.Bytes(e.Participants.encode())
	return x.buf
}

func (e UpdateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserStatus)
	x.Int(e.UserID)
	x.Bytes(e.Status.encode())
	return x.buf
}

func (e UpdateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserName)
	x.Int(e.UserID)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Username)
	return x.buf
}

func (e UpdateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserPhoto)
	x.Int(e.UserID)
	x.Int(e.Date)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.Previous.encode())
	return x.buf
}

func (e UpdateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateNewEncryptedMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Qts)
	return x.buf
}

func (e UpdateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateEncryptedChatTyping)
	x.Int(e.ChatID)
	return x.buf
}

func (e UpdateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateEncryption)
	x.Bytes(e.Chat.encode())
	x.Int(e.Date)
	return x.buf
}

func (e UpdateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateEncryptedMessagesRead)
	x.Int(e.ChatID)
	x.Int(e.MaxDate)
	x.Int(e.Date)
	return x.buf
}

func (e UpdateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatParticipantAdd)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	x.Int(e.Version)
	return x.buf
}

func (e UpdateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatParticipantDelete)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Int(e.Version)
	return x.buf
}

func (e UpdateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDcOptions)
	x.Vector(e.DcOptions)
	return x.buf
}

func (e UpdateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserBlocked)
	x.Int(e.UserID)
	x.Bytes(e.Blocked.encode())
	return x.buf
}

func (e UpdateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.NotifySettings.encode())
	return x.buf
}

func (e UpdateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateServiceNotification)
	x.Int(e.Flags)
	//flag e.Popup
	if e.Flags&2 != 0 {
		x.Int(e.InboxDate)
	}
	x.String(e.Type)
	x.String(e.Message)
	x.Bytes(e.Media.encode())
	x.Vector(e.Entities)
	return x.buf
}

func (e UpdatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e UpdateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserPhone)
	x.Int(e.UserID)
	x.String(e.Phone)
	return x.buf
}

func (e UpdateReadHistoryInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateReadHistoryInbox)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	x.Int(e.StillUnreadCount)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateReadHistoryOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateReadHistoryOutbox)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateWebPage)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateReadMessagesContents)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateChannelTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelTooLong)
	x.Int(e.Flags)
	x.Int(e.ChannelID)
	if e.Flags&1 != 0 {
		x.Int(e.Pts)
	}
	return x.buf
}

func (e UpdateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannel)
	x.Int(e.ChannelID)
	return x.buf
}

func (e UpdateNewChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateNewChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateReadChannelInbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateReadChannelInbox)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.Int(e.ChannelID)
	x.Int(e.MaxID)
	x.Int(e.StillUnreadCount)
	x.Int(e.Pts)
	return x.buf
}

func (e UpdateDeleteChannelMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDeleteChannelMessages)
	x.Int(e.ChannelID)
	x.VectorInt(e.Messages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateChannelMessageViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelMessageViews)
	x.Int(e.ChannelID)
	x.Int(e.ID)
	x.Int(e.Views)
	return x.buf
}

func (e UpdateChatParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatParticipantAdmin)
	x.Int(e.ChatID)
	x.Int(e.UserID)
	x.Bytes(e.IsAdmin.encode())
	x.Int(e.Version)
	return x.buf
}

func (e UpdateNewStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateNewStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e UpdateStickerSetsOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateStickerSetsOrder)
	x.Int(e.Flags)
	//flag e.Masks
	x.VectorLong(e.Order)
	return x.buf
}

func (e UpdateStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateStickerSets)
	return x.buf
}

func (e UpdateSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateSavedGifs)
	return x.buf
}

func (e UpdateBotInlineQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotInlineQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.String(e.Query)
	if e.Flags&1 != 0 {
		x.Bytes(e.Geo.encode())
	}
	x.String(e.Offset)
	return x.buf
}

func (e UpdateBotInlineSend) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotInlineSend)
	x.Int(e.Flags)
	x.Int(e.UserID)
	x.String(e.Query)
	if e.Flags&1 != 0 {
		x.Bytes(e.Geo.encode())
	}
	x.String(e.ID)
	if e.Flags&2 != 0 {
		x.Bytes(e.MsgID.encode())
	}
	return x.buf
}

func (e UpdateEditChannelMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateEditChannelMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateChannelPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelPinnedMessage)
	x.Int(e.ChannelID)
	x.Int(e.ID)
	return x.buf
}

func (e UpdateBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Long(e.ChatInstance)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Data)
	}
	if e.Flags&2 != 0 {
		x.String(e.GameShortName)
	}
	return x.buf
}

func (e UpdateEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateEditMessage)
	x.Bytes(e.Message.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateInlineBotCallbackQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateInlineBotCallbackQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.Bytes(e.MsgID.encode())
	x.Long(e.ChatInstance)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Data)
	}
	if e.Flags&2 != 0 {
		x.String(e.GameShortName)
	}
	return x.buf
}

func (e UpdateReadChannelOutbox) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateReadChannelOutbox)
	x.Int(e.ChannelID)
	x.Int(e.MaxID)
	return x.buf
}

func (e UpdateDraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDraftMessage)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Draft.encode())
	return x.buf
}

func (e UpdateReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateReadFeaturedStickers)
	return x.buf
}

func (e UpdateRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateRecentStickers)
	return x.buf
}

func (e UpdateConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateConfig)
	return x.buf
}

func (e UpdatePtsChanged) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePtsChanged)
	return x.buf
}

func (e UpdateChannelWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelWebPage)
	x.Int(e.ChannelID)
	x.Bytes(e.Webpage.encode())
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdateDialogPinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDialogPinned)
	x.Int(e.Flags)
	//flag e.Pinned
	if e.Flags&2 != 0 {
		x.Int(e.FolderID)
	}
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e UpdatePinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePinnedDialogs)
	x.Int(e.Flags)
	if e.Flags&2 != 0 {
		x.Int(e.FolderID)
	}
	if e.Flags&1 != 0 {
		x.Vector(e.Order)
	}
	return x.buf
}

func (e UpdateBotWebhookJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotWebhookJSON)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e UpdateBotWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotWebhookJSONQuery)
	x.Long(e.QueryID)
	x.Bytes(e.Data.encode())
	x.Int(e.Timeout)
	return x.buf
}

func (e UpdateBotShippingQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotShippingQuery)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.StringBytes(e.Payload)
	x.Bytes(e.ShippingAddress.encode())
	return x.buf
}

func (e UpdateBotPrecheckoutQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateBotPrecheckoutQuery)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	x.Int(e.UserID)
	x.StringBytes(e.Payload)
	if e.Flags&1 != 0 {
		x.Bytes(e.Info.encode())
	}
	if e.Flags&2 != 0 {
		x.String(e.ShippingOptionID)
	}
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	return x.buf
}

func (e UpdatePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePhoneCall)
	x.Bytes(e.PhoneCall.encode())
	return x.buf
}

func (e UpdateLangPackTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateLangPackTooLong)
	x.String(e.LangCode)
	return x.buf
}

func (e UpdateLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateLangPack)
	x.Bytes(e.Difference.encode())
	return x.buf
}

func (e UpdateFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateFavedStickers)
	return x.buf
}

func (e UpdateChannelReadMessagesContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelReadMessagesContents)
	x.Int(e.ChannelID)
	x.VectorInt(e.Messages)
	return x.buf
}

func (e UpdateContactsReset) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateContactsReset)
	return x.buf
}

func (e UpdateChannelAvailableMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelAvailableMessages)
	x.Int(e.ChannelID)
	x.Int(e.AvailableMinID)
	return x.buf
}

func (e UpdateDialogUnreadMark) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDialogUnreadMark)
	x.Int(e.Flags)
	//flag e.Unread
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e UpdateUserPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateUserPinnedMessage)
	x.Int(e.UserID)
	x.Int(e.ID)
	return x.buf
}

func (e UpdateChatPinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatPinnedMessage)
	x.Int(e.ChatID)
	x.Int(e.ID)
	x.Int(e.Version)
	return x.buf
}

func (e UpdateMessagePoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateMessagePoll)
	x.Int(e.Flags)
	x.Long(e.PollID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Poll.encode())
	}
	x.Bytes(e.Results.encode())
	return x.buf
}

func (e UpdateChatDefaultBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChatDefaultBannedRights)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.DefaultBannedRights.encode())
	x.Int(e.Version)
	return x.buf
}

func (e UpdateFolderPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateFolderPeers)
	x.Vector(e.FolderPeers)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e UpdatePeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePeerSettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e UpdatePeerLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePeerLocated)
	x.Vector(e.Peers)
	return x.buf
}

func (e UpdateNewScheduledMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateNewScheduledMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e UpdateDeleteScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDeleteScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.Messages)
	return x.buf
}

func (e UpdateTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateTheme)
	x.Bytes(e.Theme.encode())
	return x.buf
}

func (e UpdateGeoLiveViewed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateGeoLiveViewed)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	return x.buf
}

func (e UpdateLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateLoginToken)
	return x.buf
}

func (e UpdateMessagePollVote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateMessagePollVote)
	x.Long(e.PollID)
	x.Int(e.UserID)
	x.Vector(e.Options)
	return x.buf
}

func (e UpdateDialogFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDialogFilter)
	x.Int(e.Flags)
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Filter.encode())
	}
	return x.buf
}

func (e UpdateDialogFilterOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDialogFilterOrder)
	x.VectorInt(e.Order)
	return x.buf
}

func (e UpdateDialogFilters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateDialogFilters)
	return x.buf
}

func (e UpdatePhoneCallSignalingData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatePhoneCallSignalingData)
	x.Long(e.PhoneCallID)
	x.StringBytes(e.Data)
	return x.buf
}

func (e UpdateChannelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateChannelParticipant)
	x.Int(e.Flags)
	x.Int(e.ChannelID)
	x.Int(e.Date)
	x.Int(e.UserID)
	if e.Flags&1 != 0 {
		x.Bytes(e.PrevParticipant.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.NewParticipant.encode())
	}
	x.Int(e.Qts)
	return x.buf
}

func (e UpdatesState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesState)
	x.Int(e.Pts)
	x.Int(e.Qts)
	x.Int(e.Date)
	x.Int(e.Seq)
	x.Int(e.UnreadCount)
	return x.buf
}

func (e UpdatesDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesDifferenceEmpty)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e UpdatesDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesDifference)
	x.Vector(e.NewMessages)
	x.Vector(e.NewEncryptedMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e UpdatesDifferenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesDifferenceSlice)
	x.Vector(e.NewMessages)
	x.Vector(e.NewEncryptedMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.IntermediateState.encode())
	return x.buf
}

func (e UpdatesDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesDifferenceTooLong)
	x.Int(e.Pts)
	return x.buf
}

func (e UpdatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesTooLong)
	return x.buf
}

func (e UpdateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateShortMessage)
	x.Int(e.Flags)
	//flag e.Out
	//flag e.Mentioned
	//flag e.MediaUnread
	//flag e.Silent
	x.Int(e.ID)
	x.Int(e.UserID)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	if e.Flags&4 != 0 {
		x.Bytes(e.FwdFrom.encode())
	}
	if e.Flags&2048 != 0 {
		x.Int(e.ViaBotID)
	}
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e UpdateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateShortChatMessage)
	x.Int(e.Flags)
	//flag e.Out
	//flag e.Mentioned
	//flag e.MediaUnread
	//flag e.Silent
	x.Int(e.ID)
	x.Int(e.FromID)
	x.Int(e.ChatID)
	x.String(e.Message)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	if e.Flags&4 != 0 {
		x.Bytes(e.FwdFrom.encode())
	}
	if e.Flags&2048 != 0 {
		x.Int(e.ViaBotID)
	}
	if e.Flags&8 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e UpdateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateShort)
	x.Bytes(e.Update.encode())
	x.Int(e.Date)
	return x.buf
}

func (e UpdatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesCombined)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.SeqStart)
	x.Int(e.Seq)
	return x.buf
}

func (e Updates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdates)
	x.Vector(e.Updates)
	x.Vector(e.Users)
	x.Vector(e.Chats)
	x.Int(e.Date)
	x.Int(e.Seq)
	return x.buf
}

func (e UpdateShortSentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdateShortSentMessage)
	x.Int(e.Flags)
	//flag e.Out
	x.Int(e.ID)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	x.Int(e.Date)
	if e.Flags&512 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&128 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e PhotosPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosPhotos)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e PhotosPhotosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosPhotosSlice)
	x.Int(e.Count)
	x.Vector(e.Photos)
	x.Vector(e.Users)
	return x.buf
}

func (e PhotosPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosPhoto)
	x.Bytes(e.Photo.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e UploadFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadFile)
	x.Bytes(e.Type.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e UploadFileCdnRedirect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadFileCdnRedirect)
	x.Int(e.DcID)
	x.StringBytes(e.FileToken)
	x.StringBytes(e.EncryptionKey)
	x.StringBytes(e.EncryptionIv)
	x.Vector(e.FileHashes)
	return x.buf
}

func (e DcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDcOption)
	x.Int(e.Flags)
	//flag e.Ipv6
	//flag e.MediaOnly
	//flag e.TcpoOnly
	//flag e.Cdn
	//flag e.Static
	x.Int(e.ID)
	x.String(e.IpAddress)
	x.Int(e.Port)
	if e.Flags&1024 != 0 {
		x.StringBytes(e.Secret)
	}
	return x.buf
}

func (e Config) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcConfig)
	x.Int(e.Flags)
	//flag e.PhonecallsEnabled
	//flag e.DefaultP2pContacts
	//flag e.PreloadFeaturedStickers
	//flag e.IgnorePhoneEntities
	//flag e.RevokePmInbox
	//flag e.BlockedMode
	//flag e.PfsEnabled
	x.Int(e.Date)
	x.Int(e.Expires)
	x.Bytes(e.TestMode.encode())
	x.Int(e.ThisDc)
	x.Vector(e.DcOptions)
	x.String(e.DcTxtDomainName)
	x.Int(e.ChatSizeMax)
	x.Int(e.MegagroupSizeMax)
	x.Int(e.ForwardedCountMax)
	x.Int(e.OnlineUpdatePeriodMs)
	x.Int(e.OfflineBlurTimeoutMs)
	x.Int(e.OfflineIdleTimeoutMs)
	x.Int(e.OnlineCloudTimeoutMs)
	x.Int(e.NotifyCloudDelayMs)
	x.Int(e.NotifyDefaultDelayMs)
	x.Int(e.PushChatPeriodMs)
	x.Int(e.PushChatLimit)
	x.Int(e.SavedGifsLimit)
	x.Int(e.EditTimeLimit)
	x.Int(e.RevokeTimeLimit)
	x.Int(e.RevokePmTimeLimit)
	x.Int(e.RatingEDecay)
	x.Int(e.StickersRecentLimit)
	x.Int(e.StickersFavedLimit)
	x.Int(e.ChannelsReadMediaPeriod)
	if e.Flags&1 != 0 {
		x.Int(e.TmpSessions)
	}
	x.Int(e.PinnedDialogsCountMax)
	x.Int(e.PinnedInfolderCountMax)
	x.Int(e.CallReceiveTimeoutMs)
	x.Int(e.CallRingTimeoutMs)
	x.Int(e.CallConnectTimeoutMs)
	x.Int(e.CallPacketTimeoutMs)
	x.String(e.MeUrlPrefix)
	if e.Flags&128 != 0 {
		x.String(e.AutoupdateUrlPrefix)
	}
	if e.Flags&512 != 0 {
		x.String(e.GifSearchUsername)
	}
	if e.Flags&1024 != 0 {
		x.String(e.VenueSearchUsername)
	}
	if e.Flags&2048 != 0 {
		x.String(e.ImgSearchUsername)
	}
	if e.Flags&4096 != 0 {
		x.String(e.StaticMapsProvider)
	}
	x.Int(e.CaptionLengthMax)
	x.Int(e.MessageLengthMax)
	x.Int(e.WebfileDcID)
	if e.Flags&4 != 0 {
		x.String(e.SuggestedLangCode)
	}
	if e.Flags&4 != 0 {
		x.Int(e.LangPackVersion)
	}
	if e.Flags&4 != 0 {
		x.Int(e.BaseLangPackVersion)
	}
	return x.buf
}

func (e NearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcNearestDc)
	x.String(e.Country)
	x.Int(e.ThisDc)
	x.Int(e.NearestDc)
	return x.buf
}

func (e HelpAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpAppUpdate)
	x.Int(e.Flags)
	//flag e.CanNotSkip
	x.Int(e.ID)
	x.String(e.Version)
	x.String(e.Text)
	x.Vector(e.Entities)
	if e.Flags&2 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&4 != 0 {
		x.String(e.Url)
	}
	return x.buf
}

func (e HelpNoAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpNoAppUpdate)
	return x.buf
}

func (e HelpInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpInviteText)
	x.String(e.Message)
	return x.buf
}

func (e EncryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedChatEmpty)
	x.Int(e.ID)
	return x.buf
}

func (e EncryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedChatWaiting)
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	return x.buf
}

func (e EncryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedChatRequested)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GA)
	return x.buf
}

func (e EncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedChat)
	x.Int(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	return x.buf
}

func (e EncryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedChatDiscarded)
	x.Int(e.ID)
	return x.buf
}

func (e InputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputEncryptedChat)
	x.Int(e.ChatID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e EncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedFileEmpty)
	return x.buf
}

func (e EncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.Int(e.DcID)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e InputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputEncryptedFileEmpty)
	return x.buf
}

func (e InputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputEncryptedFileUploaded)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Md5Checksum)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e InputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputEncryptedFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputEncryptedFileBigUploaded)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.Int(e.KeyFingerprint)
	return x.buf
}

func (e EncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedMessage)
	x.Long(e.RandomID)
	x.Int(e.ChatID)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e EncryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEncryptedMessageService)
	x.Long(e.RandomID)
	x.Int(e.ChatID)
	x.Int(e.Date)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e MessagesDhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDhConfigNotModified)
	x.StringBytes(e.Random)
	return x.buf
}

func (e MessagesDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDhConfig)
	x.Int(e.G)
	x.StringBytes(e.P)
	x.Int(e.Version)
	x.StringBytes(e.Random)
	return x.buf
}

func (e MessagesSentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSentEncryptedMessage)
	x.Int(e.Date)
	return x.buf
}

func (e MessagesSentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSentEncryptedFile)
	x.Int(e.Date)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e InputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputDocumentEmpty)
	return x.buf
}

func (e InputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputDocument)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	return x.buf
}

func (e DocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e Document) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocument)
	x.Int(e.Flags)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.StringBytes(e.FileReference)
	x.Int(e.Date)
	x.String(e.MimeType)
	x.Int(e.Size)
	if e.Flags&1 != 0 {
		x.Vector(e.Thumbs)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.VideoThumbs)
	}
	x.Int(e.DcID)
	x.Vector(e.Attributes)
	return x.buf
}

func (e HelpSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpSupport)
	x.String(e.PhoneNumber)
	x.Bytes(e.User.encode())
	return x.buf
}

func (e NotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcNotifyPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e NotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcNotifyUsers)
	return x.buf
}

func (e NotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcNotifyChats)
	return x.buf
}

func (e NotifyBroadcasts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcNotifyBroadcasts)
	return x.buf
}

func (e SendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageTypingAction)
	return x.buf
}

func (e SendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageCancelAction)
	return x.buf
}

func (e SendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageRecordVideoAction)
	return x.buf
}

func (e SendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageUploadVideoAction)
	x.Int(e.Progress)
	return x.buf
}

func (e SendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageRecordAudioAction)
	return x.buf
}

func (e SendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageUploadAudioAction)
	x.Int(e.Progress)
	return x.buf
}

func (e SendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageUploadPhotoAction)
	x.Int(e.Progress)
	return x.buf
}

func (e SendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageUploadDocumentAction)
	x.Int(e.Progress)
	return x.buf
}

func (e SendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageGeoLocationAction)
	return x.buf
}

func (e SendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageChooseContactAction)
	return x.buf
}

func (e SendMessageGamePlayAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageGamePlayAction)
	return x.buf
}

func (e SendMessageRecordRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageRecordRoundAction)
	return x.buf
}

func (e SendMessageUploadRoundAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSendMessageUploadRoundAction)
	x.Int(e.Progress)
	return x.buf
}

func (e ContactsFound) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsFound)
	x.Vector(e.MyResults)
	x.Vector(e.Results)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e InputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e InputPrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyChatInvite)
	return x.buf
}

func (e InputPrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyPhoneCall)
	return x.buf
}

func (e InputPrivacyKeyPhoneP2P) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyPhoneP2P)
	return x.buf
}

func (e InputPrivacyKeyForwards) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyForwards)
	return x.buf
}

func (e InputPrivacyKeyProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyProfilePhoto)
	return x.buf
}

func (e InputPrivacyKeyPhoneNumber) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyPhoneNumber)
	return x.buf
}

func (e InputPrivacyKeyAddedByPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyKeyAddedByPhone)
	return x.buf
}

func (e PrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e PrivacyKeyChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyChatInvite)
	return x.buf
}

func (e PrivacyKeyPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyPhoneCall)
	return x.buf
}

func (e PrivacyKeyPhoneP2P) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyPhoneP2P)
	return x.buf
}

func (e PrivacyKeyForwards) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyForwards)
	return x.buf
}

func (e PrivacyKeyProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyProfilePhoto)
	return x.buf
}

func (e PrivacyKeyPhoneNumber) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyPhoneNumber)
	return x.buf
}

func (e PrivacyKeyAddedByPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyKeyAddedByPhone)
	return x.buf
}

func (e InputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueAllowContacts)
	return x.buf
}

func (e InputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueAllowAll)
	return x.buf
}

func (e InputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueAllowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e InputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueDisallowContacts)
	return x.buf
}

func (e InputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueDisallowAll)
	return x.buf
}

func (e InputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueDisallowUsers)
	x.Vector(e.Users)
	return x.buf
}

func (e InputPrivacyValueAllowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueAllowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e InputPrivacyValueDisallowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPrivacyValueDisallowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e PrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueAllowContacts)
	return x.buf
}

func (e PrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueAllowAll)
	return x.buf
}

func (e PrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueAllowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e PrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueDisallowContacts)
	return x.buf
}

func (e PrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueDisallowAll)
	return x.buf
}

func (e PrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueDisallowUsers)
	x.VectorInt(e.Users)
	return x.buf
}

func (e PrivacyValueAllowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueAllowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e PrivacyValueDisallowChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPrivacyValueDisallowChatParticipants)
	x.VectorInt(e.Chats)
	return x.buf
}

func (e AccountPrivacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountPrivacyRules)
	x.Vector(e.Rules)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e AccountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountDaysTTL)
	x.Int(e.Days)
	return x.buf
}

func (e DocumentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeImageSize)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e DocumentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeAnimated)
	return x.buf
}

func (e DocumentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeSticker)
	x.Int(e.Flags)
	//flag e.Mask
	x.String(e.Alt)
	x.Bytes(e.Stickerset.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.MaskCoords.encode())
	}
	return x.buf
}

func (e DocumentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeVideo)
	x.Int(e.Flags)
	//flag e.RoundMessage
	//flag e.SupportsStreaming
	x.Int(e.Duration)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e DocumentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeAudio)
	x.Int(e.Flags)
	//flag e.Voice
	x.Int(e.Duration)
	if e.Flags&1 != 0 {
		x.String(e.Title)
	}
	if e.Flags&2 != 0 {
		x.String(e.Performer)
	}
	if e.Flags&4 != 0 {
		x.StringBytes(e.Waveform)
	}
	return x.buf
}

func (e DocumentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeFilename)
	x.String(e.FileName)
	return x.buf
}

func (e DocumentAttributeHasStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDocumentAttributeHasStickers)
	return x.buf
}

func (e MessagesStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesStickersNotModified)
	return x.buf
}

func (e MessagesStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesStickers)
	x.Int(e.Hash)
	x.Vector(e.Stickers)
	return x.buf
}

func (e StickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickerPack)
	x.String(e.Emoticon)
	x.VectorLong(e.Documents)
	return x.buf
}

func (e MessagesAllStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAllStickersNotModified)
	return x.buf
}

func (e MessagesAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAllStickers)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.buf
}

func (e MessagesAffectedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAffectedMessages)
	x.Int(e.Pts)
	x.Int(e.PtsCount)
	return x.buf
}

func (e WebPageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebPageEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e WebPagePending) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebPagePending)
	x.Long(e.ID)
	x.Int(e.Date)
	return x.buf
}

func (e WebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebPage)
	x.Int(e.Flags)
	x.Long(e.ID)
	x.String(e.Url)
	x.String(e.DisplayUrl)
	x.Int(e.Hash)
	if e.Flags&1 != 0 {
		x.String(e.Type)
	}
	if e.Flags&2 != 0 {
		x.String(e.SiteName)
	}
	if e.Flags&4 != 0 {
		x.String(e.Title)
	}
	if e.Flags&8 != 0 {
		x.String(e.Description)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&32 != 0 {
		x.String(e.EmbedUrl)
	}
	if e.Flags&32 != 0 {
		x.String(e.EmbedType)
	}
	if e.Flags&64 != 0 {
		x.Int(e.EmbedWidth)
	}
	if e.Flags&64 != 0 {
		x.Int(e.EmbedHeight)
	}
	if e.Flags&128 != 0 {
		x.Int(e.Duration)
	}
	if e.Flags&256 != 0 {
		x.String(e.Author)
	}
	if e.Flags&512 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&1024 != 0 {
		x.Bytes(e.CachedPage.encode())
	}
	if e.Flags&4096 != 0 {
		x.Vector(e.Attributes)
	}
	return x.buf
}

func (e WebPageNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebPageNotModified)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.CachedPageViews)
	}
	return x.buf
}

func (e Authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthorization)
	x.Int(e.Flags)
	//flag e.Current
	//flag e.OfficialApp
	//flag e.PasswordPending
	x.Long(e.Hash)
	x.String(e.DeviceModel)
	x.String(e.Platform)
	x.String(e.SystemVersion)
	x.Int(e.ApiID)
	x.String(e.AppName)
	x.String(e.AppVersion)
	x.Int(e.DateCreated)
	x.Int(e.DateActive)
	x.String(e.Ip)
	x.String(e.Country)
	x.String(e.Region)
	return x.buf
}

func (e AccountAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountAuthorizations)
	x.Vector(e.Authorizations)
	return x.buf
}

func (e AccountPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountPassword)
	x.Int(e.Flags)
	//flag e.HasRecovery
	//flag e.HasSecureValues
	//flag e.HasPassword
	if e.Flags&4 != 0 {
		x.Bytes(e.CurrentAlgo.encode())
	}
	if e.Flags&4 != 0 {
		x.StringBytes(e.SrpB)
	}
	if e.Flags&4 != 0 {
		x.Long(e.SrpID)
	}
	if e.Flags&8 != 0 {
		x.String(e.Hint)
	}
	if e.Flags&16 != 0 {
		x.String(e.EmailUnconfirmedPattern)
	}
	x.Bytes(e.NewAlgo.encode())
	x.Bytes(e.NewSecureAlgo.encode())
	x.StringBytes(e.SecureRandom)
	return x.buf
}

func (e AccountPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountPasswordSettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.Email)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.SecureSettings.encode())
	}
	return x.buf
}

func (e AccountPasswordInputSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountPasswordInputSettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.NewAlgo.encode())
	}
	if e.Flags&1 != 0 {
		x.StringBytes(e.NewPasswordHash)
	}
	if e.Flags&1 != 0 {
		x.String(e.Hint)
	}
	if e.Flags&2 != 0 {
		x.String(e.Email)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.NewSecureSettings.encode())
	}
	return x.buf
}

func (e AuthPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthPasswordRecovery)
	x.String(e.EmailPattern)
	return x.buf
}

func (e ReceivedNotifyMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReceivedNotifyMessage)
	x.Int(e.ID)
	x.Int(e.Flags)
	return x.buf
}

func (e ChatInviteEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatInviteEmpty)
	return x.buf
}

func (e ChatInviteExported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatInviteExported)
	x.String(e.Link)
	return x.buf
}

func (e ChatInviteAlready) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatInviteAlready)
	x.Bytes(e.Chat.encode())
	return x.buf
}

func (e ChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatInvite)
	x.Int(e.Flags)
	//flag e.Channel
	//flag e.Broadcast
	//flag e.Public
	//flag e.Megagroup
	x.String(e.Title)
	x.Bytes(e.Photo.encode())
	x.Int(e.ParticipantsCount)
	if e.Flags&16 != 0 {
		x.Vector(e.Participants)
	}
	return x.buf
}

func (e ChatInvitePeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatInvitePeek)
	x.Bytes(e.Chat.encode())
	x.Int(e.Expires)
	return x.buf
}

func (e InputStickerSetEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetEmpty)
	return x.buf
}

func (e InputStickerSetID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetID)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputStickerSetShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetShortName)
	x.String(e.ShortName)
	return x.buf
}

func (e InputStickerSetAnimatedEmoji) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetAnimatedEmoji)
	return x.buf
}

func (e InputStickerSetDice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetDice)
	x.String(e.Emoticon)
	return x.buf
}

func (e StickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickerSet)
	x.Int(e.Flags)
	//flag e.Archived
	//flag e.Official
	//flag e.Masks
	//flag e.Animated
	if e.Flags&1 != 0 {
		x.Int(e.InstalledDate)
	}
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.String(e.Title)
	x.String(e.ShortName)
	if e.Flags&16 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	if e.Flags&16 != 0 {
		x.Int(e.ThumbDcID)
	}
	x.Int(e.Count)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesStickerSet)
	x.Bytes(e.Set.encode())
	x.Vector(e.Packs)
	x.Vector(e.Documents)
	return x.buf
}

func (e BotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotCommand)
	x.String(e.Command)
	x.String(e.Description)
	return x.buf
}

func (e BotInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInfo)
	x.Int(e.UserID)
	x.String(e.Description)
	x.Vector(e.Commands)
	return x.buf
}

func (e KeyboardButton) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButton)
	x.String(e.Text)
	return x.buf
}

func (e KeyboardButtonUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonUrl)
	x.String(e.Text)
	x.String(e.Url)
	return x.buf
}

func (e KeyboardButtonCallback) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonCallback)
	x.String(e.Text)
	x.StringBytes(e.Data)
	return x.buf
}

func (e KeyboardButtonRequestPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonRequestPhone)
	x.String(e.Text)
	return x.buf
}

func (e KeyboardButtonRequestGeoLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonRequestGeoLocation)
	x.String(e.Text)
	return x.buf
}

func (e KeyboardButtonSwitchInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonSwitchInline)
	x.Int(e.Flags)
	//flag e.SamePeer
	x.String(e.Text)
	x.String(e.Query)
	return x.buf
}

func (e KeyboardButtonGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonGame)
	x.String(e.Text)
	return x.buf
}

func (e KeyboardButtonBuy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonBuy)
	x.String(e.Text)
	return x.buf
}

func (e KeyboardButtonUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonUrlAuth)
	x.Int(e.Flags)
	x.String(e.Text)
	if e.Flags&1 != 0 {
		x.String(e.FwdText)
	}
	x.String(e.Url)
	x.Int(e.ButtonID)
	return x.buf
}

func (e InputKeyboardButtonUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputKeyboardButtonUrlAuth)
	x.Int(e.Flags)
	//flag e.RequestWriteAccess
	x.String(e.Text)
	if e.Flags&2 != 0 {
		x.String(e.FwdText)
	}
	x.String(e.Url)
	x.Bytes(e.Bot.encode())
	return x.buf
}

func (e KeyboardButtonRequestPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonRequestPoll)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.Quiz.encode())
	}
	x.String(e.Text)
	return x.buf
}

func (e KeyboardButtonRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcKeyboardButtonRow)
	x.Vector(e.Buttons)
	return x.buf
}

func (e ReplyKeyboardHide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReplyKeyboardHide)
	x.Int(e.Flags)
	//flag e.Selective
	return x.buf
}

func (e ReplyKeyboardForceReply) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReplyKeyboardForceReply)
	x.Int(e.Flags)
	//flag e.SingleUse
	//flag e.Selective
	return x.buf
}

func (e ReplyKeyboardMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReplyKeyboardMarkup)
	x.Int(e.Flags)
	//flag e.Resize
	//flag e.SingleUse
	//flag e.Selective
	x.Vector(e.Rows)
	return x.buf
}

func (e ReplyInlineMarkup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcReplyInlineMarkup)
	x.Vector(e.Rows)
	return x.buf
}

func (e MessageEntityUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityUnknown)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityMention) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityMention)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityHashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityHashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityBotCommand) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityBotCommand)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityEmail)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityBold)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityItalic)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityCode)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityPre) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityPre)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Language)
	return x.buf
}

func (e MessageEntityTextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityTextUrl)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.String(e.Url)
	return x.buf
}

func (e MessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Int(e.UserID)
	return x.buf
}

func (e InputMessageEntityMentionName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessageEntityMentionName)
	x.Int(e.Offset)
	x.Int(e.Length)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e MessageEntityPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityPhone)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityCashtag) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityCashtag)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityUnderline)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityStrike)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityBlockquote)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e MessageEntityBankCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageEntityBankCard)
	x.Int(e.Offset)
	x.Int(e.Length)
	return x.buf
}

func (e InputChannelEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputChannelEmpty)
	return x.buf
}

func (e InputChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputChannel)
	x.Int(e.ChannelID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputChannelFromMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputChannelFromMessage)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ChannelID)
	return x.buf
}

func (e ContactsResolvedPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsResolvedPeer)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e MessageRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageRange)
	x.Int(e.MinID)
	x.Int(e.MaxID)
	return x.buf
}

func (e UpdatesChannelDifferenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesChannelDifferenceEmpty)
	x.Int(e.Flags)
	//flag e.Final
	x.Int(e.Pts)
	if e.Flags&2 != 0 {
		x.Int(e.Timeout)
	}
	return x.buf
}

func (e UpdatesChannelDifferenceTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesChannelDifferenceTooLong)
	x.Int(e.Flags)
	//flag e.Final
	if e.Flags&2 != 0 {
		x.Int(e.Timeout)
	}
	x.Bytes(e.Dialog.encode())
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e UpdatesChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesChannelDifference)
	x.Int(e.Flags)
	//flag e.Final
	x.Int(e.Pts)
	if e.Flags&2 != 0 {
		x.Int(e.Timeout)
	}
	x.Vector(e.NewMessages)
	x.Vector(e.OtherUpdates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e ChannelMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelMessagesFilterEmpty)
	return x.buf
}

func (e ChannelMessagesFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelMessagesFilter)
	x.Int(e.Flags)
	//flag e.ExcludeNewMessages
	x.Vector(e.Ranges)
	return x.buf
}

func (e ChannelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipant)
	x.Int(e.UserID)
	x.Int(e.Date)
	return x.buf
}

func (e ChannelParticipantSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantSelf)
	x.Int(e.UserID)
	x.Int(e.InviterID)
	x.Int(e.Date)
	return x.buf
}

func (e ChannelParticipantCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantCreator)
	x.Int(e.Flags)
	x.Int(e.UserID)
	if e.Flags&1 != 0 {
		x.String(e.Rank)
	}
	return x.buf
}

func (e ChannelParticipantAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantAdmin)
	x.Int(e.Flags)
	//flag e.CanEdit
	//flag e.Self
	x.Int(e.UserID)
	if e.Flags&2 != 0 {
		x.Int(e.InviterID)
	}
	x.Int(e.PromotedBy)
	x.Int(e.Date)
	x.Bytes(e.AdminRights.encode())
	if e.Flags&4 != 0 {
		x.String(e.Rank)
	}
	return x.buf
}

func (e ChannelParticipantBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantBanned)
	x.Int(e.Flags)
	//flag e.Left
	x.Int(e.UserID)
	x.Int(e.KickedBy)
	x.Int(e.Date)
	x.Bytes(e.BannedRights.encode())
	return x.buf
}

func (e ChannelParticipantsRecent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsRecent)
	return x.buf
}

func (e ChannelParticipantsAdmins) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsAdmins)
	return x.buf
}

func (e ChannelParticipantsKicked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsKicked)
	x.String(e.Q)
	return x.buf
}

func (e ChannelParticipantsBots) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsBots)
	return x.buf
}

func (e ChannelParticipantsBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsBanned)
	x.String(e.Q)
	return x.buf
}

func (e ChannelParticipantsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsSearch)
	x.String(e.Q)
	return x.buf
}

func (e ChannelParticipantsContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelParticipantsContacts)
	x.String(e.Q)
	return x.buf
}

func (e ChannelsChannelParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsChannelParticipants)
	x.Int(e.Count)
	x.Vector(e.Participants)
	x.Vector(e.Users)
	return x.buf
}

func (e ChannelsChannelParticipantsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsChannelParticipantsNotModified)
	return x.buf
}

func (e ChannelsChannelParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsChannelParticipant)
	x.Bytes(e.Participant.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e HelpTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpTermsOfService)
	x.Int(e.Flags)
	//flag e.Popup
	x.Bytes(e.ID.encode())
	x.String(e.Text)
	x.Vector(e.Entities)
	if e.Flags&2 != 0 {
		x.Int(e.MinAgeConfirm)
	}
	return x.buf
}

func (e MessagesSavedGifsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSavedGifsNotModified)
	return x.buf
}

func (e MessagesSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSavedGifs)
	x.Int(e.Hash)
	x.Vector(e.Gifs)
	return x.buf
}

func (e InputBotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e InputBotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageText)
	x.Int(e.Flags)
	//flag e.NoWebpage
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e InputBotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.GeoPoint.encode())
	x.Int(e.Period)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e InputBotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e InputBotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e InputBotInlineMessageGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageGame)
	x.Int(e.Flags)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e InputBotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineResult)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.String(e.Description)
	}
	if e.Flags&8 != 0 {
		x.String(e.Url)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.Content.encode())
	}
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e InputBotInlineResultPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineResultPhoto)
	x.String(e.ID)
	x.String(e.Type)
	x.Bytes(e.Photo.encode())
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e InputBotInlineResultDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineResultDocument)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.String(e.Description)
	}
	x.Bytes(e.Document.encode())
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e InputBotInlineResultGame) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineResultGame)
	x.String(e.ID)
	x.String(e.ShortName)
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e BotInlineMessageMediaAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineMessageMediaAuto)
	x.Int(e.Flags)
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e BotInlineMessageText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineMessageText)
	x.Int(e.Flags)
	//flag e.NoWebpage
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e BotInlineMessageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineMessageMediaGeo)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.Int(e.Period)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e BotInlineMessageMediaVenue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineMessageMediaVenue)
	x.Int(e.Flags)
	x.Bytes(e.Geo.encode())
	x.String(e.Title)
	x.String(e.Address)
	x.String(e.Provider)
	x.String(e.VenueID)
	x.String(e.VenueType)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e BotInlineMessageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineMessageMediaContact)
	x.Int(e.Flags)
	x.String(e.PhoneNumber)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Vcard)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	return x.buf
}

func (e BotInlineResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineResult)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.String(e.Description)
	}
	if e.Flags&8 != 0 {
		x.String(e.Url)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.Content.encode())
	}
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e BotInlineMediaResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotInlineMediaResult)
	x.Int(e.Flags)
	x.String(e.ID)
	x.String(e.Type)
	if e.Flags&1 != 0 {
		x.Bytes(e.Photo.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&4 != 0 {
		x.String(e.Title)
	}
	if e.Flags&8 != 0 {
		x.String(e.Description)
	}
	x.Bytes(e.SendMessage.encode())
	return x.buf
}

func (e MessagesBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesBotResults)
	x.Int(e.Flags)
	//flag e.Gallery
	x.Long(e.QueryID)
	if e.Flags&2 != 0 {
		x.String(e.NextOffset)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.SwitchPm.encode())
	}
	x.Vector(e.Results)
	x.Int(e.CacheTime)
	x.Vector(e.Users)
	return x.buf
}

func (e ExportedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcExportedMessageLink)
	x.String(e.Link)
	x.String(e.Html)
	return x.buf
}

func (e MessageFwdHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageFwdHeader)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FromID)
	}
	if e.Flags&32 != 0 {
		x.String(e.FromName)
	}
	x.Int(e.Date)
	if e.Flags&2 != 0 {
		x.Int(e.ChannelID)
	}
	if e.Flags&4 != 0 {
		x.Int(e.ChannelPost)
	}
	if e.Flags&8 != 0 {
		x.String(e.PostAuthor)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.SavedFromPeer.encode())
	}
	if e.Flags&16 != 0 {
		x.Int(e.SavedFromMsgID)
	}
	if e.Flags&64 != 0 {
		x.String(e.PsaType)
	}
	return x.buf
}

func (e AuthCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthCodeTypeSms)
	return x.buf
}

func (e AuthCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthCodeTypeCall)
	return x.buf
}

func (e AuthCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthCodeTypeFlashCall)
	return x.buf
}

func (e AuthSentCodeTypeApp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSentCodeTypeApp)
	x.Int(e.Length)
	return x.buf
}

func (e AuthSentCodeTypeSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSentCodeTypeSms)
	x.Int(e.Length)
	return x.buf
}

func (e AuthSentCodeTypeCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSentCodeTypeCall)
	x.Int(e.Length)
	return x.buf
}

func (e AuthSentCodeTypeFlashCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSentCodeTypeFlashCall)
	x.String(e.Pattern)
	return x.buf
}

func (e MessagesBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesBotCallbackAnswer)
	x.Int(e.Flags)
	//flag e.Alert
	//flag e.HasUrl
	//flag e.NativeUi
	if e.Flags&1 != 0 {
		x.String(e.Message)
	}
	if e.Flags&4 != 0 {
		x.String(e.Url)
	}
	x.Int(e.CacheTime)
	return x.buf
}

func (e MessagesMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesMessageEditData)
	x.Int(e.Flags)
	//flag e.Caption
	return x.buf
}

func (e InputBotInlineMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputBotInlineMessageID)
	x.Int(e.DcID)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InlineBotSwitchPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInlineBotSwitchPM)
	x.String(e.Text)
	x.String(e.StartParam)
	return x.buf
}

func (e MessagesPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesPeerDialogs)
	x.Vector(e.Dialogs)
	x.Vector(e.Messages)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	x.Bytes(e.State.encode())
	return x.buf
}

func (e TopPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeer)
	x.Bytes(e.Peer.encode())
	x.Double(e.Rating)
	return x.buf
}

func (e TopPeerCategoryBotsPM) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryBotsPM)
	return x.buf
}

func (e TopPeerCategoryBotsInline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryBotsInline)
	return x.buf
}

func (e TopPeerCategoryCorrespondents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryCorrespondents)
	return x.buf
}

func (e TopPeerCategoryGroups) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryGroups)
	return x.buf
}

func (e TopPeerCategoryChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryChannels)
	return x.buf
}

func (e TopPeerCategoryPhoneCalls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryPhoneCalls)
	return x.buf
}

func (e TopPeerCategoryForwardUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryForwardUsers)
	return x.buf
}

func (e TopPeerCategoryForwardChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryForwardChats)
	return x.buf
}

func (e TopPeerCategoryPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTopPeerCategoryPeers)
	x.Bytes(e.Category.encode())
	x.Int(e.Count)
	x.Vector(e.Peers)
	return x.buf
}

func (e ContactsTopPeersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsTopPeersNotModified)
	return x.buf
}

func (e ContactsTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsTopPeers)
	x.Vector(e.Categories)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e ContactsTopPeersDisabled) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsTopPeersDisabled)
	return x.buf
}

func (e DraftMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDraftMessageEmpty)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.Date)
	}
	return x.buf
}

func (e DraftMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDraftMessage)
	x.Int(e.Flags)
	//flag e.NoWebpage
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.String(e.Message)
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	x.Int(e.Date)
	return x.buf
}

func (e MessagesFeaturedStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFeaturedStickersNotModified)
	x.Int(e.Count)
	return x.buf
}

func (e MessagesFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFeaturedStickers)
	x.Int(e.Hash)
	x.Int(e.Count)
	x.Vector(e.Sets)
	x.VectorLong(e.Unread)
	return x.buf
}

func (e MessagesRecentStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesRecentStickersNotModified)
	return x.buf
}

func (e MessagesRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesRecentStickers)
	x.Int(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Stickers)
	x.VectorInt(e.Dates)
	return x.buf
}

func (e MessagesArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesArchivedStickers)
	x.Int(e.Count)
	x.Vector(e.Sets)
	return x.buf
}

func (e MessagesStickerSetInstallResultSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesStickerSetInstallResultSuccess)
	return x.buf
}

func (e MessagesStickerSetInstallResultArchive) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesStickerSetInstallResultArchive)
	x.Vector(e.Sets)
	return x.buf
}

func (e StickerSetCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickerSetCovered)
	x.Bytes(e.Set.encode())
	x.Bytes(e.Cover.encode())
	return x.buf
}

func (e StickerSetMultiCovered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickerSetMultiCovered)
	x.Bytes(e.Set.encode())
	x.Vector(e.Covers)
	return x.buf
}

func (e MaskCoords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMaskCoords)
	x.Int(e.N)
	x.Double(e.X)
	x.Double(e.Y)
	x.Double(e.Zoom)
	return x.buf
}

func (e InputStickeredMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickeredMediaPhoto)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e InputStickeredMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickeredMediaDocument)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e Game) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcGame)
	x.Int(e.Flags)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.String(e.ShortName)
	x.String(e.Title)
	x.String(e.Description)
	x.Bytes(e.Photo.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Document.encode())
	}
	return x.buf
}

func (e InputGameID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputGameID)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputGameShortName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputGameShortName)
	x.Bytes(e.BotID.encode())
	x.String(e.ShortName)
	return x.buf
}

func (e HighScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHighScore)
	x.Int(e.Pos)
	x.Int(e.UserID)
	x.Int(e.Score)
	return x.buf
}

func (e MessagesHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesHighScores)
	x.Vector(e.Scores)
	x.Vector(e.Users)
	return x.buf
}

func (e TextEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextEmpty)
	return x.buf
}

func (e TextPlain) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextPlain)
	x.String(e.Text)
	return x.buf
}

func (e TextBold) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextBold)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextItalic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextItalic)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextUnderline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextUnderline)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextStrike) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextStrike)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextFixed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextFixed)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextUrl)
	x.Bytes(e.Text.encode())
	x.String(e.Url)
	x.Long(e.WebpageID)
	return x.buf
}

func (e TextEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextEmail)
	x.Bytes(e.Text.encode())
	x.String(e.Email)
	return x.buf
}

func (e TextConcat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextConcat)
	x.Vector(e.Texts)
	return x.buf
}

func (e TextSubscript) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextSubscript)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextSuperscript) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextSuperscript)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextMarked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextMarked)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e TextPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextPhone)
	x.Bytes(e.Text.encode())
	x.String(e.Phone)
	return x.buf
}

func (e TextImage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextImage)
	x.Long(e.DocumentID)
	x.Int(e.W)
	x.Int(e.H)
	return x.buf
}

func (e TextAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTextAnchor)
	x.Bytes(e.Text.encode())
	x.String(e.Name)
	return x.buf
}

func (e PageBlockUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockUnsupported)
	return x.buf
}

func (e PageBlockTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockTitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockSubtitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockSubtitle)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockAuthorDate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockAuthorDate)
	x.Bytes(e.Author.encode())
	x.Int(e.PublishedDate)
	return x.buf
}

func (e PageBlockHeader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockHeader)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockSubheader) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockSubheader)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockParagraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockParagraph)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockPreformatted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockPreformatted)
	x.Bytes(e.Text.encode())
	x.String(e.Language)
	return x.buf
}

func (e PageBlockFooter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockFooter)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockDivider) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockDivider)
	return x.buf
}

func (e PageBlockAnchor) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockAnchor)
	x.String(e.Name)
	return x.buf
}

func (e PageBlockList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockList)
	x.Vector(e.Items)
	return x.buf
}

func (e PageBlockBlockquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockBlockquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockPullquote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockPullquote)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockPhoto)
	x.Int(e.Flags)
	x.Long(e.PhotoID)
	x.Bytes(e.Caption.encode())
	if e.Flags&1 != 0 {
		x.String(e.Url)
	}
	if e.Flags&1 != 0 {
		x.Long(e.WebpageID)
	}
	return x.buf
}

func (e PageBlockVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockVideo)
	x.Int(e.Flags)
	//flag e.Autoplay
	//flag e.Loop
	x.Long(e.VideoID)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockCover) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockCover)
	x.Bytes(e.Cover.encode())
	return x.buf
}

func (e PageBlockEmbed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockEmbed)
	x.Int(e.Flags)
	//flag e.FullWidth
	//flag e.AllowScrolling
	if e.Flags&2 != 0 {
		x.String(e.Url)
	}
	if e.Flags&4 != 0 {
		x.String(e.Html)
	}
	if e.Flags&16 != 0 {
		x.Long(e.PosterPhotoID)
	}
	if e.Flags&32 != 0 {
		x.Int(e.W)
	}
	if e.Flags&32 != 0 {
		x.Int(e.H)
	}
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockEmbedPost) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockEmbedPost)
	x.String(e.Url)
	x.Long(e.WebpageID)
	x.Long(e.AuthorPhotoID)
	x.String(e.Author)
	x.Int(e.Date)
	x.Vector(e.Blocks)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockCollage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockCollage)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockSlideshow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockSlideshow)
	x.Vector(e.Items)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e PageBlockAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockAudio)
	x.Long(e.AudioID)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PageBlockKicker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockKicker)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageBlockTable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockTable)
	x.Int(e.Flags)
	//flag e.Bordered
	//flag e.Striped
	x.Bytes(e.Title.encode())
	x.Vector(e.Rows)
	return x.buf
}

func (e PageBlockOrderedList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockOrderedList)
	x.Vector(e.Items)
	return x.buf
}

func (e PageBlockDetails) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockDetails)
	x.Int(e.Flags)
	//flag e.Open
	x.Vector(e.Blocks)
	x.Bytes(e.Title.encode())
	return x.buf
}

func (e PageBlockRelatedArticles) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockRelatedArticles)
	x.Bytes(e.Title.encode())
	x.Vector(e.Articles)
	return x.buf
}

func (e PageBlockMap) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageBlockMap)
	x.Bytes(e.Geo.encode())
	x.Int(e.Zoom)
	x.Int(e.W)
	x.Int(e.H)
	x.Bytes(e.Caption.encode())
	return x.buf
}

func (e PhoneCallDiscardReasonMissed) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallDiscardReasonMissed)
	return x.buf
}

func (e PhoneCallDiscardReasonDisconnect) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallDiscardReasonDisconnect)
	return x.buf
}

func (e PhoneCallDiscardReasonHangup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallDiscardReasonHangup)
	return x.buf
}

func (e PhoneCallDiscardReasonBusy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallDiscardReasonBusy)
	return x.buf
}

func (e DataJSON) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDataJSON)
	x.String(e.Data)
	return x.buf
}

func (e LabeledPrice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLabeledPrice)
	x.String(e.Label)
	x.Long(e.Amount)
	return x.buf
}

func (e Invoice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvoice)
	x.Int(e.Flags)
	//flag e.Test
	//flag e.NameRequested
	//flag e.PhoneRequested
	//flag e.EmailRequested
	//flag e.ShippingAddressRequested
	//flag e.Flexible
	//flag e.PhoneToProvider
	//flag e.EmailToProvider
	x.String(e.Currency)
	x.Vector(e.Prices)
	return x.buf
}

func (e PaymentCharge) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentCharge)
	x.String(e.ID)
	x.String(e.ProviderChargeID)
	return x.buf
}

func (e PostAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPostAddress)
	x.String(e.StreetLine1)
	x.String(e.StreetLine2)
	x.String(e.City)
	x.String(e.State)
	x.String(e.CountryIso2)
	x.String(e.PostCode)
	return x.buf
}

func (e PaymentRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentRequestedInfo)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.Name)
	}
	if e.Flags&2 != 0 {
		x.String(e.Phone)
	}
	if e.Flags&4 != 0 {
		x.String(e.Email)
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.ShippingAddress.encode())
	}
	return x.buf
}

func (e PaymentSavedCredentialsCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentSavedCredentialsCard)
	x.String(e.ID)
	x.String(e.Title)
	return x.buf
}

func (e WebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebDocument)
	x.String(e.Url)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e WebDocumentNoProxy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebDocumentNoProxy)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e InputWebDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputWebDocument)
	x.String(e.Url)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Vector(e.Attributes)
	return x.buf
}

func (e InputWebFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputWebFileLocation)
	x.String(e.Url)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputWebFileGeoPointLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputWebFileGeoPointLocation)
	x.Bytes(e.GeoPoint.encode())
	x.Long(e.AccessHash)
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Zoom)
	x.Int(e.Scale)
	return x.buf
}

func (e UploadWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadWebFile)
	x.Int(e.Size)
	x.String(e.MimeType)
	x.Bytes(e.FileType.encode())
	x.Int(e.Mtime)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e PaymentsPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsPaymentForm)
	x.Int(e.Flags)
	//flag e.CanSaveCredentials
	//flag e.PasswordMissing
	x.Int(e.BotID)
	x.Bytes(e.Invoice.encode())
	x.Int(e.ProviderID)
	x.String(e.Url)
	if e.Flags&16 != 0 {
		x.String(e.NativeProvider)
	}
	if e.Flags&16 != 0 {
		x.Bytes(e.NativeParams.encode())
	}
	if e.Flags&1 != 0 {
		x.Bytes(e.SavedInfo.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.SavedCredentials.encode())
	}
	x.Vector(e.Users)
	return x.buf
}

func (e PaymentsValidatedRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsValidatedRequestedInfo)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.ID)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.ShippingOptions)
	}
	return x.buf
}

func (e PaymentsPaymentResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsPaymentResult)
	x.Bytes(e.Updates.encode())
	return x.buf
}

func (e PaymentsPaymentVerificationNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsPaymentVerificationNeeded)
	x.String(e.Url)
	return x.buf
}

func (e PaymentsPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsPaymentReceipt)
	x.Int(e.Flags)
	x.Int(e.Date)
	x.Int(e.BotID)
	x.Bytes(e.Invoice.encode())
	x.Int(e.ProviderID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Info.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Shipping.encode())
	}
	x.String(e.Currency)
	x.Long(e.TotalAmount)
	x.String(e.CredentialsTitle)
	x.Vector(e.Users)
	return x.buf
}

func (e PaymentsSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsSavedInfo)
	x.Int(e.Flags)
	//flag e.HasSavedCredentials
	if e.Flags&1 != 0 {
		x.Bytes(e.SavedInfo.encode())
	}
	return x.buf
}

func (e InputPaymentCredentialsSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPaymentCredentialsSaved)
	x.String(e.ID)
	x.StringBytes(e.TmpPassword)
	return x.buf
}

func (e InputPaymentCredentials) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPaymentCredentials)
	x.Int(e.Flags)
	//flag e.Save
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e InputPaymentCredentialsApplePay) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPaymentCredentialsApplePay)
	x.Bytes(e.PaymentData.encode())
	return x.buf
}

func (e InputPaymentCredentialsAndroidPay) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPaymentCredentialsAndroidPay)
	x.Bytes(e.PaymentToken.encode())
	x.String(e.GoogleTransactionID)
	return x.buf
}

func (e AccountTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountTmpPassword)
	x.StringBytes(e.TmpPassword)
	x.Int(e.ValidUntil)
	return x.buf
}

func (e ShippingOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcShippingOption)
	x.String(e.ID)
	x.String(e.Title)
	x.Vector(e.Prices)
	return x.buf
}

func (e InputStickerSetItem) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputStickerSetItem)
	x.Int(e.Flags)
	x.Bytes(e.Document.encode())
	x.String(e.Emoji)
	if e.Flags&1 != 0 {
		x.Bytes(e.MaskCoords.encode())
	}
	return x.buf
}

func (e InputPhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputPhoneCall)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e PhoneCallEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallEmpty)
	x.Long(e.ID)
	return x.buf
}

func (e PhoneCallWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallWaiting)
	x.Int(e.Flags)
	//flag e.Video
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.Bytes(e.Protocol.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReceiveDate)
	}
	return x.buf
}

func (e PhoneCallRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallRequested)
	x.Int(e.Flags)
	//flag e.Video
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GAHash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e PhoneCallAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallAccepted)
	x.Int(e.Flags)
	//flag e.Video
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GB)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e PhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCall)
	x.Int(e.Flags)
	//flag e.P2pAllowed
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Date)
	x.Int(e.AdminID)
	x.Int(e.ParticipantID)
	x.StringBytes(e.GAOrB)
	x.Long(e.KeyFingerprint)
	x.Bytes(e.Protocol.encode())
	x.Vector(e.Connections)
	x.Int(e.StartDate)
	return x.buf
}

func (e PhoneCallDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallDiscarded)
	x.Int(e.Flags)
	//flag e.NeedRating
	//flag e.NeedDebug
	//flag e.Video
	x.Long(e.ID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Reason.encode())
	}
	if e.Flags&2 != 0 {
		x.Int(e.Duration)
	}
	return x.buf
}

func (e PhoneConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneConnection)
	x.Long(e.ID)
	x.String(e.Ip)
	x.String(e.Ipv6)
	x.Int(e.Port)
	x.StringBytes(e.PeerTag)
	return x.buf
}

func (e PhoneCallProtocol) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneCallProtocol)
	x.Int(e.Flags)
	//flag e.UdpP2p
	//flag e.UdpReflector
	x.Int(e.MinLayer)
	x.Int(e.MaxLayer)
	x.VectorString(e.LibraryVersions)
	return x.buf
}

func (e PhonePhoneCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhonePhoneCall)
	x.Bytes(e.PhoneCall.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e UploadCdnFileReuploadNeeded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadCdnFileReuploadNeeded)
	x.StringBytes(e.RequestToken)
	return x.buf
}

func (e UploadCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadCdnFile)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e CdnPublicKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcCdnPublicKey)
	x.Int(e.DcID)
	x.String(e.PublicKey)
	return x.buf
}

func (e CdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcCdnConfig)
	x.Vector(e.PublicKeys)
	return x.buf
}

func (e LangPackString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangPackString)
	x.String(e.Key)
	x.String(e.Value)
	return x.buf
}

func (e LangPackStringPluralized) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangPackStringPluralized)
	x.Int(e.Flags)
	x.String(e.Key)
	if e.Flags&1 != 0 {
		x.String(e.ZeroValue)
	}
	if e.Flags&2 != 0 {
		x.String(e.OneValue)
	}
	if e.Flags&4 != 0 {
		x.String(e.TwoValue)
	}
	if e.Flags&8 != 0 {
		x.String(e.FewValue)
	}
	if e.Flags&16 != 0 {
		x.String(e.ManyValue)
	}
	x.String(e.OtherValue)
	return x.buf
}

func (e LangPackStringDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangPackStringDeleted)
	x.String(e.Key)
	return x.buf
}

func (e LangPackDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangPackDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	x.Int(e.Version)
	x.Vector(e.Strings)
	return x.buf
}

func (e LangPackLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangPackLanguage)
	x.Int(e.Flags)
	//flag e.Official
	//flag e.Rtl
	//flag e.Beta
	x.String(e.Name)
	x.String(e.NativeName)
	x.String(e.LangCode)
	if e.Flags&2 != 0 {
		x.String(e.BaseLangCode)
	}
	x.String(e.PluralCode)
	x.Int(e.StringsCount)
	x.Int(e.TranslatedCount)
	x.String(e.TranslationsUrl)
	return x.buf
}

func (e ChannelAdminLogEventActionChangeTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangeTitle)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}

func (e ChannelAdminLogEventActionChangeAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangeAbout)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}

func (e ChannelAdminLogEventActionChangeUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangeUsername)
	x.String(e.PrevValue)
	x.String(e.NewValue)
	return x.buf
}

func (e ChannelAdminLogEventActionChangePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangePhoto)
	x.Bytes(e.PrevPhoto.encode())
	x.Bytes(e.NewPhoto.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionToggleInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionToggleInvites)
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionToggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionToggleSignatures)
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionUpdatePinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionUpdatePinned)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionEditMessage)
	x.Bytes(e.PrevMessage.encode())
	x.Bytes(e.NewMessage.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionDeleteMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionDeleteMessage)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionParticipantJoin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionParticipantJoin)
	return x.buf
}

func (e ChannelAdminLogEventActionParticipantLeave) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionParticipantLeave)
	return x.buf
}

func (e ChannelAdminLogEventActionParticipantInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionParticipantInvite)
	x.Bytes(e.Participant.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionParticipantToggleBan) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionParticipantToggleBan)
	x.Bytes(e.PrevParticipant.encode())
	x.Bytes(e.NewParticipant.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionParticipantToggleAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionParticipantToggleAdmin)
	x.Bytes(e.PrevParticipant.encode())
	x.Bytes(e.NewParticipant.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionChangeStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangeStickerSet)
	x.Bytes(e.PrevStickerset.encode())
	x.Bytes(e.NewStickerset.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionTogglePreHistoryHidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionTogglePreHistoryHidden)
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionDefaultBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionDefaultBannedRights)
	x.Bytes(e.PrevBannedRights.encode())
	x.Bytes(e.NewBannedRights.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionStopPoll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionStopPoll)
	x.Bytes(e.Message.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionChangeLinkedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangeLinkedChat)
	x.Int(e.PrevValue)
	x.Int(e.NewValue)
	return x.buf
}

func (e ChannelAdminLogEventActionChangeLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionChangeLocation)
	x.Bytes(e.PrevValue.encode())
	x.Bytes(e.NewValue.encode())
	return x.buf
}

func (e ChannelAdminLogEventActionToggleSlowMode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventActionToggleSlowMode)
	x.Int(e.PrevValue)
	x.Int(e.NewValue)
	return x.buf
}

func (e ChannelAdminLogEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEvent)
	x.Long(e.ID)
	x.Int(e.Date)
	x.Int(e.UserID)
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e ChannelsAdminLogResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsAdminLogResults)
	x.Vector(e.Events)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e ChannelAdminLogEventsFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelAdminLogEventsFilter)
	x.Int(e.Flags)
	//flag e.Join
	//flag e.Leave
	//flag e.Invite
	//flag e.Ban
	//flag e.Unban
	//flag e.Kick
	//flag e.Unkick
	//flag e.Promote
	//flag e.Demote
	//flag e.Info
	//flag e.Settings
	//flag e.Pinned
	//flag e.Edit
	//flag e.Delete
	return x.buf
}

func (e PopularContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPopularContact)
	x.Long(e.ClientID)
	x.Int(e.Importers)
	return x.buf
}

func (e MessagesFavedStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFavedStickersNotModified)
	return x.buf
}

func (e MessagesFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFavedStickers)
	x.Int(e.Hash)
	x.Vector(e.Packs)
	x.Vector(e.Stickers)
	return x.buf
}

func (e RecentMeUrlUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRecentMeUrlUnknown)
	x.String(e.Url)
	return x.buf
}

func (e RecentMeUrlUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRecentMeUrlUser)
	x.String(e.Url)
	x.Int(e.UserID)
	return x.buf
}

func (e RecentMeUrlChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRecentMeUrlChat)
	x.String(e.Url)
	x.Int(e.ChatID)
	return x.buf
}

func (e RecentMeUrlChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRecentMeUrlChatInvite)
	x.String(e.Url)
	x.Bytes(e.ChatInvite.encode())
	return x.buf
}

func (e RecentMeUrlStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRecentMeUrlStickerSet)
	x.String(e.Url)
	x.Bytes(e.Set.encode())
	return x.buf
}

func (e HelpRecentMeUrls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpRecentMeUrls)
	x.Vector(e.Urls)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e InputSingleMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputSingleMedia)
	x.Int(e.Flags)
	x.Bytes(e.Media.encode())
	x.Long(e.RandomID)
	x.String(e.Message)
	if e.Flags&1 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e WebAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebAuthorization)
	x.Long(e.Hash)
	x.Int(e.BotID)
	x.String(e.Domain)
	x.String(e.Browser)
	x.String(e.Platform)
	x.Int(e.DateCreated)
	x.Int(e.DateActive)
	x.String(e.Ip)
	x.String(e.Region)
	return x.buf
}

func (e AccountWebAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountWebAuthorizations)
	x.Vector(e.Authorizations)
	x.Vector(e.Users)
	return x.buf
}

func (e InputMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessageID)
	x.Int(e.ID)
	return x.buf
}

func (e InputMessageReplyTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessageReplyTo)
	x.Int(e.ID)
	return x.buf
}

func (e InputMessagePinned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputMessagePinned)
	return x.buf
}

func (e InputDialogPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputDialogPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e InputDialogPeerFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputDialogPeerFolder)
	x.Int(e.FolderID)
	return x.buf
}

func (e DialogPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDialogPeer)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e DialogPeerFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDialogPeerFolder)
	x.Int(e.FolderID)
	return x.buf
}

func (e MessagesFoundStickerSetsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFoundStickerSetsNotModified)
	return x.buf
}

func (e MessagesFoundStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFoundStickerSets)
	x.Int(e.Hash)
	x.Vector(e.Sets)
	return x.buf
}

func (e FileHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFileHash)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.StringBytes(e.Hash)
	return x.buf
}

func (e InputClientProxy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputClientProxy)
	x.String(e.Address)
	x.Int(e.Port)
	return x.buf
}

func (e HelpTermsOfServiceUpdateEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpTermsOfServiceUpdateEmpty)
	x.Int(e.Expires)
	return x.buf
}

func (e HelpTermsOfServiceUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpTermsOfServiceUpdate)
	x.Int(e.Expires)
	x.Bytes(e.TermsOfService.encode())
	return x.buf
}

func (e InputSecureFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputSecureFileUploaded)
	x.Long(e.ID)
	x.Int(e.Parts)
	x.String(e.Md5Checksum)
	x.StringBytes(e.FileHash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e InputSecureFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputSecureFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e SecureFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureFileEmpty)
	return x.buf
}

func (e SecureFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureFile)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.Int(e.Size)
	x.Int(e.DcID)
	x.Int(e.Date)
	x.StringBytes(e.FileHash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e SecureData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureData)
	x.StringBytes(e.Data)
	x.StringBytes(e.DataHash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e SecurePlainPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecurePlainPhone)
	x.String(e.Phone)
	return x.buf
}

func (e SecurePlainEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecurePlainEmail)
	x.String(e.Email)
	return x.buf
}

func (e SecureValueTypePersonalDetails) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypePersonalDetails)
	return x.buf
}

func (e SecureValueTypePassport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypePassport)
	return x.buf
}

func (e SecureValueTypeDriverLicense) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeDriverLicense)
	return x.buf
}

func (e SecureValueTypeIdentityCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeIdentityCard)
	return x.buf
}

func (e SecureValueTypeInternalPassport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeInternalPassport)
	return x.buf
}

func (e SecureValueTypeAddress) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeAddress)
	return x.buf
}

func (e SecureValueTypeUtilityBill) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeUtilityBill)
	return x.buf
}

func (e SecureValueTypeBankStatement) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeBankStatement)
	return x.buf
}

func (e SecureValueTypeRentalAgreement) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeRentalAgreement)
	return x.buf
}

func (e SecureValueTypePassportRegistration) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypePassportRegistration)
	return x.buf
}

func (e SecureValueTypeTemporaryRegistration) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeTemporaryRegistration)
	return x.buf
}

func (e SecureValueTypePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypePhone)
	return x.buf
}

func (e SecureValueTypeEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueTypeEmail)
	return x.buf
}

func (e SecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValue)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Data.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.FrontSide.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReverseSide.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Selfie.encode())
	}
	if e.Flags&64 != 0 {
		x.Vector(e.Translation)
	}
	if e.Flags&16 != 0 {
		x.Vector(e.Files)
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.PlainData.encode())
	}
	x.StringBytes(e.Hash)
	return x.buf
}

func (e InputSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputSecureValue)
	x.Int(e.Flags)
	x.Bytes(e.Type.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Data.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.FrontSide.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReverseSide.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Selfie.encode())
	}
	if e.Flags&64 != 0 {
		x.Vector(e.Translation)
	}
	if e.Flags&16 != 0 {
		x.Vector(e.Files)
	}
	if e.Flags&32 != 0 {
		x.Bytes(e.PlainData.encode())
	}
	return x.buf
}

func (e SecureValueHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueHash)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.Hash)
	return x.buf
}

func (e SecureValueErrorData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorData)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.DataHash)
	x.String(e.Field)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorFrontSide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorFrontSide)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorReverseSide) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorReverseSide)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorSelfie) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorSelfie)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorFile)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorFiles) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorFiles)
	x.Bytes(e.Type.encode())
	x.Vector(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueError)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.Hash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorTranslationFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorTranslationFile)
	x.Bytes(e.Type.encode())
	x.StringBytes(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureValueErrorTranslationFiles) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureValueErrorTranslationFiles)
	x.Bytes(e.Type.encode())
	x.Vector(e.FileHash)
	x.String(e.Text)
	return x.buf
}

func (e SecureCredentialsEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureCredentialsEncrypted)
	x.StringBytes(e.Data)
	x.StringBytes(e.Hash)
	x.StringBytes(e.Secret)
	return x.buf
}

func (e AccountAuthorizationForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountAuthorizationForm)
	x.Int(e.Flags)
	x.Vector(e.RequiredTypes)
	x.Vector(e.Values)
	x.Vector(e.Errors)
	x.Vector(e.Users)
	if e.Flags&1 != 0 {
		x.String(e.PrivacyPolicyUrl)
	}
	return x.buf
}

func (e AccountSentEmailCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSentEmailCode)
	x.String(e.EmailPattern)
	x.Int(e.Length)
	return x.buf
}

func (e HelpDeepLinkInfoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpDeepLinkInfoEmpty)
	return x.buf
}

func (e HelpDeepLinkInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpDeepLinkInfo)
	x.Int(e.Flags)
	//flag e.UpdateApp
	x.String(e.Message)
	if e.Flags&2 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e SavedPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSavedPhoneContact)
	x.String(e.Phone)
	x.String(e.FirstName)
	x.String(e.LastName)
	x.Int(e.Date)
	return x.buf
}

func (e AccountTakeout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountTakeout)
	x.Long(e.ID)
	return x.buf
}

func (e PasswordKdfAlgoUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPasswordKdfAlgoUnknown)
	return x.buf
}

func (e PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow)
	x.StringBytes(e.Salt1)
	x.StringBytes(e.Salt2)
	x.Int(e.G)
	x.StringBytes(e.P)
	return x.buf
}

func (e SecurePasswordKdfAlgoUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecurePasswordKdfAlgoUnknown)
	return x.buf
}

func (e SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000)
	x.StringBytes(e.Salt)
	return x.buf
}

func (e SecurePasswordKdfAlgoSHA512) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecurePasswordKdfAlgoSHA512)
	x.StringBytes(e.Salt)
	return x.buf
}

func (e SecureSecretSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureSecretSettings)
	x.Bytes(e.SecureAlgo.encode())
	x.StringBytes(e.SecureSecret)
	x.Long(e.SecureSecretID)
	return x.buf
}

func (e InputCheckPasswordEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputCheckPasswordEmpty)
	return x.buf
}

func (e InputCheckPasswordSRP) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputCheckPasswordSRP)
	x.Long(e.SrpID)
	x.StringBytes(e.A)
	x.StringBytes(e.M1)
	return x.buf
}

func (e SecureRequiredType) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureRequiredType)
	x.Int(e.Flags)
	//flag e.NativeNames
	//flag e.SelfieRequired
	//flag e.TranslationRequired
	x.Bytes(e.Type.encode())
	return x.buf
}

func (e SecureRequiredTypeOneOf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcSecureRequiredTypeOneOf)
	x.Vector(e.Types)
	return x.buf
}

func (e HelpPassportConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpPassportConfigNotModified)
	return x.buf
}

func (e HelpPassportConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpPassportConfig)
	x.Int(e.Hash)
	x.Bytes(e.CountriesLangs.encode())
	return x.buf
}

func (e InputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputAppEvent)
	x.Double(e.Time)
	x.String(e.Type)
	x.Long(e.Peer)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e JsonObjectValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonObjectValue)
	x.String(e.Key)
	x.Bytes(e.Value.encode())
	return x.buf
}

func (e JsonNull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonNull)
	return x.buf
}

func (e JsonBool) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonBool)
	x.Bytes(e.Value.encode())
	return x.buf
}

func (e JsonNumber) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonNumber)
	x.Double(e.Value)
	return x.buf
}

func (e JsonString) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonString)
	x.String(e.Value)
	return x.buf
}

func (e JsonArray) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonArray)
	x.Vector(e.Value)
	return x.buf
}

func (e JsonObject) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcJsonObject)
	x.Vector(e.Value)
	return x.buf
}

func (e PageTableCell) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageTableCell)
	x.Int(e.Flags)
	//flag e.Header
	//flag e.AlignCenter
	//flag e.AlignRight
	//flag e.ValignMiddle
	//flag e.ValignBottom
	if e.Flags&128 != 0 {
		x.Bytes(e.Text.encode())
	}
	if e.Flags&2 != 0 {
		x.Int(e.Colspan)
	}
	if e.Flags&4 != 0 {
		x.Int(e.Rowspan)
	}
	return x.buf
}

func (e PageTableRow) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageTableRow)
	x.Vector(e.Cells)
	return x.buf
}

func (e PageCaption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageCaption)
	x.Bytes(e.Text.encode())
	x.Bytes(e.Credit.encode())
	return x.buf
}

func (e PageListItemText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageListItemText)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageListItemBlocks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageListItemBlocks)
	x.Vector(e.Blocks)
	return x.buf
}

func (e PageListOrderedItemText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageListOrderedItemText)
	x.String(e.Num)
	x.Bytes(e.Text.encode())
	return x.buf
}

func (e PageListOrderedItemBlocks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageListOrderedItemBlocks)
	x.String(e.Num)
	x.Vector(e.Blocks)
	return x.buf
}

func (e PageRelatedArticle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPageRelatedArticle)
	x.Int(e.Flags)
	x.String(e.Url)
	x.Long(e.WebpageID)
	if e.Flags&1 != 0 {
		x.String(e.Title)
	}
	if e.Flags&2 != 0 {
		x.String(e.Description)
	}
	if e.Flags&4 != 0 {
		x.Long(e.PhotoID)
	}
	if e.Flags&8 != 0 {
		x.String(e.Author)
	}
	if e.Flags&16 != 0 {
		x.Int(e.PublishedDate)
	}
	return x.buf
}

func (e Page) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPage)
	x.Int(e.Flags)
	//flag e.Part
	//flag e.Rtl
	//flag e.V2
	x.String(e.Url)
	x.Vector(e.Blocks)
	x.Vector(e.Photos)
	x.Vector(e.Documents)
	if e.Flags&8 != 0 {
		x.Int(e.Views)
	}
	return x.buf
}

func (e HelpSupportName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpSupportName)
	x.String(e.Name)
	return x.buf
}

func (e HelpUserInfoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpUserInfoEmpty)
	return x.buf
}

func (e HelpUserInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpUserInfo)
	x.String(e.Message)
	x.Vector(e.Entities)
	x.String(e.Author)
	x.Int(e.Date)
	return x.buf
}

func (e PollAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPollAnswer)
	x.String(e.Text)
	x.StringBytes(e.Option)
	return x.buf
}

func (e Poll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPoll)
	x.Long(e.ID)
	x.Int(e.Flags)
	//flag e.Closed
	//flag e.PublicVoters
	//flag e.MultipleChoice
	//flag e.Quiz
	x.String(e.Question)
	x.Vector(e.Answers)
	if e.Flags&16 != 0 {
		x.Int(e.ClosePeriod)
	}
	if e.Flags&32 != 0 {
		x.Int(e.CloseDate)
	}
	return x.buf
}

func (e PollAnswerVoters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPollAnswerVoters)
	x.Int(e.Flags)
	//flag e.Chosen
	//flag e.Correct
	x.StringBytes(e.Option)
	x.Int(e.Voters)
	return x.buf
}

func (e PollResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPollResults)
	x.Int(e.Flags)
	//flag e.Min
	if e.Flags&2 != 0 {
		x.Vector(e.Results)
	}
	if e.Flags&4 != 0 {
		x.Int(e.TotalVoters)
	}
	if e.Flags&8 != 0 {
		x.VectorInt(e.RecentVoters)
	}
	if e.Flags&16 != 0 {
		x.String(e.Solution)
	}
	if e.Flags&16 != 0 {
		x.Vector(e.SolutionEntities)
	}
	return x.buf
}

func (e ChatOnlines) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatOnlines)
	x.Int(e.Onlines)
	return x.buf
}

func (e StatsURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsURL)
	x.String(e.Url)
	return x.buf
}

func (e ChatAdminRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatAdminRights)
	x.Int(e.Flags)
	//flag e.ChangeInfo
	//flag e.PostMessages
	//flag e.EditMessages
	//flag e.DeleteMessages
	//flag e.BanUsers
	//flag e.InviteUsers
	//flag e.PinMessages
	//flag e.AddAdmins
	return x.buf
}

func (e ChatBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChatBannedRights)
	x.Int(e.Flags)
	//flag e.ViewMessages
	//flag e.SendMessages
	//flag e.SendMedia
	//flag e.SendStickers
	//flag e.SendGifs
	//flag e.SendGames
	//flag e.SendInline
	//flag e.EmbedLinks
	//flag e.SendPolls
	//flag e.ChangeInfo
	//flag e.InviteUsers
	//flag e.PinMessages
	x.Int(e.UntilDate)
	return x.buf
}

func (e InputWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputWallPaper)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputWallPaperSlug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputWallPaperSlug)
	x.String(e.Slug)
	return x.buf
}

func (e InputWallPaperNoFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputWallPaperNoFile)
	return x.buf
}

func (e AccountWallPapersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountWallPapersNotModified)
	return x.buf
}

func (e AccountWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountWallPapers)
	x.Int(e.Hash)
	x.Vector(e.Wallpapers)
	return x.buf
}

func (e CodeSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcCodeSettings)
	x.Int(e.Flags)
	//flag e.AllowFlashcall
	//flag e.CurrentNumber
	//flag e.AllowAppHash
	return x.buf
}

func (e WallPaperSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWallPaperSettings)
	x.Int(e.Flags)
	//flag e.Blur
	//flag e.Motion
	if e.Flags&1 != 0 {
		x.Int(e.BackgroundColor)
	}
	if e.Flags&16 != 0 {
		x.Int(e.SecondBackgroundColor)
	}
	if e.Flags&8 != 0 {
		x.Int(e.Intensity)
	}
	if e.Flags&16 != 0 {
		x.Int(e.Rotation)
	}
	return x.buf
}

func (e AutoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAutoDownloadSettings)
	x.Int(e.Flags)
	//flag e.Disabled
	//flag e.VideoPreloadLarge
	//flag e.AudioPreloadNext
	//flag e.PhonecallsLessData
	x.Int(e.PhotoSizeMax)
	x.Int(e.VideoSizeMax)
	x.Int(e.FileSizeMax)
	x.Int(e.VideoUploadMaxbitrate)
	return x.buf
}

func (e AccountAutoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountAutoDownloadSettings)
	x.Bytes(e.Low.encode())
	x.Bytes(e.Medium.encode())
	x.Bytes(e.High.encode())
	return x.buf
}

func (e EmojiKeyword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEmojiKeyword)
	x.String(e.Keyword)
	x.VectorString(e.Emoticons)
	return x.buf
}

func (e EmojiKeywordDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEmojiKeywordDeleted)
	x.String(e.Keyword)
	x.VectorString(e.Emoticons)
	return x.buf
}

func (e EmojiKeywordsDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEmojiKeywordsDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	x.Int(e.Version)
	x.Vector(e.Keywords)
	return x.buf
}

func (e EmojiURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEmojiURL)
	x.String(e.Url)
	return x.buf
}

func (e EmojiLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcEmojiLanguage)
	x.String(e.LangCode)
	return x.buf
}

func (e FileLocationToBeDeprecated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFileLocationToBeDeprecated)
	x.Long(e.VolumeID)
	x.Int(e.LocalID)
	return x.buf
}

func (e Folder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFolder)
	x.Int(e.Flags)
	//flag e.AutofillNewBroadcasts
	//flag e.AutofillPublicGroups
	//flag e.AutofillNewCorrespondents
	x.Int(e.ID)
	x.String(e.Title)
	if e.Flags&8 != 0 {
		x.Bytes(e.Photo.encode())
	}
	return x.buf
}

func (e InputFolderPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputFolderPeer)
	x.Bytes(e.Peer.encode())
	x.Int(e.FolderID)
	return x.buf
}

func (e FolderPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFolderPeer)
	x.Bytes(e.Peer.encode())
	x.Int(e.FolderID)
	return x.buf
}

func (e MessagesSearchCounter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSearchCounter)
	x.Int(e.Flags)
	//flag e.Inexact
	x.Bytes(e.Filter.encode())
	x.Int(e.Count)
	return x.buf
}

func (e UrlAuthResultRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUrlAuthResultRequest)
	x.Int(e.Flags)
	//flag e.RequestWriteAccess
	x.Bytes(e.Bot.encode())
	x.String(e.Domain)
	return x.buf
}

func (e UrlAuthResultAccepted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUrlAuthResultAccepted)
	x.String(e.Url)
	return x.buf
}

func (e UrlAuthResultDefault) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUrlAuthResultDefault)
	return x.buf
}

func (e ChannelLocationEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelLocationEmpty)
	return x.buf
}

func (e ChannelLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelLocation)
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Address)
	return x.buf
}

func (e PeerLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerLocated)
	x.Bytes(e.Peer.encode())
	x.Int(e.Expires)
	x.Int(e.Distance)
	return x.buf
}

func (e PeerSelfLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPeerSelfLocated)
	x.Int(e.Expires)
	return x.buf
}

func (e RestrictionReason) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcRestrictionReason)
	x.String(e.Platform)
	x.String(e.Reason)
	x.String(e.Text)
	return x.buf
}

func (e InputTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputTheme)
	x.Long(e.ID)
	x.Long(e.AccessHash)
	return x.buf
}

func (e InputThemeSlug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputThemeSlug)
	x.String(e.Slug)
	return x.buf
}

func (e Theme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcTheme)
	x.Int(e.Flags)
	//flag e.Creator
	//flag e.Default
	x.Long(e.ID)
	x.Long(e.AccessHash)
	x.String(e.Slug)
	x.String(e.Title)
	if e.Flags&4 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Settings.encode())
	}
	x.Int(e.InstallsCount)
	return x.buf
}

func (e AccountThemesNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountThemesNotModified)
	return x.buf
}

func (e AccountThemes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountThemes)
	x.Int(e.Hash)
	x.Vector(e.Themes)
	return x.buf
}

func (e AuthLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthLoginToken)
	x.Int(e.Expires)
	x.StringBytes(e.Token)
	return x.buf
}

func (e AuthLoginTokenMigrateTo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthLoginTokenMigrateTo)
	x.Int(e.DcID)
	x.StringBytes(e.Token)
	return x.buf
}

func (e AuthLoginTokenSuccess) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthLoginTokenSuccess)
	x.Bytes(e.Authorization.encode())
	return x.buf
}

func (e AccountContentSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountContentSettings)
	x.Int(e.Flags)
	//flag e.SensitiveEnabled
	//flag e.SensitiveCanChange
	return x.buf
}

func (e MessagesInactiveChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesInactiveChats)
	x.VectorInt(e.Dates)
	x.Vector(e.Chats)
	x.Vector(e.Users)
	return x.buf
}

func (e BaseThemeClassic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBaseThemeClassic)
	return x.buf
}

func (e BaseThemeDay) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBaseThemeDay)
	return x.buf
}

func (e BaseThemeNight) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBaseThemeNight)
	return x.buf
}

func (e BaseThemeTinted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBaseThemeTinted)
	return x.buf
}

func (e BaseThemeArctic) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBaseThemeArctic)
	return x.buf
}

func (e InputThemeSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInputThemeSettings)
	x.Int(e.Flags)
	x.Bytes(e.BaseTheme.encode())
	x.Int(e.AccentColor)
	if e.Flags&1 != 0 {
		x.Int(e.MessageTopColor)
	}
	if e.Flags&1 != 0 {
		x.Int(e.MessageBottomColor)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Wallpaper.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.WallpaperSettings.encode())
	}
	return x.buf
}

func (e ThemeSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcThemeSettings)
	x.Int(e.Flags)
	x.Bytes(e.BaseTheme.encode())
	x.Int(e.AccentColor)
	if e.Flags&1 != 0 {
		x.Int(e.MessageTopColor)
	}
	if e.Flags&1 != 0 {
		x.Int(e.MessageBottomColor)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Wallpaper.encode())
	}
	return x.buf
}

func (e WebPageAttributeTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcWebPageAttributeTheme)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Vector(e.Documents)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e MessageUserVote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageUserVote)
	x.Int(e.UserID)
	x.StringBytes(e.Option)
	x.Int(e.Date)
	return x.buf
}

func (e MessageUserVoteInputOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageUserVoteInputOption)
	x.Int(e.UserID)
	x.Int(e.Date)
	return x.buf
}

func (e MessageUserVoteMultiple) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageUserVoteMultiple)
	x.Int(e.UserID)
	x.Vector(e.Options)
	x.Int(e.Date)
	return x.buf
}

func (e MessagesVotesList) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesVotesList)
	x.Int(e.Flags)
	x.Int(e.Count)
	x.Vector(e.Votes)
	x.Vector(e.Users)
	if e.Flags&1 != 0 {
		x.String(e.NextOffset)
	}
	return x.buf
}

func (e BankCardOpenUrl) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBankCardOpenUrl)
	x.String(e.Url)
	x.String(e.Name)
	return x.buf
}

func (e PaymentsBankCardData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsBankCardData)
	x.String(e.Title)
	x.Vector(e.OpenUrls)
	return x.buf
}

func (e DialogFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDialogFilter)
	x.Int(e.Flags)
	//flag e.Contacts
	//flag e.NonContacts
	//flag e.Groups
	//flag e.Broadcasts
	//flag e.Bots
	//flag e.ExcludeMuted
	//flag e.ExcludeRead
	//flag e.ExcludeArchived
	x.Int(e.ID)
	x.String(e.Title)
	if e.Flags&33554432 != 0 {
		x.String(e.Emoticon)
	}
	x.Vector(e.PinnedPeers)
	x.Vector(e.IncludePeers)
	x.Vector(e.ExcludePeers)
	return x.buf
}

func (e DialogFilterSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcDialogFilterSuggested)
	x.Bytes(e.Filter.encode())
	x.String(e.Description)
	return x.buf
}

func (e StatsDateRangeDays) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsDateRangeDays)
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	return x.buf
}

func (e StatsAbsValueAndPrev) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsAbsValueAndPrev)
	x.Double(e.Current)
	x.Double(e.Previous)
	return x.buf
}

func (e StatsPercentValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsPercentValue)
	x.Double(e.Part)
	x.Double(e.Total)
	return x.buf
}

func (e StatsGraphAsync) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGraphAsync)
	x.String(e.Token)
	return x.buf
}

func (e StatsGraphError) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGraphError)
	x.String(e.Error)
	return x.buf
}

func (e StatsGraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGraph)
	x.Int(e.Flags)
	x.Bytes(e.Json.encode())
	if e.Flags&1 != 0 {
		x.String(e.ZoomToken)
	}
	return x.buf
}

func (e MessageInteractionCounters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessageInteractionCounters)
	x.Int(e.MsgID)
	x.Int(e.Views)
	x.Int(e.Forwards)
	return x.buf
}

func (e StatsBroadcastStats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsBroadcastStats)
	x.Bytes(e.Period.encode())
	x.Bytes(e.Followers.encode())
	x.Bytes(e.ViewsPerPost.encode())
	x.Bytes(e.SharesPerPost.encode())
	x.Bytes(e.EnabledNotifications.encode())
	x.Bytes(e.GrowthGraph.encode())
	x.Bytes(e.FollowersGraph.encode())
	x.Bytes(e.MuteGraph.encode())
	x.Bytes(e.TopHoursGraph.encode())
	x.Bytes(e.InteractionsGraph.encode())
	x.Bytes(e.IvInteractionsGraph.encode())
	x.Bytes(e.ViewsBySourceGraph.encode())
	x.Bytes(e.NewFollowersBySourceGraph.encode())
	x.Bytes(e.LanguagesGraph.encode())
	x.Vector(e.RecentMessageInteractions)
	return x.buf
}

func (e HelpPromoDataEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpPromoDataEmpty)
	x.Int(e.Expires)
	return x.buf
}

func (e HelpPromoData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpPromoData)
	x.Int(e.Flags)
	//flag e.Proxy
	x.Int(e.Expires)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Chats)
	x.Vector(e.Users)
	if e.Flags&2 != 0 {
		x.String(e.PsaType)
	}
	if e.Flags&4 != 0 {
		x.String(e.PsaMessage)
	}
	return x.buf
}

func (e VideoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcVideoSize)
	x.Int(e.Flags)
	x.String(e.Type)
	x.Bytes(e.Location.encode())
	x.Int(e.W)
	x.Int(e.H)
	x.Int(e.Size)
	if e.Flags&1 != 0 {
		x.Double(e.VideoStartTs)
	}
	return x.buf
}

func (e StatsGroupTopPoster) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGroupTopPoster)
	x.Int(e.UserID)
	x.Int(e.Messages)
	x.Int(e.AvgChars)
	return x.buf
}

func (e StatsGroupTopAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGroupTopAdmin)
	x.Int(e.UserID)
	x.Int(e.Deleted)
	x.Int(e.Kicked)
	x.Int(e.Banned)
	return x.buf
}

func (e StatsGroupTopInviter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGroupTopInviter)
	x.Int(e.UserID)
	x.Int(e.Invitations)
	return x.buf
}

func (e StatsMegagroupStats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsMegagroupStats)
	x.Bytes(e.Period.encode())
	x.Bytes(e.Members.encode())
	x.Bytes(e.Messages.encode())
	x.Bytes(e.Viewers.encode())
	x.Bytes(e.Posters.encode())
	x.Bytes(e.GrowthGraph.encode())
	x.Bytes(e.MembersGraph.encode())
	x.Bytes(e.NewMembersBySourceGraph.encode())
	x.Bytes(e.LanguagesGraph.encode())
	x.Bytes(e.MessagesGraph.encode())
	x.Bytes(e.ActionsGraph.encode())
	x.Bytes(e.TopHoursGraph.encode())
	x.Bytes(e.WeekdaysGraph.encode())
	x.Vector(e.TopPosters)
	x.Vector(e.TopAdmins)
	x.Vector(e.TopInviters)
	x.Vector(e.Users)
	return x.buf
}

func (e GlobalPrivacySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcGlobalPrivacySettings)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.ArchiveAndMuteNewNoncontactPeers.encode())
	}
	return x.buf
}

func (e InvokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvokeAfterMsg)
	x.Long(e.MsgID)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e InvokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvokeAfterMsgs)
	x.VectorLong(e.MsgIds)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e InitConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInitConnection)
	x.Int(e.Flags)
	x.Int(e.ApiID)
	x.String(e.DeviceModel)
	x.String(e.SystemVersion)
	x.String(e.AppVersion)
	x.String(e.SystemLangCode)
	x.String(e.LangPack)
	x.String(e.LangCode)
	if e.Flags&1 != 0 {
		x.Bytes(e.Proxy.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Params.encode())
	}
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e InvokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvokeWithLayer)
	x.Int(e.Layer)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e InvokeWithoutUpdates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvokeWithoutUpdates)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e InvokeWithMessagesRange) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvokeWithMessagesRange)
	x.Bytes(e.Range.encode())
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e InvokeWithTakeout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcInvokeWithTakeout)
	x.Long(e.TakeoutID)
	x.Bytes(e.Query.encode())
	return x.buf
}

func (e AuthSendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSendCode)
	x.String(e.PhoneNumber)
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AuthSignUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSignUp)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.FirstName)
	x.String(e.LastName)
	return x.buf
}

func (e AuthSignIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthSignIn)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e AuthLogOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthLogOut)
	return x.buf
}

func (e AuthResetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthResetAuthorizations)
	return x.buf
}

func (e AuthExportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthExportAuthorization)
	x.Int(e.DcID)
	return x.buf
}

func (e AuthImportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthImportAuthorization)
	x.Int(e.ID)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e AuthBindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthBindTempAuthKey)
	x.Long(e.PermAuthKeyID)
	x.Long(e.Nonce)
	x.Int(e.ExpiresAt)
	x.StringBytes(e.EncryptedMessage)
	return x.buf
}

func (e AuthImportBotAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthImportBotAuthorization)
	x.Int(e.Flags)
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	x.String(e.BotAuthToken)
	return x.buf
}

func (e AuthCheckPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthCheckPassword)
	x.Bytes(e.Password.encode())
	return x.buf
}

func (e AuthRequestPasswordRecovery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthRequestPasswordRecovery)
	return x.buf
}

func (e AuthRecoverPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthRecoverPassword)
	x.String(e.Code)
	return x.buf
}

func (e AuthResendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthResendCode)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}

func (e AuthCancelCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthCancelCode)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	return x.buf
}

func (e AuthDropTempAuthKeys) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthDropTempAuthKeys)
	x.VectorLong(e.ExceptAuthKeys)
	return x.buf
}

func (e AuthExportLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthExportLoginToken)
	x.Int(e.ApiID)
	x.String(e.ApiHash)
	x.VectorInt(e.ExceptIds)
	return x.buf
}

func (e AuthImportLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthImportLoginToken)
	x.StringBytes(e.Token)
	return x.buf
}

func (e AuthAcceptLoginToken) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAuthAcceptLoginToken)
	x.StringBytes(e.Token)
	return x.buf
}

func (e AccountRegisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountRegisterDevice)
	x.Int(e.Flags)
	//flag e.NoMuted
	x.Int(e.TokenType)
	x.String(e.Token)
	x.Bytes(e.AppSandbox.encode())
	x.StringBytes(e.Secret)
	x.VectorInt(e.OtherUids)
	return x.buf
}

func (e AccountUnregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUnregisterDevice)
	x.Int(e.TokenType)
	x.String(e.Token)
	x.VectorInt(e.OtherUids)
	return x.buf
}

func (e AccountUpdateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdateNotifySettings)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountGetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetNotifySettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e AccountResetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountResetNotifySettings)
	return x.buf
}

func (e AccountUpdateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdateProfile)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.String(e.FirstName)
	}
	if e.Flags&2 != 0 {
		x.String(e.LastName)
	}
	if e.Flags&4 != 0 {
		x.String(e.About)
	}
	return x.buf
}

func (e AccountUpdateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdateStatus)
	x.Bytes(e.Offline.encode())
	return x.buf
}

func (e AccountGetWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetWallPapers)
	x.Int(e.Hash)
	return x.buf
}

func (e AccountReportPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountReportPeer)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Reason.encode())
	return x.buf
}

func (e AccountCheckUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountCheckUsername)
	x.String(e.Username)
	return x.buf
}

func (e AccountUpdateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdateUsername)
	x.String(e.Username)
	return x.buf
}

func (e AccountGetPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetPrivacy)
	x.Bytes(e.Key.encode())
	return x.buf
}

func (e AccountSetPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSetPrivacy)
	x.Bytes(e.Key.encode())
	x.Vector(e.Rules)
	return x.buf
}

func (e AccountDeleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountDeleteAccount)
	x.String(e.Reason)
	return x.buf
}

func (e AccountGetAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetAccountTTL)
	return x.buf
}

func (e AccountSetAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSetAccountTTL)
	x.Bytes(e.Ttl.encode())
	return x.buf
}

func (e AccountSendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSendChangePhoneCode)
	x.String(e.PhoneNumber)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountChangePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountChangePhone)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e AccountUpdateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdateDeviceLocked)
	x.Int(e.Period)
	return x.buf
}

func (e AccountGetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetAuthorizations)
	return x.buf
}

func (e AccountResetAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountResetAuthorization)
	x.Long(e.Hash)
	return x.buf
}

func (e AccountGetPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetPassword)
	return x.buf
}

func (e AccountGetPasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetPasswordSettings)
	x.Bytes(e.Password.encode())
	return x.buf
}

func (e AccountUpdatePasswordSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdatePasswordSettings)
	x.Bytes(e.Password.encode())
	x.Bytes(e.NewSettings.encode())
	return x.buf
}

func (e AccountSendConfirmPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSendConfirmPhoneCode)
	x.String(e.Hash)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountConfirmPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountConfirmPhone)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e AccountGetTmpPassword) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetTmpPassword)
	x.Bytes(e.Password.encode())
	x.Int(e.Period)
	return x.buf
}

func (e AccountGetWebAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetWebAuthorizations)
	return x.buf
}

func (e AccountResetWebAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountResetWebAuthorization)
	x.Long(e.Hash)
	return x.buf
}

func (e AccountResetWebAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountResetWebAuthorizations)
	return x.buf
}

func (e AccountGetAllSecureValues) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetAllSecureValues)
	return x.buf
}

func (e AccountGetSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetSecureValue)
	x.Vector(e.Types)
	return x.buf
}

func (e AccountSaveSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSaveSecureValue)
	x.Bytes(e.Value.encode())
	x.Long(e.SecureSecretID)
	return x.buf
}

func (e AccountDeleteSecureValue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountDeleteSecureValue)
	x.Vector(e.Types)
	return x.buf
}

func (e AccountGetAuthorizationForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetAuthorizationForm)
	x.Int(e.BotID)
	x.String(e.Scope)
	x.String(e.PublicKey)
	return x.buf
}

func (e AccountAcceptAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountAcceptAuthorization)
	x.Int(e.BotID)
	x.String(e.Scope)
	x.String(e.PublicKey)
	x.Vector(e.ValueHashes)
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e AccountSendVerifyPhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSendVerifyPhoneCode)
	x.String(e.PhoneNumber)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountVerifyPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountVerifyPhone)
	x.String(e.PhoneNumber)
	x.String(e.PhoneCodeHash)
	x.String(e.PhoneCode)
	return x.buf
}

func (e AccountSendVerifyEmailCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSendVerifyEmailCode)
	x.String(e.Email)
	return x.buf
}

func (e AccountVerifyEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountVerifyEmail)
	x.String(e.Email)
	x.String(e.Code)
	return x.buf
}

func (e AccountInitTakeoutSession) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountInitTakeoutSession)
	x.Int(e.Flags)
	//flag e.Contacts
	//flag e.MessageUsers
	//flag e.MessageChats
	//flag e.MessageMegagroups
	//flag e.MessageChannels
	//flag e.Files
	if e.Flags&32 != 0 {
		x.Int(e.FileMaxSize)
	}
	return x.buf
}

func (e AccountFinishTakeoutSession) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountFinishTakeoutSession)
	x.Int(e.Flags)
	//flag e.Success
	return x.buf
}

func (e AccountConfirmPasswordEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountConfirmPasswordEmail)
	x.String(e.Code)
	return x.buf
}

func (e AccountResendPasswordEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountResendPasswordEmail)
	return x.buf
}

func (e AccountCancelPasswordEmail) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountCancelPasswordEmail)
	return x.buf
}

func (e AccountGetContactSignUpNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetContactSignUpNotification)
	return x.buf
}

func (e AccountSetContactSignUpNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSetContactSignUpNotification)
	x.Bytes(e.Silent.encode())
	return x.buf
}

func (e AccountGetNotifyExceptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetNotifyExceptions)
	x.Int(e.Flags)
	//flag e.CompareSound
	if e.Flags&1 != 0 {
		x.Bytes(e.Peer.encode())
	}
	return x.buf
}

func (e AccountGetWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetWallPaper)
	x.Bytes(e.Wallpaper.encode())
	return x.buf
}

func (e AccountUploadWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUploadWallPaper)
	x.Bytes(e.File.encode())
	x.String(e.MimeType)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountSaveWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSaveWallPaper)
	x.Bytes(e.Wallpaper.encode())
	x.Bytes(e.Unsave.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountInstallWallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountInstallWallPaper)
	x.Bytes(e.Wallpaper.encode())
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountResetWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountResetWallPapers)
	return x.buf
}

func (e AccountGetAutoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetAutoDownloadSettings)
	return x.buf
}

func (e AccountSaveAutoDownloadSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSaveAutoDownloadSettings)
	x.Int(e.Flags)
	//flag e.Low
	//flag e.High
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e AccountUploadTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUploadTheme)
	x.Int(e.Flags)
	x.Bytes(e.File.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	x.String(e.FileName)
	x.String(e.MimeType)
	return x.buf
}

func (e AccountCreateTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountCreateTheme)
	x.Int(e.Flags)
	x.String(e.Slug)
	x.String(e.Title)
	if e.Flags&4 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e AccountUpdateTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountUpdateTheme)
	x.Int(e.Flags)
	x.String(e.Format)
	x.Bytes(e.Theme.encode())
	if e.Flags&1 != 0 {
		x.String(e.Slug)
	}
	if e.Flags&2 != 0 {
		x.String(e.Title)
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.Document.encode())
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.Settings.encode())
	}
	return x.buf
}

func (e AccountSaveTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSaveTheme)
	x.Bytes(e.Theme.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e AccountInstallTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountInstallTheme)
	x.Int(e.Flags)
	//flag e.Dark
	if e.Flags&2 != 0 {
		x.String(e.Format)
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Theme.encode())
	}
	return x.buf
}

func (e AccountGetTheme) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetTheme)
	x.String(e.Format)
	x.Bytes(e.Theme.encode())
	x.Long(e.DocumentID)
	return x.buf
}

func (e AccountGetThemes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetThemes)
	x.String(e.Format)
	x.Int(e.Hash)
	return x.buf
}

func (e AccountSetContentSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSetContentSettings)
	x.Int(e.Flags)
	//flag e.SensitiveEnabled
	return x.buf
}

func (e AccountGetContentSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetContentSettings)
	return x.buf
}

func (e AccountGetMultiWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetMultiWallPapers)
	x.Vector(e.Wallpapers)
	return x.buf
}

func (e AccountGetGlobalPrivacySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountGetGlobalPrivacySettings)
	return x.buf
}

func (e AccountSetGlobalPrivacySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcAccountSetGlobalPrivacySettings)
	x.Bytes(e.Settings.encode())
	return x.buf
}

func (e UsersGetUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUsersGetUsers)
	x.Vector(e.ID)
	return x.buf
}

func (e UsersGetFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUsersGetFullUser)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e UsersSetSecureValueErrors) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUsersSetSecureValueErrors)
	x.Bytes(e.ID.encode())
	x.Vector(e.Errors)
	return x.buf
}

func (e ContactsGetContactIDs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetContactIDs)
	x.Int(e.Hash)
	return x.buf
}

func (e ContactsGetStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetStatuses)
	return x.buf
}

func (e ContactsGetContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetContacts)
	x.Int(e.Hash)
	return x.buf
}

func (e ContactsImportContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsImportContacts)
	x.Vector(e.Contacts)
	return x.buf
}

func (e ContactsDeleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsDeleteContacts)
	x.Vector(e.ID)
	return x.buf
}

func (e ContactsDeleteByPhones) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsDeleteByPhones)
	x.VectorString(e.Phones)
	return x.buf
}

func (e ContactsBlock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsBlock)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e ContactsUnblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsUnblock)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e ContactsGetBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetBlocked)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e ContactsSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsSearch)
	x.String(e.Q)
	x.Int(e.Limit)
	return x.buf
}

func (e ContactsResolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsResolveUsername)
	x.String(e.Username)
	return x.buf
}

func (e ContactsGetTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetTopPeers)
	x.Int(e.Flags)
	//flag e.Correspondents
	//flag e.BotsPm
	//flag e.BotsInline
	//flag e.PhoneCalls
	//flag e.ForwardUsers
	//flag e.ForwardChats
	//flag e.Groups
	//flag e.Channels
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e ContactsResetTopPeerRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsResetTopPeerRating)
	x.Bytes(e.Category.encode())
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e ContactsResetSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsResetSaved)
	return x.buf
}

func (e ContactsGetSaved) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetSaved)
	return x.buf
}

func (e ContactsToggleTopPeers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsToggleTopPeers)
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e ContactsAddContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsAddContact)
	x.Int(e.Flags)
	//flag e.AddPhonePrivacyException
	x.Bytes(e.ID.encode())
	x.String(e.FirstName)
	x.String(e.LastName)
	x.String(e.Phone)
	return x.buf
}

func (e ContactsAcceptContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsAcceptContact)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e ContactsGetLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcContactsGetLocated)
	x.Int(e.Flags)
	//flag e.Background
	x.Bytes(e.GeoPoint.encode())
	if e.Flags&1 != 0 {
		x.Int(e.SelfExpires)
	}
	return x.buf
}

func (e MessagesGetMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetMessages)
	x.Vector(e.ID)
	return x.buf
}

func (e MessagesGetDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetDialogs)
	x.Int(e.Flags)
	//flag e.ExcludePinned
	if e.Flags&2 != 0 {
		x.Int(e.FolderID)
	}
	x.Int(e.OffsetDate)
	x.Int(e.OffsetID)
	x.Bytes(e.OffsetPeer.encode())
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesGetHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.OffsetID)
	x.Int(e.OffsetDate)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxID)
	x.Int(e.MinID)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesSearch) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSearch)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.String(e.Q)
	if e.Flags&1 != 0 {
		x.Bytes(e.FromID.encode())
	}
	x.Bytes(e.Filter.encode())
	x.Int(e.MinDate)
	x.Int(e.MaxDate)
	x.Int(e.OffsetID)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxID)
	x.Int(e.MinID)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesReadHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReadHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e MessagesDeleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDeleteHistory)
	x.Int(e.Flags)
	//flag e.JustClear
	//flag e.Revoke
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e MessagesDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDeleteMessages)
	x.Int(e.Flags)
	//flag e.Revoke
	x.VectorInt(e.ID)
	return x.buf
}

func (e MessagesReceivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReceivedMessages)
	x.Int(e.MaxID)
	return x.buf
}

func (e MessagesSetTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Action.encode())
	return x.buf
}

func (e MessagesSendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendMessage)
	x.Int(e.Flags)
	//flag e.NoWebpage
	//flag e.Silent
	//flag e.Background
	//flag e.ClearDraft
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.String(e.Message)
	x.Long(e.RandomID)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e MessagesSendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendMedia)
	x.Int(e.Flags)
	//flag e.Silent
	//flag e.Background
	//flag e.ClearDraft
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Bytes(e.Media.encode())
	x.String(e.Message)
	x.Long(e.RandomID)
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e MessagesForwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesForwardMessages)
	x.Int(e.Flags)
	//flag e.Silent
	//flag e.Background
	//flag e.WithMyScore
	//flag e.Grouped
	x.Bytes(e.FromPeer.encode())
	x.VectorInt(e.ID)
	x.VectorLong(e.RandomID)
	x.Bytes(e.ToPeer.encode())
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e MessagesReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReportSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesGetPeerSettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetPeerSettings)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesReport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReport)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	x.Bytes(e.Reason.encode())
	return x.buf
}

func (e MessagesGetChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetChats)
	x.VectorInt(e.ID)
	return x.buf
}

func (e MessagesGetFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetFullChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e MessagesEditChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditChatTitle)
	x.Int(e.ChatID)
	x.String(e.Title)
	return x.buf
}

func (e MessagesEditChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditChatPhoto)
	x.Int(e.ChatID)
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e MessagesAddChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAddChatUser)
	x.Int(e.ChatID)
	x.Bytes(e.UserID.encode())
	x.Int(e.FwdLimit)
	return x.buf
}

func (e MessagesDeleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDeleteChatUser)
	x.Int(e.ChatID)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e MessagesCreateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesCreateChat)
	x.Vector(e.Users)
	x.String(e.Title)
	return x.buf
}

func (e MessagesGetDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetDhConfig)
	x.Int(e.Version)
	x.Int(e.RandomLength)
	return x.buf
}

func (e MessagesRequestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesRequestEncryption)
	x.Bytes(e.UserID.encode())
	x.Int(e.RandomID)
	x.StringBytes(e.GA)
	return x.buf
}

func (e MessagesAcceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAcceptEncryption)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Long(e.KeyFingerprint)
	return x.buf
}

func (e MessagesDiscardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDiscardEncryption)
	x.Int(e.ChatID)
	return x.buf
}

func (e MessagesSetEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetEncryptedTyping)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Typing.encode())
	return x.buf
}

func (e MessagesReadEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReadEncryptedHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.MaxDate)
	return x.buf
}

func (e MessagesSendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendEncrypted)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.StringBytes(e.Data)
	return x.buf
}

func (e MessagesSendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.StringBytes(e.Data)
	x.Bytes(e.File.encode())
	return x.buf
}

func (e MessagesSendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendEncryptedService)
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.StringBytes(e.Data)
	return x.buf
}

func (e MessagesReceivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReceivedQueue)
	x.Int(e.MaxQts)
	return x.buf
}

func (e MessagesReportEncryptedSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReportEncryptedSpam)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesReadMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReadMessageContents)
	x.VectorInt(e.ID)
	return x.buf
}

func (e MessagesGetStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetStickers)
	x.String(e.Emoticon)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesGetAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetAllStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesGetWebPagePreview) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetWebPagePreview)
	x.Int(e.Flags)
	x.String(e.Message)
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e MessagesExportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesExportChatInvite)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesCheckChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesCheckChatInvite)
	x.String(e.Hash)
	return x.buf
}

func (e MessagesImportChatInvite) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesImportChatInvite)
	x.String(e.Hash)
	return x.buf
}

func (e MessagesGetStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e MessagesInstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesInstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Archived.encode())
	return x.buf
}

func (e MessagesUninstallStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesUninstallStickerSet)
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e MessagesStartBot) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesStartBot)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	x.Long(e.RandomID)
	x.String(e.StartParam)
	return x.buf
}

func (e MessagesGetMessagesViews) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetMessagesViews)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	x.Bytes(e.Increment.encode())
	return x.buf
}

func (e MessagesEditChatAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditChatAdmin)
	x.Int(e.ChatID)
	x.Bytes(e.UserID.encode())
	x.Bytes(e.IsAdmin.encode())
	return x.buf
}

func (e MessagesMigrateChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesMigrateChat)
	x.Int(e.ChatID)
	return x.buf
}

func (e MessagesSearchGlobal) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSearchGlobal)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Int(e.FolderID)
	}
	x.String(e.Q)
	x.Int(e.OffsetRate)
	x.Bytes(e.OffsetPeer.encode())
	x.Int(e.OffsetID)
	x.Int(e.Limit)
	return x.buf
}

func (e MessagesReorderStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReorderStickerSets)
	x.Int(e.Flags)
	//flag e.Masks
	x.VectorLong(e.Order)
	return x.buf
}

func (e MessagesGetDocumentByHash) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetDocumentByHash)
	x.StringBytes(e.Sha256)
	x.Int(e.Size)
	x.String(e.MimeType)
	return x.buf
}

func (e MessagesGetSavedGifs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetSavedGifs)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesSaveGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSaveGif)
	x.Bytes(e.ID.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e MessagesGetInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetInlineBotResults)
	x.Int(e.Flags)
	x.Bytes(e.Bot.encode())
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Bytes(e.GeoPoint.encode())
	}
	x.String(e.Query)
	x.String(e.Offset)
	return x.buf
}

func (e MessagesSetInlineBotResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetInlineBotResults)
	x.Int(e.Flags)
	//flag e.Gallery
	//flag e.Private
	x.Long(e.QueryID)
	x.Vector(e.Results)
	x.Int(e.CacheTime)
	if e.Flags&4 != 0 {
		x.String(e.NextOffset)
	}
	if e.Flags&8 != 0 {
		x.Bytes(e.SwitchPm.encode())
	}
	return x.buf
}

func (e MessagesSendInlineBotResult) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendInlineBotResult)
	x.Int(e.Flags)
	//flag e.Silent
	//flag e.Background
	//flag e.ClearDraft
	//flag e.HideVia
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Long(e.RandomID)
	x.Long(e.QueryID)
	x.String(e.ID)
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e MessagesGetMessageEditData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetMessageEditData)
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	return x.buf
}

func (e MessagesEditMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditMessage)
	x.Int(e.Flags)
	//flag e.NoWebpage
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	if e.Flags&2048 != 0 {
		x.String(e.Message)
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	if e.Flags&32768 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e MessagesEditInlineBotMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditInlineBotMessage)
	x.Int(e.Flags)
	//flag e.NoWebpage
	x.Bytes(e.ID.encode())
	if e.Flags&2048 != 0 {
		x.String(e.Message)
	}
	if e.Flags&16384 != 0 {
		x.Bytes(e.Media.encode())
	}
	if e.Flags&4 != 0 {
		x.Bytes(e.ReplyMarkup.encode())
	}
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e MessagesGetBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetBotCallbackAnswer)
	x.Int(e.Flags)
	//flag e.Game
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Data)
	}
	return x.buf
}

func (e MessagesSetBotCallbackAnswer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetBotCallbackAnswer)
	x.Int(e.Flags)
	//flag e.Alert
	x.Long(e.QueryID)
	if e.Flags&1 != 0 {
		x.String(e.Message)
	}
	if e.Flags&4 != 0 {
		x.String(e.Url)
	}
	x.Int(e.CacheTime)
	return x.buf
}

func (e MessagesGetPeerDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetPeerDialogs)
	x.Vector(e.Peers)
	return x.buf
}

func (e MessagesSaveDraft) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSaveDraft)
	x.Int(e.Flags)
	//flag e.NoWebpage
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Bytes(e.Peer.encode())
	x.String(e.Message)
	if e.Flags&8 != 0 {
		x.Vector(e.Entities)
	}
	return x.buf
}

func (e MessagesGetAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetAllDrafts)
	return x.buf
}

func (e MessagesGetFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetFeaturedStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesReadFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReadFeaturedStickers)
	x.VectorLong(e.ID)
	return x.buf
}

func (e MessagesGetRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetRecentStickers)
	x.Int(e.Flags)
	//flag e.Attached
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesSaveRecentSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSaveRecentSticker)
	x.Int(e.Flags)
	//flag e.Attached
	x.Bytes(e.ID.encode())
	x.Bytes(e.Unsave.encode())
	return x.buf
}

func (e MessagesClearRecentStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesClearRecentStickers)
	x.Int(e.Flags)
	//flag e.Attached
	return x.buf
}

func (e MessagesGetArchivedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetArchivedStickers)
	x.Int(e.Flags)
	//flag e.Masks
	x.Long(e.OffsetID)
	x.Int(e.Limit)
	return x.buf
}

func (e MessagesGetMaskStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetMaskStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesGetAttachedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetAttachedStickers)
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e MessagesSetGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetGameScore)
	x.Int(e.Flags)
	//flag e.EditMessage
	//flag e.Force
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	x.Bytes(e.UserID.encode())
	x.Int(e.Score)
	return x.buf
}

func (e MessagesSetInlineGameScore) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetInlineGameScore)
	x.Int(e.Flags)
	//flag e.EditMessage
	//flag e.Force
	x.Bytes(e.ID.encode())
	x.Bytes(e.UserID.encode())
	x.Int(e.Score)
	return x.buf
}

func (e MessagesGetGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetGameHighScores)
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e MessagesGetInlineGameHighScores) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetInlineGameHighScores)
	x.Bytes(e.ID.encode())
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e MessagesGetCommonChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetCommonChats)
	x.Bytes(e.UserID.encode())
	x.Int(e.MaxID)
	x.Int(e.Limit)
	return x.buf
}

func (e MessagesGetAllChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetAllChats)
	x.VectorInt(e.ExceptIds)
	return x.buf
}

func (e MessagesGetWebPage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetWebPage)
	x.String(e.Url)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesToggleDialogPin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesToggleDialogPin)
	x.Int(e.Flags)
	//flag e.Pinned
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesReorderPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReorderPinnedDialogs)
	x.Int(e.Flags)
	//flag e.Force
	x.Int(e.FolderID)
	x.Vector(e.Order)
	return x.buf
}

func (e MessagesGetPinnedDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetPinnedDialogs)
	x.Int(e.FolderID)
	return x.buf
}

func (e MessagesSetBotShippingResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetBotShippingResults)
	x.Int(e.Flags)
	x.Long(e.QueryID)
	if e.Flags&1 != 0 {
		x.String(e.Error)
	}
	if e.Flags&2 != 0 {
		x.Vector(e.ShippingOptions)
	}
	return x.buf
}

func (e MessagesSetBotPrecheckoutResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSetBotPrecheckoutResults)
	x.Int(e.Flags)
	//flag e.Success
	x.Long(e.QueryID)
	if e.Flags&1 != 0 {
		x.String(e.Error)
	}
	return x.buf
}

func (e MessagesUploadMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesUploadMedia)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Media.encode())
	return x.buf
}

func (e MessagesSendScreenshotNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendScreenshotNotification)
	x.Bytes(e.Peer.encode())
	x.Int(e.ReplyToMsgID)
	x.Long(e.RandomID)
	return x.buf
}

func (e MessagesGetFavedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetFavedStickers)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesFaveSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesFaveSticker)
	x.Bytes(e.ID.encode())
	x.Bytes(e.Unfave.encode())
	return x.buf
}

func (e MessagesGetUnreadMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetUnreadMentions)
	x.Bytes(e.Peer.encode())
	x.Int(e.OffsetID)
	x.Int(e.AddOffset)
	x.Int(e.Limit)
	x.Int(e.MaxID)
	x.Int(e.MinID)
	return x.buf
}

func (e MessagesReadMentions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesReadMentions)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesGetRecentLocations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetRecentLocations)
	x.Bytes(e.Peer.encode())
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesSendMultiMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendMultiMedia)
	x.Int(e.Flags)
	//flag e.Silent
	//flag e.Background
	//flag e.ClearDraft
	x.Bytes(e.Peer.encode())
	if e.Flags&1 != 0 {
		x.Int(e.ReplyToMsgID)
	}
	x.Vector(e.MultiMedia)
	if e.Flags&1024 != 0 {
		x.Int(e.ScheduleDate)
	}
	return x.buf
}

func (e MessagesUploadEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesUploadEncryptedFile)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.File.encode())
	return x.buf
}

func (e MessagesSearchStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSearchStickerSets)
	x.Int(e.Flags)
	//flag e.ExcludeFeatured
	x.String(e.Q)
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesGetSplitRanges) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetSplitRanges)
	return x.buf
}

func (e MessagesMarkDialogUnread) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesMarkDialogUnread)
	x.Int(e.Flags)
	//flag e.Unread
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesGetDialogUnreadMarks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetDialogUnreadMarks)
	return x.buf
}

func (e MessagesClearAllDrafts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesClearAllDrafts)
	return x.buf
}

func (e MessagesUpdatePinnedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesUpdatePinnedMessage)
	x.Int(e.Flags)
	//flag e.Silent
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	return x.buf
}

func (e MessagesSendVote) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendVote)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Vector(e.Options)
	return x.buf
}

func (e MessagesGetPollResults) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetPollResults)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	return x.buf
}

func (e MessagesGetOnlines) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetOnlines)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesGetStatsURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetStatsURL)
	x.Int(e.Flags)
	//flag e.Dark
	x.Bytes(e.Peer.encode())
	x.String(e.Params)
	return x.buf
}

func (e MessagesEditChatAbout) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditChatAbout)
	x.Bytes(e.Peer.encode())
	x.String(e.About)
	return x.buf
}

func (e MessagesEditChatDefaultBannedRights) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesEditChatDefaultBannedRights)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.BannedRights.encode())
	return x.buf
}

func (e MessagesGetEmojiKeywords) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetEmojiKeywords)
	x.String(e.LangCode)
	return x.buf
}

func (e MessagesGetEmojiKeywordsDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetEmojiKeywordsDifference)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	return x.buf
}

func (e MessagesGetEmojiKeywordsLanguages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetEmojiKeywordsLanguages)
	x.VectorString(e.LangCodes)
	return x.buf
}

func (e MessagesGetEmojiURL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetEmojiURL)
	x.String(e.LangCode)
	return x.buf
}

func (e MessagesGetSearchCounters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetSearchCounters)
	x.Bytes(e.Peer.encode())
	x.Vector(e.Filters)
	return x.buf
}

func (e MessagesRequestUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesRequestUrlAuth)
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ButtonID)
	return x.buf
}

func (e MessagesAcceptUrlAuth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesAcceptUrlAuth)
	x.Int(e.Flags)
	//flag e.WriteAllowed
	x.Bytes(e.Peer.encode())
	x.Int(e.MsgID)
	x.Int(e.ButtonID)
	return x.buf
}

func (e MessagesHidePeerSettingsBar) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesHidePeerSettingsBar)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e MessagesGetScheduledHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetScheduledHistory)
	x.Bytes(e.Peer.encode())
	x.Int(e.Hash)
	return x.buf
}

func (e MessagesGetScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e MessagesSendScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesSendScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e MessagesDeleteScheduledMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesDeleteScheduledMessages)
	x.Bytes(e.Peer.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e MessagesGetPollVotes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetPollVotes)
	x.Int(e.Flags)
	x.Bytes(e.Peer.encode())
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.StringBytes(e.Option)
	}
	if e.Flags&2 != 0 {
		x.String(e.Offset)
	}
	x.Int(e.Limit)
	return x.buf
}

func (e MessagesToggleStickerSets) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesToggleStickerSets)
	x.Int(e.Flags)
	//flag e.Uninstall
	//flag e.Archive
	//flag e.Unarchive
	x.Vector(e.Stickersets)
	return x.buf
}

func (e MessagesGetDialogFilters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetDialogFilters)
	return x.buf
}

func (e MessagesGetSuggestedDialogFilters) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetSuggestedDialogFilters)
	return x.buf
}

func (e MessagesUpdateDialogFilter) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesUpdateDialogFilter)
	x.Int(e.Flags)
	x.Int(e.ID)
	if e.Flags&1 != 0 {
		x.Bytes(e.Filter.encode())
	}
	return x.buf
}

func (e MessagesUpdateDialogFiltersOrder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesUpdateDialogFiltersOrder)
	x.VectorInt(e.Order)
	return x.buf
}

func (e MessagesGetOldFeaturedStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcMessagesGetOldFeaturedStickers)
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e UpdatesGetState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesGetState)
	return x.buf
}

func (e UpdatesGetDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesGetDifference)
	x.Int(e.Flags)
	x.Int(e.Pts)
	if e.Flags&1 != 0 {
		x.Int(e.PtsTotalLimit)
	}
	x.Int(e.Date)
	x.Int(e.Qts)
	return x.buf
}

func (e UpdatesGetChannelDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUpdatesGetChannelDifference)
	x.Int(e.Flags)
	//flag e.Force
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Pts)
	x.Int(e.Limit)
	return x.buf
}

func (e PhotosUpdateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosUpdateProfilePhoto)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e PhotosUploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosUploadProfilePhoto)
	x.Int(e.Flags)
	if e.Flags&1 != 0 {
		x.Bytes(e.File.encode())
	}
	if e.Flags&2 != 0 {
		x.Bytes(e.Video.encode())
	}
	if e.Flags&4 != 0 {
		x.Double(e.VideoStartTs)
	}
	return x.buf
}

func (e PhotosDeletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosDeletePhotos)
	x.Vector(e.ID)
	return x.buf
}

func (e PhotosGetUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhotosGetUserPhotos)
	x.Bytes(e.UserID.encode())
	x.Int(e.Offset)
	x.Long(e.MaxID)
	x.Int(e.Limit)
	return x.buf
}

func (e UploadSaveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadSaveFilePart)
	x.Long(e.FileID)
	x.Int(e.FilePart)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e UploadGetFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadGetFile)
	x.Int(e.Flags)
	//flag e.Precise
	//flag e.CdnSupported
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e UploadSaveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadSaveBigFilePart)
	x.Long(e.FileID)
	x.Int(e.FilePart)
	x.Int(e.FileTotalParts)
	x.StringBytes(e.Bytes)
	return x.buf
}

func (e UploadGetWebFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadGetWebFile)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e UploadGetCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadGetCdnFile)
	x.StringBytes(e.FileToken)
	x.Int(e.Offset)
	x.Int(e.Limit)
	return x.buf
}

func (e UploadReuploadCdnFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadReuploadCdnFile)
	x.StringBytes(e.FileToken)
	x.StringBytes(e.RequestToken)
	return x.buf
}

func (e UploadGetCdnFileHashes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadGetCdnFileHashes)
	x.StringBytes(e.FileToken)
	x.Int(e.Offset)
	return x.buf
}

func (e UploadGetFileHashes) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcUploadGetFileHashes)
	x.Bytes(e.Location.encode())
	x.Int(e.Offset)
	return x.buf
}

func (e HelpGetConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetConfig)
	return x.buf
}

func (e HelpGetNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetNearestDc)
	return x.buf
}

func (e HelpGetAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetAppUpdate)
	x.String(e.Source)
	return x.buf
}

func (e HelpGetInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetInviteText)
	return x.buf
}

func (e HelpGetSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetSupport)
	return x.buf
}

func (e HelpGetAppChangelog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetAppChangelog)
	x.String(e.PrevAppVersion)
	return x.buf
}

func (e HelpSetBotUpdatesStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpSetBotUpdatesStatus)
	x.Int(e.PendingUpdatesCount)
	x.String(e.Message)
	return x.buf
}

func (e HelpGetCdnConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetCdnConfig)
	return x.buf
}

func (e HelpGetRecentMeUrls) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetRecentMeUrls)
	x.String(e.Referer)
	return x.buf
}

func (e HelpGetTermsOfServiceUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetTermsOfServiceUpdate)
	return x.buf
}

func (e HelpAcceptTermsOfService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpAcceptTermsOfService)
	x.Bytes(e.ID.encode())
	return x.buf
}

func (e HelpGetDeepLinkInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetDeepLinkInfo)
	x.String(e.Path)
	return x.buf
}

func (e HelpGetAppConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetAppConfig)
	return x.buf
}

func (e HelpSaveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpSaveAppLog)
	x.Vector(e.Events)
	return x.buf
}

func (e HelpGetPassportConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetPassportConfig)
	x.Int(e.Hash)
	return x.buf
}

func (e HelpGetSupportName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetSupportName)
	return x.buf
}

func (e HelpGetUserInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetUserInfo)
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e HelpEditUserInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpEditUserInfo)
	x.Bytes(e.UserID.encode())
	x.String(e.Message)
	x.Vector(e.Entities)
	return x.buf
}

func (e HelpGetPromoData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpGetPromoData)
	return x.buf
}

func (e HelpHidePromoData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpHidePromoData)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e HelpDismissSuggestion) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcHelpDismissSuggestion)
	x.String(e.Suggestion)
	return x.buf
}

func (e ChannelsReadHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsReadHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e ChannelsDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsDeleteMessages)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e ChannelsDeleteUserHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsDeleteUserHistory)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e ChannelsReportSpam) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsReportSpam)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e ChannelsGetMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetMessages)
	x.Bytes(e.Channel.encode())
	x.Vector(e.ID)
	return x.buf
}

func (e ChannelsGetParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetParticipants)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Filter.encode())
	x.Int(e.Offset)
	x.Int(e.Limit)
	x.Int(e.Hash)
	return x.buf
}

func (e ChannelsGetParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetParticipant)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	return x.buf
}

func (e ChannelsGetChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetChannels)
	x.Vector(e.ID)
	return x.buf
}

func (e ChannelsGetFullChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetFullChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e ChannelsCreateChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsCreateChannel)
	x.Int(e.Flags)
	//flag e.Broadcast
	//flag e.Megagroup
	x.String(e.Title)
	x.String(e.About)
	if e.Flags&4 != 0 {
		x.Bytes(e.GeoPoint.encode())
	}
	if e.Flags&4 != 0 {
		x.String(e.Address)
	}
	return x.buf
}

func (e ChannelsEditAdmin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsEditAdmin)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.Bytes(e.AdminRights.encode())
	x.String(e.Rank)
	return x.buf
}

func (e ChannelsEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsEditTitle)
	x.Bytes(e.Channel.encode())
	x.String(e.Title)
	return x.buf
}

func (e ChannelsEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsEditPhoto)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Photo.encode())
	return x.buf
}

func (e ChannelsCheckUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsCheckUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

func (e ChannelsUpdateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsUpdateUsername)
	x.Bytes(e.Channel.encode())
	x.String(e.Username)
	return x.buf
}

func (e ChannelsJoinChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsJoinChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e ChannelsLeaveChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsLeaveChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e ChannelsInviteToChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsInviteToChannel)
	x.Bytes(e.Channel.encode())
	x.Vector(e.Users)
	return x.buf
}

func (e ChannelsDeleteChannel) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsDeleteChannel)
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e ChannelsExportMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsExportMessageLink)
	x.Bytes(e.Channel.encode())
	x.Int(e.ID)
	x.Bytes(e.Grouped.encode())
	return x.buf
}

func (e ChannelsToggleSignatures) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsToggleSignatures)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e ChannelsGetAdminedPublicChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetAdminedPublicChannels)
	x.Int(e.Flags)
	//flag e.ByLocation
	//flag e.CheckLimit
	return x.buf
}

func (e ChannelsEditBanned) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsEditBanned)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.Bytes(e.BannedRights.encode())
	return x.buf
}

func (e ChannelsGetAdminLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetAdminLog)
	x.Int(e.Flags)
	x.Bytes(e.Channel.encode())
	x.String(e.Q)
	if e.Flags&1 != 0 {
		x.Bytes(e.EventsFilter.encode())
	}
	if e.Flags&2 != 0 {
		x.Vector(e.Admins)
	}
	x.Long(e.MaxID)
	x.Long(e.MinID)
	x.Int(e.Limit)
	return x.buf
}

func (e ChannelsSetStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsSetStickers)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Stickerset.encode())
	return x.buf
}

func (e ChannelsReadMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsReadMessageContents)
	x.Bytes(e.Channel.encode())
	x.VectorInt(e.ID)
	return x.buf
}

func (e ChannelsDeleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsDeleteHistory)
	x.Bytes(e.Channel.encode())
	x.Int(e.MaxID)
	return x.buf
}

func (e ChannelsTogglePreHistoryHidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsTogglePreHistoryHidden)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.Enabled.encode())
	return x.buf
}

func (e ChannelsGetLeftChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetLeftChannels)
	x.Int(e.Offset)
	return x.buf
}

func (e ChannelsGetGroupsForDiscussion) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetGroupsForDiscussion)
	return x.buf
}

func (e ChannelsSetDiscussionGroup) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsSetDiscussionGroup)
	x.Bytes(e.Broadcast.encode())
	x.Bytes(e.Group.encode())
	return x.buf
}

func (e ChannelsEditCreator) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsEditCreator)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.UserID.encode())
	x.Bytes(e.Password.encode())
	return x.buf
}

func (e ChannelsEditLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsEditLocation)
	x.Bytes(e.Channel.encode())
	x.Bytes(e.GeoPoint.encode())
	x.String(e.Address)
	return x.buf
}

func (e ChannelsToggleSlowMode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsToggleSlowMode)
	x.Bytes(e.Channel.encode())
	x.Int(e.Seconds)
	return x.buf
}

func (e ChannelsGetInactiveChannels) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcChannelsGetInactiveChannels)
	return x.buf
}

func (e BotsSendCustomRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotsSendCustomRequest)
	x.String(e.CustomMethod)
	x.Bytes(e.Params.encode())
	return x.buf
}

func (e BotsAnswerWebhookJSONQuery) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotsAnswerWebhookJSONQuery)
	x.Long(e.QueryID)
	x.Bytes(e.Data.encode())
	return x.buf
}

func (e BotsSetBotCommands) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcBotsSetBotCommands)
	x.Vector(e.Commands)
	return x.buf
}

func (e PaymentsGetPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsGetPaymentForm)
	x.Int(e.MsgID)
	return x.buf
}

func (e PaymentsGetPaymentReceipt) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsGetPaymentReceipt)
	x.Int(e.MsgID)
	return x.buf
}

func (e PaymentsValidateRequestedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsValidateRequestedInfo)
	x.Int(e.Flags)
	//flag e.Save
	x.Int(e.MsgID)
	x.Bytes(e.Info.encode())
	return x.buf
}

func (e PaymentsSendPaymentForm) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsSendPaymentForm)
	x.Int(e.Flags)
	x.Int(e.MsgID)
	if e.Flags&1 != 0 {
		x.String(e.RequestedInfoID)
	}
	if e.Flags&2 != 0 {
		x.String(e.ShippingOptionID)
	}
	x.Bytes(e.Credentials.encode())
	return x.buf
}

func (e PaymentsGetSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsGetSavedInfo)
	return x.buf
}

func (e PaymentsClearSavedInfo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsClearSavedInfo)
	x.Int(e.Flags)
	//flag e.Credentials
	//flag e.Info
	return x.buf
}

func (e PaymentsGetBankCardData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPaymentsGetBankCardData)
	x.String(e.Number)
	return x.buf
}

func (e StickersCreateStickerSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickersCreateStickerSet)
	x.Int(e.Flags)
	//flag e.Masks
	//flag e.Animated
	x.Bytes(e.UserID.encode())
	x.String(e.Title)
	x.String(e.ShortName)
	if e.Flags&4 != 0 {
		x.Bytes(e.Thumb.encode())
	}
	x.Vector(e.Stickers)
	return x.buf
}

func (e StickersRemoveStickerFromSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickersRemoveStickerFromSet)
	x.Bytes(e.Sticker.encode())
	return x.buf
}

func (e StickersChangeStickerPosition) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickersChangeStickerPosition)
	x.Bytes(e.Sticker.encode())
	x.Int(e.Position)
	return x.buf
}

func (e StickersAddStickerToSet) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickersAddStickerToSet)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Sticker.encode())
	return x.buf
}

func (e StickersSetStickerSetThumb) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStickersSetStickerSetThumb)
	x.Bytes(e.Stickerset.encode())
	x.Bytes(e.Thumb.encode())
	return x.buf
}

func (e PhoneGetCallConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneGetCallConfig)
	return x.buf
}

func (e PhoneRequestCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneRequestCall)
	x.Int(e.Flags)
	//flag e.Video
	x.Bytes(e.UserID.encode())
	x.Int(e.RandomID)
	x.StringBytes(e.GAHash)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e PhoneAcceptCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneAcceptCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GB)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e PhoneConfirmCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneConfirmCall)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.GA)
	x.Long(e.KeyFingerprint)
	x.Bytes(e.Protocol.encode())
	return x.buf
}

func (e PhoneReceivedCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneReceivedCall)
	x.Bytes(e.Peer.encode())
	return x.buf
}

func (e PhoneDiscardCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneDiscardCall)
	x.Int(e.Flags)
	//flag e.Video
	x.Bytes(e.Peer.encode())
	x.Int(e.Duration)
	x.Bytes(e.Reason.encode())
	x.Long(e.ConnectionID)
	return x.buf
}

func (e PhoneSetCallRating) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneSetCallRating)
	x.Int(e.Flags)
	//flag e.UserInitiative
	x.Bytes(e.Peer.encode())
	x.Int(e.Rating)
	x.String(e.Comment)
	return x.buf
}

func (e PhoneSaveCallDebug) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneSaveCallDebug)
	x.Bytes(e.Peer.encode())
	x.Bytes(e.Debug.encode())
	return x.buf
}

func (e PhoneSendSignalingData) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcPhoneSendSignalingData)
	x.Bytes(e.Peer.encode())
	x.StringBytes(e.Data)
	return x.buf
}

func (e LangpackGetLangPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangpackGetLangPack)
	x.String(e.LangPack)
	x.String(e.LangCode)
	return x.buf
}

func (e LangpackGetStrings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangpackGetStrings)
	x.String(e.LangPack)
	x.String(e.LangCode)
	x.VectorString(e.Keys)
	return x.buf
}

func (e LangpackGetDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangpackGetDifference)
	x.String(e.LangPack)
	x.String(e.LangCode)
	x.Int(e.FromVersion)
	return x.buf
}

func (e LangpackGetLanguages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangpackGetLanguages)
	x.String(e.LangPack)
	return x.buf
}

func (e LangpackGetLanguage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcLangpackGetLanguage)
	x.String(e.LangPack)
	x.String(e.LangCode)
	return x.buf
}

func (e FoldersEditPeerFolders) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFoldersEditPeerFolders)
	x.Vector(e.FolderPeers)
	return x.buf
}

func (e FoldersDeleteFolder) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcFoldersDeleteFolder)
	x.Int(e.FolderID)
	return x.buf
}

func (e StatsGetBroadcastStats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGetBroadcastStats)
	x.Int(e.Flags)
	//flag e.Dark
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e StatsLoadAsyncGraph) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsLoadAsyncGraph)
	x.Int(e.Flags)
	x.String(e.Token)
	if e.Flags&1 != 0 {
		x.Long(e.X)
	}
	return x.buf
}

func (e StatsGetMegagroupStats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crcStatsGetMegagroupStats)
	x.Int(e.Flags)
	//flag e.Dark
	x.Bytes(e.Channel.encode())
	return x.buf
}

func (e ReqPqMulti) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ReqDHParams) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e SetClientDHParams) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e RpcDropAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e GetFutureSalts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e Ping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PingDelayDisconnect) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e DestroySession) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HttpWait) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e DestroyAuthKey) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e True) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e BoolFalse) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e BoolTrue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e Error) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e IpPort) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InvokeAfterMsg) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InvokeAfterMsgs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InitConnection) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InvokeWithLayer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InvokeWithoutUpdates) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InvokeWithMessagesRange) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e InvokeWithTakeout) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthSendCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthSignUp) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthSignIn) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthLogOut) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthResetAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthExportAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthImportAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthBindTempAuthKey) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthImportBotAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthCheckPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthRequestPasswordRecovery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthRecoverPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthResendCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthCancelCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthDropTempAuthKeys) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthExportLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthImportLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AuthAcceptLoginToken) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountRegisterDevice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUnregisterDevice) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdateNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountResetNotifySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdateProfile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdateStatus) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetWallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountReportPeer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountCheckUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdateUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetPrivacy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSetPrivacy) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountDeleteAccount) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetAccountTTL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSetAccountTTL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSendChangePhoneCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountChangePhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdateDeviceLocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountResetAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetPasswordSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdatePasswordSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSendConfirmPhoneCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountConfirmPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetTmpPassword) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetWebAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountResetWebAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountResetWebAuthorizations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetAllSecureValues) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e AccountGetSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e AccountSaveSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountDeleteSecureValue) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetAuthorizationForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountAcceptAuthorization) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSendVerifyPhoneCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountVerifyPhone) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSendVerifyEmailCode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountVerifyEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountInitTakeoutSession) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountFinishTakeoutSession) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountConfirmPasswordEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountResendPasswordEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountCancelPasswordEmail) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetContactSignUpNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSetContactSignUpNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetNotifyExceptions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUploadWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSaveWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountInstallWallPaper) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountResetWallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetAutoDownloadSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSaveAutoDownloadSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUploadTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountCreateTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountUpdateTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSaveTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountInstallTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetTheme) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetThemes) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSetContentSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetContentSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountGetMultiWallPapers) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e AccountGetGlobalPrivacySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e AccountSetGlobalPrivacySettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UsersGetUsers) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e UsersGetFullUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UsersSetSecureValueErrors) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsGetContactIDs) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorInt(dbuf.VectorInt())
}

func (e ContactsGetStatuses) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e ContactsGetContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsImportContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsDeleteContacts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsDeleteByPhones) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsBlock) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsUnblock) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsGetBlocked) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsSearch) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsResolveUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsGetTopPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsResetTopPeerRating) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsResetSaved) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsGetSaved) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e ContactsToggleTopPeers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsAddContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsAcceptContact) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ContactsGetLocated) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSearch) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReadHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesDeleteHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesDeleteMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReceivedMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesSetTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesForwardMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReportSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetPeerSettings) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReport) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetFullChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesEditChatTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesEditChatPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesAddChatUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesDeleteChatUser) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesCreateChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetDhConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesRequestEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesAcceptEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesDiscardEncryption) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSetEncryptedTyping) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReadEncryptedHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendEncrypted) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendEncryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendEncryptedService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReceivedQueue) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorLong(dbuf.VectorLong())
}

func (e MessagesReportEncryptedSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReadMessageContents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetAllStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetWebPagePreview) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesExportChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesCheckChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesImportChatInvite) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesInstallStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesUninstallStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesStartBot) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetMessagesViews) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorInt(dbuf.VectorInt())
}

func (e MessagesEditChatAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesMigrateChat) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSearchGlobal) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReorderStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetDocumentByHash) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetSavedGifs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSaveGif) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetInlineBotResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSetInlineBotResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendInlineBotResult) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetMessageEditData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesEditMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesEditInlineBotMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetBotCallbackAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSetBotCallbackAnswer) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetPeerDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSaveDraft) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetAllDrafts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReadFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetRecentStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSaveRecentSticker) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesClearRecentStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetArchivedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetMaskStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetAttachedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesSetGameScore) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSetInlineGameScore) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetGameHighScores) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetInlineGameHighScores) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetCommonChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetAllChats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetWebPage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesToggleDialogPin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReorderPinnedDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetPinnedDialogs) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSetBotShippingResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSetBotPrecheckoutResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesUploadMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendScreenshotNotification) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetFavedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesFaveSticker) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetUnreadMentions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesReadMentions) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetRecentLocations) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendMultiMedia) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesUploadEncryptedFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSearchStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetSplitRanges) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesMarkDialogUnread) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetDialogUnreadMarks) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesClearAllDrafts) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesUpdatePinnedMessage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendVote) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetPollResults) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetOnlines) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetStatsURL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesEditChatAbout) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesEditChatDefaultBannedRights) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetEmojiKeywords) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetEmojiKeywordsDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetEmojiKeywordsLanguages) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesGetEmojiURL) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetSearchCounters) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesRequestUrlAuth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesAcceptUrlAuth) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesHidePeerSettingsBar) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetScheduledHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesSendScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesDeleteScheduledMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetPollVotes) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesToggleStickerSets) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetDialogFilters) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesGetSuggestedDialogFilters) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e MessagesUpdateDialogFilter) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesUpdateDialogFiltersOrder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e MessagesGetOldFeaturedStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UpdatesGetState) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UpdatesGetDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UpdatesGetChannelDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhotosUpdateProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhotosUploadProfilePhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhotosDeletePhotos) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorLong(dbuf.VectorLong())
}

func (e PhotosGetUserPhotos) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UploadSaveFilePart) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UploadGetFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UploadSaveBigFilePart) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UploadGetWebFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UploadGetCdnFile) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e UploadReuploadCdnFile) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e UploadGetCdnFileHashes) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e UploadGetFileHashes) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e HelpGetConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetNearestDc) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetAppUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetInviteText) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetSupport) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetAppChangelog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpSetBotUpdatesStatus) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetCdnConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetRecentMeUrls) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetTermsOfServiceUpdate) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpAcceptTermsOfService) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetDeepLinkInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetAppConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpSaveAppLog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetPassportConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetSupportName) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetUserInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpEditUserInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpGetPromoData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpHidePromoData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e HelpDismissSuggestion) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsReadHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsDeleteMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsDeleteUserHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsReportSpam) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetMessages) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetParticipants) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetParticipant) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetFullChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsCreateChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsEditAdmin) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsEditTitle) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsEditPhoto) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsCheckUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsUpdateUsername) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsJoinChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsLeaveChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsInviteToChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsDeleteChannel) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsExportMessageLink) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsToggleSignatures) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetAdminedPublicChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsEditBanned) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetAdminLog) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsSetStickers) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsReadMessageContents) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsDeleteHistory) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsTogglePreHistoryHidden) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetLeftChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetGroupsForDiscussion) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsSetDiscussionGroup) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsEditCreator) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsEditLocation) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsToggleSlowMode) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e ChannelsGetInactiveChannels) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e BotsSendCustomRequest) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e BotsAnswerWebhookJSONQuery) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e BotsSetBotCommands) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsGetPaymentForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsGetPaymentReceipt) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsValidateRequestedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsSendPaymentForm) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsGetSavedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsClearSavedInfo) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PaymentsGetBankCardData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StickersCreateStickerSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StickersRemoveStickerFromSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StickersChangeStickerPosition) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StickersAddStickerToSet) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StickersSetStickerSetThumb) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneGetCallConfig) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneRequestCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneAcceptCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneConfirmCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneReceivedCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneDiscardCall) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneSetCallRating) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneSaveCallDebug) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e PhoneSendSignalingData) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e LangpackGetLangPack) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e LangpackGetStrings) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e LangpackGetDifference) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e LangpackGetLanguages) decodeResponse(dbuf *DecodeBuf) TL {
	return VectorObject(dbuf.Vector())
}

func (e LangpackGetLanguage) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e FoldersEditPeerFolders) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e FoldersDeleteFolder) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StatsGetBroadcastStats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StatsLoadAsyncGraph) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func (e StatsGetMegagroupStats) decodeResponse(dbuf *DecodeBuf) TL {
	return dbuf.Object()
}

func readFlags(m *DecodeBuf, flagsPtr *int32) int32 {
	flags := m.Int()
	*flagsPtr = flags
	return flags
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crcResPQ:
		r = ResPQ{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
			m.VectorLong(),
		}

	case crcPQInnerDataDc:
		r = PQInnerDataDc{
			m.String(),
			m.String(),
			m.String(),
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(32),
			m.Int(),
		}

	case crcPQInnerDataTempDc:
		r = PQInnerDataTempDc{
			m.String(),
			m.String(),
			m.String(),
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(32),
			m.Int(),
			m.Int(),
		}

	case crcServerDHParamsFail:
		r = ServerDHParamsFail{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case crcServerDHParamsOk:
		r = ServerDHParamsOk{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
		}

	case crcServerDHInnerData:
		r = ServerDHInnerData{
			m.Bytes(16),
			m.Bytes(16),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crcClientDHInnerData:
		r = ClientDHInnerData{
			m.Bytes(16),
			m.Bytes(16),
			m.Long(),
			m.String(),
		}

	case crcDhGenOk:
		r = DhGenOk{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case crcDhGenRetry:
		r = DhGenRetry{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case crcDhGenFail:
		r = DhGenFail{
			m.Bytes(16),
			m.Bytes(16),
			m.Bytes(16),
		}

	case crcBindAuthKeyInner:
		r = BindAuthKeyInner{
			m.Long(),
			m.Long(),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crcRpcError:
		r = RpcError{
			m.Int(),
			m.String(),
		}

	case crcRpcAnswerUnknown:
		r = RpcAnswerUnknown{}

	case crcRpcAnswerDroppedRunning:
		r = RpcAnswerDroppedRunning{}

	case crcRpcAnswerDropped:
		r = RpcAnswerDropped{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crcFutureSalt:
		r = FutureSalt{
			m.Int(),
			m.Int(),
			m.Long(),
		}

	case crcFutureSalts:
		r = FutureSalts{
			m.Long(),
			m.Int(),
			m.Vector(),
		}

	case crcPong:
		r = Pong{
			m.Long(),
			m.Long(),
		}

	case crcDestroySessionOk:
		r = DestroySessionOk{
			m.Long(),
		}

	case crcDestroySessionNone:
		r = DestroySessionNone{
			m.Long(),
		}

	case crcNewSessionCreated:
		r = NewSessionCreated{
			m.Long(),
			m.Long(),
			m.Long(),
		}

	case crcMsgsAck:
		r = MsgsAck{
			m.VectorLong(),
		}

	case crcBadMsgNotification:
		r = BadMsgNotification{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crcBadServerSalt:
		r = BadServerSalt{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Long(),
		}

	case crcMsgResendReq:
		r = MsgResendReq{
			m.VectorLong(),
		}

	case crcMsgsStateReq:
		r = MsgsStateReq{
			m.VectorLong(),
		}

	case crcMsgsStateInfo:
		r = MsgsStateInfo{
			m.Long(),
			m.String(),
		}

	case crcMsgsAllInfo:
		r = MsgsAllInfo{
			m.VectorLong(),
			m.String(),
		}

	case crcMsgDetailedInfo:
		r = MsgDetailedInfo{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crcMsgNewDetailedInfo:
		r = MsgNewDetailedInfo{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crcDestroyAuthKeyOk:
		r = DestroyAuthKeyOk{}

	case crcDestroyAuthKeyNone:
		r = DestroyAuthKeyNone{}

	case crcDestroyAuthKeyFail:
		r = DestroyAuthKeyFail{}

	case crcReqPqMulti:
		r = ReqPqMulti{
			m.Bytes(16),
		}

	case crcReqDHParams:
		r = ReqDHParams{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crcSetClientDHParams:
		r = SetClientDHParams{
			m.Bytes(16),
			m.Bytes(16),
			m.String(),
		}

	case crcRpcDropAnswer:
		r = RpcDropAnswer{
			m.Long(),
		}

	case crcGetFutureSalts:
		r = GetFutureSalts{
			m.Int(),
		}

	case crcPing:
		r = Ping{
			m.Long(),
		}

	case crcPingDelayDisconnect:
		r = PingDelayDisconnect{
			m.Long(),
			m.Int(),
		}

	case crcDestroySession:
		r = DestroySession{
			m.Long(),
		}

	case crcHttpWait:
		r = HttpWait{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcDestroyAuthKey:
		r = DestroyAuthKey{}

	case crcTrue:
		r = True{}

	case crcBoolFalse:
		r = BoolFalse{}

	case crcBoolTrue:
		r = BoolTrue{}

	case crcError:
		r = Error{
			m.Int(),
			m.String(),
		}

	case crcIpPort:
		r = IpPort{
			m.Int(),
			m.Int(),
		}

	case crcInputPeerEmpty:
		r = InputPeerEmpty{}

	case crcInputPeerSelf:
		r = InputPeerSelf{}

	case crcInputPeerChat:
		r = InputPeerChat{
			m.Int(),
		}

	case crcInputPeerUser:
		r = InputPeerUser{
			m.Int(),
			m.Long(),
		}

	case crcInputPeerChannel:
		r = InputPeerChannel{
			m.Int(),
			m.Long(),
		}

	case crcInputPeerUserFromMessage:
		r = InputPeerUserFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcInputPeerChannelFromMessage:
		r = InputPeerChannelFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcInputUserEmpty:
		r = InputUserEmpty{}

	case crcInputUserSelf:
		r = InputUserSelf{}

	case crcInputUser:
		r = InputUser{
			m.Int(),
			m.Long(),
		}

	case crcInputUserFromMessage:
		r = InputUserFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcInputPhoneContact:
		r = InputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcInputFile:
		r = InputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crcInputFileBig:
		r = InputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crcInputMediaEmpty:
		r = InputMediaEmpty{}

	case crcInputMediaUploadedPhoto:
		var flags int32
		r = InputMediaUploadedPhoto{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crcInputMediaPhoto:
		var flags int32
		r = InputMediaPhoto{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case crcInputMediaGeoPoint:
		r = InputMediaGeoPoint{
			m.Object(),
		}

	case crcInputMediaContact:
		r = InputMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcInputMediaUploadedDocument:
		var flags int32
		r = InputMediaUploadedDocument{
			readFlags(m, &flags),
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.String(),
			m.Vector(),
			m.FlaggedVector(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crcInputMediaDocument:
		var flags int32
		r = InputMediaDocument{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case crcInputMediaVenue:
		r = InputMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcInputMediaPhotoExternal:
		var flags int32
		r = InputMediaPhotoExternal{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crcInputMediaDocumentExternal:
		var flags int32
		r = InputMediaDocumentExternal{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedInt(flags, 0),
		}

	case crcInputMediaGame:
		r = InputMediaGame{
			m.Object(),
		}

	case crcInputMediaInvoice:
		var flags int32
		r = InputMediaInvoice{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crcInputMediaGeoLive:
		var flags int32
		r = InputMediaGeoLive{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.FlaggedInt(flags, 1),
		}

	case crcInputMediaPoll:
		var flags int32
		r = InputMediaPoll{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedVector(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedVector(flags, 1),
		}

	case crcInputMediaDice:
		r = InputMediaDice{
			m.String(),
		}

	case crcInputChatPhotoEmpty:
		r = InputChatPhotoEmpty{}

	case crcInputChatUploadedPhoto:
		var flags int32
		r = InputChatUploadedPhoto{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedDouble(flags, 2),
		}

	case crcInputChatPhoto:
		r = InputChatPhoto{
			m.Object(),
		}

	case crcInputGeoPointEmpty:
		r = InputGeoPointEmpty{}

	case crcInputGeoPoint:
		r = InputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crcInputPhotoEmpty:
		r = InputPhotoEmpty{}

	case crcInputPhoto:
		r = InputPhoto{
			m.Long(),
			m.Long(),
			m.StringBytes(),
		}

	case crcInputFileLocation:
		r = InputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
			m.StringBytes(),
		}

	case crcInputEncryptedFileLocation:
		r = InputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crcInputDocumentFileLocation:
		r = InputDocumentFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.String(),
		}

	case crcInputSecureFileLocation:
		r = InputSecureFileLocation{
			m.Long(),
			m.Long(),
		}

	case crcInputTakeoutFileLocation:
		r = InputTakeoutFileLocation{}

	case crcInputPhotoFileLocation:
		r = InputPhotoFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.String(),
		}

	case crcInputPhotoLegacyFileLocation:
		r = InputPhotoLegacyFileLocation{
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crcInputPeerPhotoFileLocation:
		var flags int32
		r = InputPeerPhotoFileLocation{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case crcInputStickerSetThumb:
		r = InputStickerSetThumb{
			m.Object(),
			m.Long(),
			m.Int(),
		}

	case crcPeerUser:
		r = PeerUser{
			m.Int(),
		}

	case crcPeerChat:
		r = PeerChat{
			m.Int(),
		}

	case crcPeerChannel:
		r = PeerChannel{
			m.Int(),
		}

	case crcStorageFileUnknown:
		r = StorageFileUnknown{}

	case crcStorageFilePartial:
		r = StorageFilePartial{}

	case crcStorageFileJpeg:
		r = StorageFileJpeg{}

	case crcStorageFileGif:
		r = StorageFileGif{}

	case crcStorageFilePng:
		r = StorageFilePng{}

	case crcStorageFilePdf:
		r = StorageFilePdf{}

	case crcStorageFileMp3:
		r = StorageFileMp3{}

	case crcStorageFileMov:
		r = StorageFileMov{}

	case crcStorageFileMp4:
		r = StorageFileMp4{}

	case crcStorageFileWebp:
		r = StorageFileWebp{}

	case crcUserEmpty:
		r = UserEmpty{
			m.Int(),
		}

	case crcUser:
		var flags int32
		r = User{
			readFlags(m, &flags),
			flags&1024 != 0,     //flag #10
			flags&2048 != 0,     //flag #11
			flags&4096 != 0,     //flag #12
			flags&8192 != 0,     //flag #13
			flags&16384 != 0,    //flag #14
			flags&32768 != 0,    //flag #15
			flags&65536 != 0,    //flag #16
			flags&131072 != 0,   //flag #17
			flags&262144 != 0,   //flag #18
			flags&1048576 != 0,  //flag #20
			flags&2097152 != 0,  //flag #21
			flags&8388608 != 0,  //flag #23
			flags&16777216 != 0, //flag #24
			m.Int(),
			m.FlaggedLong(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 5),
			m.FlaggedObject(flags, 6),
			m.FlaggedInt(flags, 14),
			m.FlaggedVector(flags, 18),
			m.FlaggedString(flags, 19),
			m.FlaggedString(flags, 22),
		}

	case crcUserProfilePhotoEmpty:
		r = UserProfilePhotoEmpty{}

	case crcUserProfilePhoto:
		var flags int32
		r = UserProfilePhoto{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crcUserStatusEmpty:
		r = UserStatusEmpty{}

	case crcUserStatusOnline:
		r = UserStatusOnline{
			m.Int(),
		}

	case crcUserStatusOffline:
		r = UserStatusOffline{
			m.Int(),
		}

	case crcUserStatusRecently:
		r = UserStatusRecently{}

	case crcUserStatusLastWeek:
		r = UserStatusLastWeek{}

	case crcUserStatusLastMonth:
		r = UserStatusLastMonth{}

	case crcChatEmpty:
		r = ChatEmpty{
			m.Int(),
		}

	case crcChat:
		var flags int32
		r = Chat{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&32 != 0, //flag #5
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 6),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 18),
		}

	case crcChatForbidden:
		r = ChatForbidden{
			m.Int(),
			m.String(),
		}

	case crcChannel:
		var flags int32
		r = Channel{
			readFlags(m, &flags),
			flags&1 != 0,       //flag #0
			flags&4 != 0,       //flag #2
			flags&32 != 0,      //flag #5
			flags&128 != 0,     //flag #7
			flags&256 != 0,     //flag #8
			flags&512 != 0,     //flag #9
			flags&2048 != 0,    //flag #11
			flags&4096 != 0,    //flag #12
			flags&524288 != 0,  //flag #19
			flags&1048576 != 0, //flag #20
			flags&2097152 != 0, //flag #21
			flags&4194304 != 0, //flag #22
			m.Int(),
			m.FlaggedLong(flags, 13),
			m.String(),
			m.FlaggedString(flags, 6),
			m.Object(),
			m.Int(),
			m.Int(),
			m.FlaggedVector(flags, 9),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 15),
			m.FlaggedObject(flags, 18),
			m.FlaggedInt(flags, 17),
		}

	case crcChannelForbidden:
		var flags int32
		r = ChannelForbidden{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&256 != 0, //flag #8
			m.Int(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 16),
		}

	case crcChatFull:
		var flags int32
		r = ChatFull{
			readFlags(m, &flags),
			flags&128 != 0, //flag #7
			flags&256 != 0, //flag #8
			m.Int(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.Object(),
			m.Object(),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 11),
		}

	case crcChannelFull:
		var flags int32
		r = ChannelFull{
			readFlags(m, &flags),
			flags&8 != 0,      //flag #3
			flags&64 != 0,     //flag #6
			flags&128 != 0,    //flag #7
			flags&1024 != 0,   //flag #10
			flags&4096 != 0,   //flag #12
			flags&65536 != 0,  //flag #16
			flags&524288 != 0, //flag #19
			m.Int(),
			m.String(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 13),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
			m.FlaggedObject(flags, 8),
			m.FlaggedInt(flags, 9),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 14),
			m.FlaggedObject(flags, 15),
			m.FlaggedInt(flags, 17),
			m.FlaggedInt(flags, 18),
			m.FlaggedInt(flags, 12),
			m.Int(),
		}

	case crcChatParticipant:
		r = ChatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcChatParticipantCreator:
		r = ChatParticipantCreator{
			m.Int(),
		}

	case crcChatParticipantAdmin:
		r = ChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcChatParticipantsForbidden:
		var flags int32
		r = ChatParticipantsForbidden{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case crcChatParticipants:
		r = ChatParticipants{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crcChatPhotoEmpty:
		r = ChatPhotoEmpty{}

	case crcChatPhoto:
		var flags int32
		r = ChatPhoto{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crcMessageEmpty:
		r = MessageEmpty{
			m.Int(),
		}

	case crcMessage:
		var flags int32
		r = Message{
			readFlags(m, &flags),
			flags&2 != 0,       //flag #1
			flags&16 != 0,      //flag #4
			flags&32 != 0,      //flag #5
			flags&8192 != 0,    //flag #13
			flags&16384 != 0,   //flag #14
			flags&262144 != 0,  //flag #18
			flags&524288 != 0,  //flag #19
			flags&2097152 != 0, //flag #21
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 6),
			m.FlaggedVector(flags, 7),
			m.FlaggedInt(flags, 10),
			m.FlaggedInt(flags, 15),
			m.FlaggedString(flags, 16),
			m.FlaggedLong(flags, 17),
			m.FlaggedVector(flags, 22),
		}

	case crcMessageService:
		var flags int32
		r = MessageService{
			readFlags(m, &flags),
			flags&2 != 0,      //flag #1
			flags&16 != 0,     //flag #4
			flags&32 != 0,     //flag #5
			flags&8192 != 0,   //flag #13
			flags&16384 != 0,  //flag #14
			flags&524288 != 0, //flag #19
			m.Int(),
			m.FlaggedInt(flags, 8),
			m.Object(),
			m.FlaggedInt(flags, 3),
			m.Int(),
			m.Object(),
		}

	case crcMessageMediaEmpty:
		r = MessageMediaEmpty{}

	case crcMessageMediaPhoto:
		var flags int32
		r = MessageMediaPhoto{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
		}

	case crcMessageMediaGeo:
		r = MessageMediaGeo{
			m.Object(),
		}

	case crcMessageMediaContact:
		r = MessageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crcMessageMediaUnsupported:
		r = MessageMediaUnsupported{}

	case crcMessageMediaDocument:
		var flags int32
		r = MessageMediaDocument{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
		}

	case crcMessageMediaWebPage:
		r = MessageMediaWebPage{
			m.Object(),
		}

	case crcMessageMediaVenue:
		r = MessageMediaVenue{
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcMessageMediaGame:
		r = MessageMediaGame{
			m.Object(),
		}

	case crcMessageMediaInvoice:
		var flags int32
		r = MessageMediaInvoice{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			flags&8 != 0, //flag #3
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 2),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crcMessageMediaGeoLive:
		r = MessageMediaGeoLive{
			m.Object(),
			m.Int(),
		}

	case crcMessageMediaPoll:
		r = MessageMediaPoll{
			m.Object(),
			m.Object(),
		}

	case crcMessageMediaDice:
		r = MessageMediaDice{
			m.Int(),
			m.String(),
		}

	case crcMessageActionEmpty:
		r = MessageActionEmpty{}

	case crcMessageActionChatCreate:
		r = MessageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crcMessageActionChatEditTitle:
		r = MessageActionChatEditTitle{
			m.String(),
		}

	case crcMessageActionChatEditPhoto:
		r = MessageActionChatEditPhoto{
			m.Object(),
		}

	case crcMessageActionChatDeletePhoto:
		r = MessageActionChatDeletePhoto{}

	case crcMessageActionChatAddUser:
		r = MessageActionChatAddUser{
			m.VectorInt(),
		}

	case crcMessageActionChatDeleteUser:
		r = MessageActionChatDeleteUser{
			m.Int(),
		}

	case crcMessageActionChatJoinedByLink:
		r = MessageActionChatJoinedByLink{
			m.Int(),
		}

	case crcMessageActionChannelCreate:
		r = MessageActionChannelCreate{
			m.String(),
		}

	case crcMessageActionChatMigrateTo:
		r = MessageActionChatMigrateTo{
			m.Int(),
		}

	case crcMessageActionChannelMigrateFrom:
		r = MessageActionChannelMigrateFrom{
			m.String(),
			m.Int(),
		}

	case crcMessageActionPinMessage:
		r = MessageActionPinMessage{}

	case crcMessageActionHistoryClear:
		r = MessageActionHistoryClear{}

	case crcMessageActionGameScore:
		r = MessageActionGameScore{
			m.Long(),
			m.Int(),
		}

	case crcMessageActionPaymentSentMe:
		var flags int32
		r = MessageActionPaymentSentMe{
			readFlags(m, &flags),
			m.String(),
			m.Long(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case crcMessageActionPaymentSent:
		r = MessageActionPaymentSent{
			m.String(),
			m.Long(),
		}

	case crcMessageActionPhoneCall:
		var flags int32
		r = MessageActionPhoneCall{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crcMessageActionScreenshotTaken:
		r = MessageActionScreenshotTaken{}

	case crcMessageActionCustomAction:
		r = MessageActionCustomAction{
			m.String(),
		}

	case crcMessageActionBotAllowed:
		r = MessageActionBotAllowed{
			m.String(),
		}

	case crcMessageActionSecureValuesSentMe:
		r = MessageActionSecureValuesSentMe{
			m.Vector(),
			m.Object(),
		}

	case crcMessageActionSecureValuesSent:
		r = MessageActionSecureValuesSent{
			m.Vector(),
		}

	case crcMessageActionContactSignUp:
		r = MessageActionContactSignUp{}

	case crcDialog:
		var flags int32
		r = Dialog{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 4),
		}

	case crcDialogFolder:
		var flags int32
		r = DialogFolder{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcPhotoEmpty:
		r = PhotoEmpty{
			m.Long(),
		}

	case crcPhoto:
		var flags int32
		r = Photo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Int(),
			m.Vector(),
			m.FlaggedVector(flags, 1),
			m.Int(),
		}

	case crcPhotoSizeEmpty:
		r = PhotoSizeEmpty{
			m.String(),
		}

	case crcPhotoSize:
		r = PhotoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcPhotoCachedSize:
		r = PhotoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crcPhotoStrippedSize:
		r = PhotoStrippedSize{
			m.String(),
			m.StringBytes(),
		}

	case crcGeoPointEmpty:
		r = GeoPointEmpty{}

	case crcGeoPoint:
		r = GeoPoint{
			m.Double(),
			m.Double(),
			m.Long(),
		}

	case crcAuthSentCode:
		var flags int32
		r = AuthSentCode{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crcAuthAuthorization:
		var flags int32
		r = AuthAuthorization{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Object(),
		}

	case crcAuthAuthorizationSignUpRequired:
		var flags int32
		r = AuthAuthorizationSignUpRequired{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
		}

	case crcAuthExportedAuthorization:
		r = AuthExportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crcInputNotifyPeer:
		r = InputNotifyPeer{
			m.Object(),
		}

	case crcInputNotifyUsers:
		r = InputNotifyUsers{}

	case crcInputNotifyChats:
		r = InputNotifyChats{}

	case crcInputNotifyBroadcasts:
		r = InputNotifyBroadcasts{}

	case crcInputPeerNotifySettings:
		var flags int32
		r = InputPeerNotifySettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case crcPeerNotifySettings:
		var flags int32
		r = PeerNotifySettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
		}

	case crcPeerSettings:
		var flags int32
		r = PeerSettings{
			readFlags(m, &flags),
			flags&1 != 0,   //flag #0
			flags&2 != 0,   //flag #1
			flags&4 != 0,   //flag #2
			flags&8 != 0,   //flag #3
			flags&16 != 0,  //flag #4
			flags&32 != 0,  //flag #5
			flags&128 != 0, //flag #7
			m.FlaggedInt(flags, 6),
		}

	case crcWallPaper:
		var flags int32
		r = WallPaper{
			m.Long(),
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.Long(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 2),
		}

	case crcWallPaperNoFile:
		var flags int32
		r = WallPaperNoFile{
			readFlags(m, &flags),
			flags&2 != 0,  //flag #1
			flags&16 != 0, //flag #4
			m.FlaggedObject(flags, 2),
		}

	case crcInputReportReasonSpam:
		r = InputReportReasonSpam{}

	case crcInputReportReasonViolence:
		r = InputReportReasonViolence{}

	case crcInputReportReasonPornography:
		r = InputReportReasonPornography{}

	case crcInputReportReasonChildAbuse:
		r = InputReportReasonChildAbuse{}

	case crcInputReportReasonOther:
		r = InputReportReasonOther{
			m.String(),
		}

	case crcInputReportReasonCopyright:
		r = InputReportReasonCopyright{}

	case crcInputReportReasonGeoIrrelevant:
		r = InputReportReasonGeoIrrelevant{}

	case crcUserFull:
		var flags int32
		r = UserFull{
			readFlags(m, &flags),
			flags&1 != 0,    //flag #0
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&128 != 0,  //flag #7
			flags&4096 != 0, //flag #12
			m.Object(),
			m.FlaggedString(flags, 1),
			m.Object(),
			m.FlaggedObject(flags, 2),
			m.Object(),
			m.FlaggedObject(flags, 3),
			m.FlaggedInt(flags, 6),
			m.Int(),
			m.FlaggedInt(flags, 11),
		}

	case crcContact:
		r = Contact{
			m.Int(),
			m.Object(),
		}

	case crcImportedContact:
		r = ImportedContact{
			m.Int(),
			m.Long(),
		}

	case crcContactBlocked:
		r = ContactBlocked{
			m.Int(),
			m.Int(),
		}

	case crcContactStatus:
		r = ContactStatus{
			m.Int(),
			m.Object(),
		}

	case crcContactsContactsNotModified:
		r = ContactsContactsNotModified{}

	case crcContactsContacts:
		r = ContactsContacts{
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case crcContactsImportedContacts:
		r = ContactsImportedContacts{
			m.Vector(),
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case crcContactsBlocked:
		r = ContactsBlocked{
			m.Vector(),
			m.Vector(),
		}

	case crcContactsBlockedSlice:
		r = ContactsBlockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesDialogs:
		r = MessagesDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesDialogsSlice:
		r = MessagesDialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesDialogsNotModified:
		r = MessagesDialogsNotModified{
			m.Int(),
		}

	case crcMessagesMessages:
		r = MessagesMessages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesMessagesSlice:
		var flags int32
		r = MessagesMessagesSlice{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesChannelMessages:
		var flags int32
		r = MessagesChannelMessages{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesMessagesNotModified:
		r = MessagesMessagesNotModified{
			m.Int(),
		}

	case crcMessagesChats:
		r = MessagesChats{
			m.Vector(),
		}

	case crcMessagesChatsSlice:
		r = MessagesChatsSlice{
			m.Int(),
			m.Vector(),
		}

	case crcMessagesChatFull:
		r = MessagesChatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessagesAffectedHistory:
		r = MessagesAffectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcInputMessagesFilterEmpty:
		r = InputMessagesFilterEmpty{}

	case crcInputMessagesFilterPhotos:
		r = InputMessagesFilterPhotos{}

	case crcInputMessagesFilterVideo:
		r = InputMessagesFilterVideo{}

	case crcInputMessagesFilterPhotoVideo:
		r = InputMessagesFilterPhotoVideo{}

	case crcInputMessagesFilterDocument:
		r = InputMessagesFilterDocument{}

	case crcInputMessagesFilterUrl:
		r = InputMessagesFilterUrl{}

	case crcInputMessagesFilterGif:
		r = InputMessagesFilterGif{}

	case crcInputMessagesFilterVoice:
		r = InputMessagesFilterVoice{}

	case crcInputMessagesFilterMusic:
		r = InputMessagesFilterMusic{}

	case crcInputMessagesFilterChatPhotos:
		r = InputMessagesFilterChatPhotos{}

	case crcInputMessagesFilterPhoneCalls:
		var flags int32
		r = InputMessagesFilterPhoneCalls{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case crcInputMessagesFilterRoundVoice:
		r = InputMessagesFilterRoundVoice{}

	case crcInputMessagesFilterRoundVideo:
		r = InputMessagesFilterRoundVideo{}

	case crcInputMessagesFilterMyMentions:
		r = InputMessagesFilterMyMentions{}

	case crcInputMessagesFilterGeo:
		r = InputMessagesFilterGeo{}

	case crcInputMessagesFilterContacts:
		r = InputMessagesFilterContacts{}

	case crcUpdateNewMessage:
		r = UpdateNewMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateMessageID:
		r = UpdateMessageID{
			m.Int(),
			m.Long(),
		}

	case crcUpdateDeleteMessages:
		r = UpdateDeleteMessages{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateUserTyping:
		r = UpdateUserTyping{
			m.Int(),
			m.Object(),
		}

	case crcUpdateChatUserTyping:
		r = UpdateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crcUpdateChatParticipants:
		r = UpdateChatParticipants{
			m.Object(),
		}

	case crcUpdateUserStatus:
		r = UpdateUserStatus{
			m.Int(),
			m.Object(),
		}

	case crcUpdateUserName:
		r = UpdateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcUpdateUserPhoto:
		r = UpdateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crcUpdateNewEncryptedMessage:
		r = UpdateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case crcUpdateEncryptedChatTyping:
		r = UpdateEncryptedChatTyping{
			m.Int(),
		}

	case crcUpdateEncryption:
		r = UpdateEncryption{
			m.Object(),
			m.Int(),
		}

	case crcUpdateEncryptedMessagesRead:
		r = UpdateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateChatParticipantAdd:
		r = UpdateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateChatParticipantDelete:
		r = UpdateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateDcOptions:
		r = UpdateDcOptions{
			m.Vector(),
		}

	case crcUpdateUserBlocked:
		r = UpdateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case crcUpdateNotifySettings:
		r = UpdateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crcUpdateServiceNotification:
		var flags int32
		r = UpdateServiceNotification{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.String(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case crcUpdatePrivacy:
		r = UpdatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case crcUpdateUserPhone:
		r = UpdateUserPhone{
			m.Int(),
			m.String(),
		}

	case crcUpdateReadHistoryInbox:
		var flags int32
		r = UpdateReadHistoryInbox{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateReadHistoryOutbox:
		r = UpdateReadHistoryOutbox{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateWebPage:
		r = UpdateWebPage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateReadMessagesContents:
		r = UpdateReadMessagesContents{
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateChannelTooLong:
		var flags int32
		r = UpdateChannelTooLong{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedInt(flags, 0),
		}

	case crcUpdateChannel:
		r = UpdateChannel{
			m.Int(),
		}

	case crcUpdateNewChannelMessage:
		r = UpdateNewChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateReadChannelInbox:
		var flags int32
		r = UpdateReadChannelInbox{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateDeleteChannelMessages:
		r = UpdateDeleteChannelMessages{
			m.Int(),
			m.VectorInt(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateChannelMessageViews:
		r = UpdateChannelMessageViews{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateChatParticipantAdmin:
		r = UpdateChatParticipantAdmin{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crcUpdateNewStickerSet:
		r = UpdateNewStickerSet{
			m.Object(),
		}

	case crcUpdateStickerSetsOrder:
		var flags int32
		r = UpdateStickerSetsOrder{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.VectorLong(),
		}

	case crcUpdateStickerSets:
		r = UpdateStickerSets{}

	case crcUpdateSavedGifs:
		r = UpdateSavedGifs{}

	case crcUpdateBotInlineQuery:
		var flags int32
		r = UpdateBotInlineQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
		}

	case crcUpdateBotInlineSend:
		var flags int32
		r = UpdateBotInlineSend{
			readFlags(m, &flags),
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.FlaggedObject(flags, 1),
		}

	case crcUpdateEditChannelMessage:
		r = UpdateEditChannelMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateChannelPinnedMessage:
		r = UpdateChannelPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crcUpdateBotCallbackQuery:
		var flags int32
		r = UpdateBotCallbackQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crcUpdateEditMessage:
		r = UpdateEditMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateInlineBotCallbackQuery:
		var flags int32
		r = UpdateInlineBotCallbackQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.Object(),
			m.Long(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
		}

	case crcUpdateReadChannelOutbox:
		r = UpdateReadChannelOutbox{
			m.Int(),
			m.Int(),
		}

	case crcUpdateDraftMessage:
		r = UpdateDraftMessage{
			m.Object(),
			m.Object(),
		}

	case crcUpdateReadFeaturedStickers:
		r = UpdateReadFeaturedStickers{}

	case crcUpdateRecentStickers:
		r = UpdateRecentStickers{}

	case crcUpdateConfig:
		r = UpdateConfig{}

	case crcUpdatePtsChanged:
		r = UpdatePtsChanged{}

	case crcUpdateChannelWebPage:
		r = UpdateChannelWebPage{
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateDialogPinned:
		var flags int32
		r = UpdateDialogPinned{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.Object(),
		}

	case crcUpdatePinnedDialogs:
		var flags int32
		r = UpdatePinnedDialogs{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 1),
			m.FlaggedVector(flags, 0),
		}

	case crcUpdateBotWebhookJSON:
		r = UpdateBotWebhookJSON{
			m.Object(),
		}

	case crcUpdateBotWebhookJSONQuery:
		r = UpdateBotWebhookJSONQuery{
			m.Long(),
			m.Object(),
			m.Int(),
		}

	case crcUpdateBotShippingQuery:
		r = UpdateBotShippingQuery{
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crcUpdateBotPrecheckoutQuery:
		var flags int32
		r = UpdateBotPrecheckoutQuery{
			readFlags(m, &flags),
			m.Long(),
			m.Int(),
			m.StringBytes(),
			m.FlaggedObject(flags, 0),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Long(),
		}

	case crcUpdatePhoneCall:
		r = UpdatePhoneCall{
			m.Object(),
		}

	case crcUpdateLangPackTooLong:
		r = UpdateLangPackTooLong{
			m.String(),
		}

	case crcUpdateLangPack:
		r = UpdateLangPack{
			m.Object(),
		}

	case crcUpdateFavedStickers:
		r = UpdateFavedStickers{}

	case crcUpdateChannelReadMessagesContents:
		r = UpdateChannelReadMessagesContents{
			m.Int(),
			m.VectorInt(),
		}

	case crcUpdateContactsReset:
		r = UpdateContactsReset{}

	case crcUpdateChannelAvailableMessages:
		r = UpdateChannelAvailableMessages{
			m.Int(),
			m.Int(),
		}

	case crcUpdateDialogUnreadMark:
		var flags int32
		r = UpdateDialogUnreadMark{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case crcUpdateUserPinnedMessage:
		r = UpdateUserPinnedMessage{
			m.Int(),
			m.Int(),
		}

	case crcUpdateChatPinnedMessage:
		r = UpdateChatPinnedMessage{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateMessagePoll:
		var flags int32
		r = UpdateMessagePoll{
			readFlags(m, &flags),
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.Object(),
		}

	case crcUpdateChatDefaultBannedRights:
		r = UpdateChatDefaultBannedRights{
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crcUpdateFolderPeers:
		r = UpdateFolderPeers{
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crcUpdatePeerSettings:
		r = UpdatePeerSettings{
			m.Object(),
			m.Object(),
		}

	case crcUpdatePeerLocated:
		r = UpdatePeerLocated{
			m.Vector(),
		}

	case crcUpdateNewScheduledMessage:
		r = UpdateNewScheduledMessage{
			m.Object(),
		}

	case crcUpdateDeleteScheduledMessages:
		r = UpdateDeleteScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crcUpdateTheme:
		r = UpdateTheme{
			m.Object(),
		}

	case crcUpdateGeoLiveViewed:
		r = UpdateGeoLiveViewed{
			m.Object(),
			m.Int(),
		}

	case crcUpdateLoginToken:
		r = UpdateLoginToken{}

	case crcUpdateMessagePollVote:
		r = UpdateMessagePollVote{
			m.Long(),
			m.Int(),
			m.Vector(),
		}

	case crcUpdateDialogFilter:
		var flags int32
		r = UpdateDialogFilter{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case crcUpdateDialogFilterOrder:
		r = UpdateDialogFilterOrder{
			m.VectorInt(),
		}

	case crcUpdateDialogFilters:
		r = UpdateDialogFilters{}

	case crcUpdatePhoneCallSignalingData:
		r = UpdatePhoneCallSignalingData{
			m.Long(),
			m.StringBytes(),
		}

	case crcUpdateChannelParticipant:
		var flags int32
		r = UpdateChannelParticipant{
			readFlags(m, &flags),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Int(),
		}

	case crcUpdatesState:
		r = UpdatesState{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdatesDifferenceEmpty:
		r = UpdatesDifferenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crcUpdatesDifference:
		r = UpdatesDifference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crcUpdatesDifferenceSlice:
		r = UpdatesDifferenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crcUpdatesDifferenceTooLong:
		r = UpdatesDifferenceTooLong{
			m.Int(),
		}

	case crcUpdatesTooLong:
		r = UpdatesTooLong{}

	case crcUpdateShortMessage:
		var flags int32
		r = UpdateShortMessage{
			readFlags(m, &flags),
			flags&2 != 0,    //flag #1
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&8192 != 0, //flag #13
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case crcUpdateShortChatMessage:
		var flags int32
		r = UpdateShortChatMessage{
			readFlags(m, &flags),
			flags&2 != 0,    //flag #1
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&8192 != 0, //flag #13
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 2),
			m.FlaggedInt(flags, 11),
			m.FlaggedInt(flags, 3),
			m.FlaggedVector(flags, 7),
		}

	case crcUpdateShort:
		r = UpdateShort{
			m.Object(),
			m.Int(),
		}

	case crcUpdatesCombined:
		r = UpdatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdates:
		r = Updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crcUpdateShortSentMessage:
		var flags int32
		r = UpdateShortSentMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedObject(flags, 9),
			m.FlaggedVector(flags, 7),
		}

	case crcPhotosPhotos:
		r = PhotosPhotos{
			m.Vector(),
			m.Vector(),
		}

	case crcPhotosPhotosSlice:
		r = PhotosPhotosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crcPhotosPhoto:
		r = PhotosPhoto{
			m.Object(),
			m.Vector(),
		}

	case crcUploadFile:
		r = UploadFile{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crcUploadFileCdnRedirect:
		r = UploadFileCdnRedirect{
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
			m.Vector(),
		}

	case crcDcOption:
		var flags int32
		r = DcOption{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.Int(),
			m.String(),
			m.Int(),
			m.FlaggedStringBytes(flags, 10),
		}

	case crcConfig:
		var flags int32
		r = Config{
			readFlags(m, &flags),
			flags&2 != 0,    //flag #1
			flags&8 != 0,    //flag #3
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&64 != 0,   //flag #6
			flags&256 != 0,  //flag #8
			flags&8192 != 0, //flag #13
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 7),
			m.FlaggedString(flags, 9),
			m.FlaggedString(flags, 10),
			m.FlaggedString(flags, 11),
			m.FlaggedString(flags, 12),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedString(flags, 2),
			m.FlaggedInt(flags, 2),
			m.FlaggedInt(flags, 2),
		}

	case crcNearestDc:
		r = NearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crcHelpAppUpdate:
		var flags int32
		r = HelpAppUpdate{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.String(),
			m.String(),
			m.Vector(),
			m.FlaggedObject(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case crcHelpNoAppUpdate:
		r = HelpNoAppUpdate{}

	case crcHelpInviteText:
		r = HelpInviteText{
			m.String(),
		}

	case crcEncryptedChatEmpty:
		r = EncryptedChatEmpty{
			m.Int(),
		}

	case crcEncryptedChatWaiting:
		r = EncryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcEncryptedChatRequested:
		var flags int32
		r = EncryptedChatRequested{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crcEncryptedChat:
		r = EncryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crcEncryptedChatDiscarded:
		r = EncryptedChatDiscarded{
			m.Int(),
		}

	case crcInputEncryptedChat:
		r = InputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crcEncryptedFileEmpty:
		r = EncryptedFileEmpty{}

	case crcEncryptedFile:
		r = EncryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcInputEncryptedFileEmpty:
		r = InputEncryptedFileEmpty{}

	case crcInputEncryptedFileUploaded:
		r = InputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crcInputEncryptedFile:
		r = InputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crcInputEncryptedFileBigUploaded:
		r = InputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crcEncryptedMessage:
		r = EncryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crcEncryptedMessageService:
		r = EncryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crcMessagesDhConfigNotModified:
		r = MessagesDhConfigNotModified{
			m.StringBytes(),
		}

	case crcMessagesDhConfig:
		r = MessagesDhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crcMessagesSentEncryptedMessage:
		r = MessagesSentEncryptedMessage{
			m.Int(),
		}

	case crcMessagesSentEncryptedFile:
		r = MessagesSentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case crcInputDocumentEmpty:
		r = InputDocumentEmpty{}

	case crcInputDocument:
		r = InputDocument{
			m.Long(),
			m.Long(),
			m.StringBytes(),
		}

	case crcDocumentEmpty:
		r = DocumentEmpty{
			m.Long(),
		}

	case crcDocument:
		var flags int32
		r = Document{
			readFlags(m, &flags),
			m.Long(),
			m.Long(),
			m.StringBytes(),
			m.Int(),
			m.String(),
			m.Int(),
			m.FlaggedVector(flags, 0),
			m.FlaggedVector(flags, 1),
			m.Int(),
			m.Vector(),
		}

	case crcHelpSupport:
		r = HelpSupport{
			m.String(),
			m.Object(),
		}

	case crcNotifyPeer:
		r = NotifyPeer{
			m.Object(),
		}

	case crcNotifyUsers:
		r = NotifyUsers{}

	case crcNotifyChats:
		r = NotifyChats{}

	case crcNotifyBroadcasts:
		r = NotifyBroadcasts{}

	case crcSendMessageTypingAction:
		r = SendMessageTypingAction{}

	case crcSendMessageCancelAction:
		r = SendMessageCancelAction{}

	case crcSendMessageRecordVideoAction:
		r = SendMessageRecordVideoAction{}

	case crcSendMessageUploadVideoAction:
		r = SendMessageUploadVideoAction{
			m.Int(),
		}

	case crcSendMessageRecordAudioAction:
		r = SendMessageRecordAudioAction{}

	case crcSendMessageUploadAudioAction:
		r = SendMessageUploadAudioAction{
			m.Int(),
		}

	case crcSendMessageUploadPhotoAction:
		r = SendMessageUploadPhotoAction{
			m.Int(),
		}

	case crcSendMessageUploadDocumentAction:
		r = SendMessageUploadDocumentAction{
			m.Int(),
		}

	case crcSendMessageGeoLocationAction:
		r = SendMessageGeoLocationAction{}

	case crcSendMessageChooseContactAction:
		r = SendMessageChooseContactAction{}

	case crcSendMessageGamePlayAction:
		r = SendMessageGamePlayAction{}

	case crcSendMessageRecordRoundAction:
		r = SendMessageRecordRoundAction{}

	case crcSendMessageUploadRoundAction:
		r = SendMessageUploadRoundAction{
			m.Int(),
		}

	case crcContactsFound:
		r = ContactsFound{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcInputPrivacyKeyStatusTimestamp:
		r = InputPrivacyKeyStatusTimestamp{}

	case crcInputPrivacyKeyChatInvite:
		r = InputPrivacyKeyChatInvite{}

	case crcInputPrivacyKeyPhoneCall:
		r = InputPrivacyKeyPhoneCall{}

	case crcInputPrivacyKeyPhoneP2P:
		r = InputPrivacyKeyPhoneP2P{}

	case crcInputPrivacyKeyForwards:
		r = InputPrivacyKeyForwards{}

	case crcInputPrivacyKeyProfilePhoto:
		r = InputPrivacyKeyProfilePhoto{}

	case crcInputPrivacyKeyPhoneNumber:
		r = InputPrivacyKeyPhoneNumber{}

	case crcInputPrivacyKeyAddedByPhone:
		r = InputPrivacyKeyAddedByPhone{}

	case crcPrivacyKeyStatusTimestamp:
		r = PrivacyKeyStatusTimestamp{}

	case crcPrivacyKeyChatInvite:
		r = PrivacyKeyChatInvite{}

	case crcPrivacyKeyPhoneCall:
		r = PrivacyKeyPhoneCall{}

	case crcPrivacyKeyPhoneP2P:
		r = PrivacyKeyPhoneP2P{}

	case crcPrivacyKeyForwards:
		r = PrivacyKeyForwards{}

	case crcPrivacyKeyProfilePhoto:
		r = PrivacyKeyProfilePhoto{}

	case crcPrivacyKeyPhoneNumber:
		r = PrivacyKeyPhoneNumber{}

	case crcPrivacyKeyAddedByPhone:
		r = PrivacyKeyAddedByPhone{}

	case crcInputPrivacyValueAllowContacts:
		r = InputPrivacyValueAllowContacts{}

	case crcInputPrivacyValueAllowAll:
		r = InputPrivacyValueAllowAll{}

	case crcInputPrivacyValueAllowUsers:
		r = InputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case crcInputPrivacyValueDisallowContacts:
		r = InputPrivacyValueDisallowContacts{}

	case crcInputPrivacyValueDisallowAll:
		r = InputPrivacyValueDisallowAll{}

	case crcInputPrivacyValueDisallowUsers:
		r = InputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case crcInputPrivacyValueAllowChatParticipants:
		r = InputPrivacyValueAllowChatParticipants{
			m.VectorInt(),
		}

	case crcInputPrivacyValueDisallowChatParticipants:
		r = InputPrivacyValueDisallowChatParticipants{
			m.VectorInt(),
		}

	case crcPrivacyValueAllowContacts:
		r = PrivacyValueAllowContacts{}

	case crcPrivacyValueAllowAll:
		r = PrivacyValueAllowAll{}

	case crcPrivacyValueAllowUsers:
		r = PrivacyValueAllowUsers{
			m.VectorInt(),
		}

	case crcPrivacyValueDisallowContacts:
		r = PrivacyValueDisallowContacts{}

	case crcPrivacyValueDisallowAll:
		r = PrivacyValueDisallowAll{}

	case crcPrivacyValueDisallowUsers:
		r = PrivacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crcPrivacyValueAllowChatParticipants:
		r = PrivacyValueAllowChatParticipants{
			m.VectorInt(),
		}

	case crcPrivacyValueDisallowChatParticipants:
		r = PrivacyValueDisallowChatParticipants{
			m.VectorInt(),
		}

	case crcAccountPrivacyRules:
		r = AccountPrivacyRules{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcAccountDaysTTL:
		r = AccountDaysTTL{
			m.Int(),
		}

	case crcDocumentAttributeImageSize:
		r = DocumentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crcDocumentAttributeAnimated:
		r = DocumentAttributeAnimated{}

	case crcDocumentAttributeSticker:
		var flags int32
		r = DocumentAttributeSticker{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case crcDocumentAttributeVideo:
		var flags int32
		r = DocumentAttributeVideo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcDocumentAttributeAudio:
		var flags int32
		r = DocumentAttributeAudio{
			readFlags(m, &flags),
			flags&1024 != 0, //flag #10
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedStringBytes(flags, 2),
		}

	case crcDocumentAttributeFilename:
		r = DocumentAttributeFilename{
			m.String(),
		}

	case crcDocumentAttributeHasStickers:
		r = DocumentAttributeHasStickers{}

	case crcMessagesStickersNotModified:
		r = MessagesStickersNotModified{}

	case crcMessagesStickers:
		r = MessagesStickers{
			m.Int(),
			m.Vector(),
		}

	case crcStickerPack:
		r = StickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crcMessagesAllStickersNotModified:
		r = MessagesAllStickersNotModified{}

	case crcMessagesAllStickers:
		r = MessagesAllStickers{
			m.Int(),
			m.Vector(),
		}

	case crcMessagesAffectedMessages:
		r = MessagesAffectedMessages{
			m.Int(),
			m.Int(),
		}

	case crcWebPageEmpty:
		r = WebPageEmpty{
			m.Long(),
		}

	case crcWebPagePending:
		r = WebPagePending{
			m.Long(),
			m.Int(),
		}

	case crcWebPage:
		var flags int32
		r = WebPage{
			readFlags(m, &flags),
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedString(flags, 5),
			m.FlaggedString(flags, 5),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 6),
			m.FlaggedInt(flags, 7),
			m.FlaggedString(flags, 8),
			m.FlaggedObject(flags, 9),
			m.FlaggedObject(flags, 10),
			m.FlaggedVector(flags, 12),
		}

	case crcWebPageNotModified:
		var flags int32
		r = WebPageNotModified{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
		}

	case crcAuthorization:
		var flags int32
		r = Authorization{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcAccountAuthorizations:
		r = AccountAuthorizations{
			m.Vector(),
		}

	case crcAccountPassword:
		var flags int32
		r = AccountPassword{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.FlaggedObject(flags, 2),
			m.FlaggedStringBytes(flags, 2),
			m.FlaggedLong(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.Object(),
			m.Object(),
			m.StringBytes(),
		}

	case crcAccountPasswordSettings:
		var flags int32
		r = AccountPasswordSettings{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case crcAccountPasswordInputSettings:
		var flags int32
		r = AccountPasswordInputSettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crcAuthPasswordRecovery:
		r = AuthPasswordRecovery{
			m.String(),
		}

	case crcReceivedNotifyMessage:
		r = ReceivedNotifyMessage{
			m.Int(),
			m.Int(),
		}

	case crcChatInviteEmpty:
		r = ChatInviteEmpty{}

	case crcChatInviteExported:
		r = ChatInviteExported{
			m.String(),
		}

	case crcChatInviteAlready:
		r = ChatInviteAlready{
			m.Object(),
		}

	case crcChatInvite:
		var flags int32
		r = ChatInvite{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.String(),
			m.Object(),
			m.Int(),
			m.FlaggedVector(flags, 4),
		}

	case crcChatInvitePeek:
		r = ChatInvitePeek{
			m.Object(),
			m.Int(),
		}

	case crcInputStickerSetEmpty:
		r = InputStickerSetEmpty{}

	case crcInputStickerSetID:
		r = InputStickerSetID{
			m.Long(),
			m.Long(),
		}

	case crcInputStickerSetShortName:
		r = InputStickerSetShortName{
			m.String(),
		}

	case crcInputStickerSetAnimatedEmoji:
		r = InputStickerSetAnimatedEmoji{}

	case crcInputStickerSetDice:
		r = InputStickerSetDice{
			m.String(),
		}

	case crcStickerSet:
		var flags int32
		r = StickerSet{
			readFlags(m, &flags),
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&32 != 0, //flag #5
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 4),
			m.FlaggedInt(flags, 4),
			m.Int(),
			m.Int(),
		}

	case crcMessagesStickerSet:
		r = MessagesStickerSet{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crcBotCommand:
		r = BotCommand{
			m.String(),
			m.String(),
		}

	case crcBotInfo:
		r = BotInfo{
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crcKeyboardButton:
		r = KeyboardButton{
			m.String(),
		}

	case crcKeyboardButtonUrl:
		r = KeyboardButtonUrl{
			m.String(),
			m.String(),
		}

	case crcKeyboardButtonCallback:
		r = KeyboardButtonCallback{
			m.String(),
			m.StringBytes(),
		}

	case crcKeyboardButtonRequestPhone:
		r = KeyboardButtonRequestPhone{
			m.String(),
		}

	case crcKeyboardButtonRequestGeoLocation:
		r = KeyboardButtonRequestGeoLocation{
			m.String(),
		}

	case crcKeyboardButtonSwitchInline:
		var flags int32
		r = KeyboardButtonSwitchInline{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.String(),
		}

	case crcKeyboardButtonGame:
		r = KeyboardButtonGame{
			m.String(),
		}

	case crcKeyboardButtonBuy:
		r = KeyboardButtonBuy{
			m.String(),
		}

	case crcKeyboardButtonUrlAuth:
		var flags int32
		r = KeyboardButtonUrlAuth{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedString(flags, 0),
			m.String(),
			m.Int(),
		}

	case crcInputKeyboardButtonUrlAuth:
		var flags int32
		r = InputKeyboardButtonUrlAuth{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Object(),
		}

	case crcKeyboardButtonRequestPoll:
		var flags int32
		r = KeyboardButtonRequestPoll{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.String(),
		}

	case crcKeyboardButtonRow:
		r = KeyboardButtonRow{
			m.Vector(),
		}

	case crcReplyKeyboardHide:
		var flags int32
		r = ReplyKeyboardHide{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
		}

	case crcReplyKeyboardForceReply:
		var flags int32
		r = ReplyKeyboardForceReply{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
		}

	case crcReplyKeyboardMarkup:
		var flags int32
		r = ReplyKeyboardMarkup{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Vector(),
		}

	case crcReplyInlineMarkup:
		r = ReplyInlineMarkup{
			m.Vector(),
		}

	case crcMessageEntityUnknown:
		r = MessageEntityUnknown{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityMention:
		r = MessageEntityMention{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityHashtag:
		r = MessageEntityHashtag{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityBotCommand:
		r = MessageEntityBotCommand{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityUrl:
		r = MessageEntityUrl{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityEmail:
		r = MessageEntityEmail{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityBold:
		r = MessageEntityBold{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityItalic:
		r = MessageEntityItalic{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityCode:
		r = MessageEntityCode{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityPre:
		r = MessageEntityPre{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crcMessageEntityTextUrl:
		r = MessageEntityTextUrl{
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crcMessageEntityMentionName:
		r = MessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcInputMessageEntityMentionName:
		r = InputMessageEntityMentionName{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crcMessageEntityPhone:
		r = MessageEntityPhone{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityCashtag:
		r = MessageEntityCashtag{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityUnderline:
		r = MessageEntityUnderline{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityStrike:
		r = MessageEntityStrike{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityBlockquote:
		r = MessageEntityBlockquote{
			m.Int(),
			m.Int(),
		}

	case crcMessageEntityBankCard:
		r = MessageEntityBankCard{
			m.Int(),
			m.Int(),
		}

	case crcInputChannelEmpty:
		r = InputChannelEmpty{}

	case crcInputChannel:
		r = InputChannel{
			m.Int(),
			m.Long(),
		}

	case crcInputChannelFromMessage:
		r = InputChannelFromMessage{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcContactsResolvedPeer:
		r = ContactsResolvedPeer{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crcMessageRange:
		r = MessageRange{
			m.Int(),
			m.Int(),
		}

	case crcUpdatesChannelDifferenceEmpty:
		var flags int32
		r = UpdatesChannelDifferenceEmpty{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.FlaggedInt(flags, 1),
		}

	case crcUpdatesChannelDifferenceTooLong:
		var flags int32
		r = UpdatesChannelDifferenceTooLong{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcUpdatesChannelDifference:
		var flags int32
		r = UpdatesChannelDifference{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcChannelMessagesFilterEmpty:
		r = ChannelMessagesFilterEmpty{}

	case crcChannelMessagesFilter:
		var flags int32
		r = ChannelMessagesFilter{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Vector(),
		}

	case crcChannelParticipant:
		r = ChannelParticipant{
			m.Int(),
			m.Int(),
		}

	case crcChannelParticipantSelf:
		r = ChannelParticipantSelf{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcChannelParticipantCreator:
		var flags int32
		r = ChannelParticipantCreator{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedString(flags, 0),
		}

	case crcChannelParticipantAdmin:
		var flags int32
		r = ChannelParticipantAdmin{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedString(flags, 2),
		}

	case crcChannelParticipantBanned:
		var flags int32
		r = ChannelParticipantBanned{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crcChannelParticipantsRecent:
		r = ChannelParticipantsRecent{}

	case crcChannelParticipantsAdmins:
		r = ChannelParticipantsAdmins{}

	case crcChannelParticipantsKicked:
		r = ChannelParticipantsKicked{
			m.String(),
		}

	case crcChannelParticipantsBots:
		r = ChannelParticipantsBots{}

	case crcChannelParticipantsBanned:
		r = ChannelParticipantsBanned{
			m.String(),
		}

	case crcChannelParticipantsSearch:
		r = ChannelParticipantsSearch{
			m.String(),
		}

	case crcChannelParticipantsContacts:
		r = ChannelParticipantsContacts{
			m.String(),
		}

	case crcChannelsChannelParticipants:
		r = ChannelsChannelParticipants{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crcChannelsChannelParticipantsNotModified:
		r = ChannelsChannelParticipantsNotModified{}

	case crcChannelsChannelParticipant:
		r = ChannelsChannelParticipant{
			m.Object(),
			m.Vector(),
		}

	case crcHelpTermsOfService:
		var flags int32
		r = HelpTermsOfService{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
			m.Vector(),
			m.FlaggedInt(flags, 1),
		}

	case crcMessagesSavedGifsNotModified:
		r = MessagesSavedGifsNotModified{}

	case crcMessagesSavedGifs:
		r = MessagesSavedGifs{
			m.Int(),
			m.Vector(),
		}

	case crcInputBotInlineMessageMediaAuto:
		var flags int32
		r = InputBotInlineMessageMediaAuto{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crcInputBotInlineMessageText:
		var flags int32
		r = InputBotInlineMessageText{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crcInputBotInlineMessageMediaGeo:
		var flags int32
		r = InputBotInlineMessageMediaGeo{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 2),
		}

	case crcInputBotInlineMessageMediaVenue:
		var flags int32
		r = InputBotInlineMessageMediaVenue{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crcInputBotInlineMessageMediaContact:
		var flags int32
		r = InputBotInlineMessageMediaContact{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crcInputBotInlineMessageGame:
		var flags int32
		r = InputBotInlineMessageGame{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 2),
		}

	case crcInputBotInlineResult:
		var flags int32
		r = InputBotInlineResult{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 5),
			m.Object(),
		}

	case crcInputBotInlineResultPhoto:
		r = InputBotInlineResultPhoto{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crcInputBotInlineResultDocument:
		var flags int32
		r = InputBotInlineResultDocument{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.Object(),
			m.Object(),
		}

	case crcInputBotInlineResultGame:
		r = InputBotInlineResultGame{
			m.String(),
			m.String(),
			m.Object(),
		}

	case crcBotInlineMessageMediaAuto:
		var flags int32
		r = BotInlineMessageMediaAuto{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crcBotInlineMessageText:
		var flags int32
		r = BotInlineMessageText{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedVector(flags, 1),
			m.FlaggedObject(flags, 2),
		}

	case crcBotInlineMessageMediaGeo:
		var flags int32
		r = BotInlineMessageMediaGeo{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 2),
		}

	case crcBotInlineMessageMediaVenue:
		var flags int32
		r = BotInlineMessageMediaVenue{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crcBotInlineMessageMediaContact:
		var flags int32
		r = BotInlineMessageMediaContact{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
		}

	case crcBotInlineResult:
		var flags int32
		r = BotInlineResult{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 5),
			m.Object(),
		}

	case crcBotInlineMediaResult:
		var flags int32
		r = BotInlineMediaResult{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.Object(),
		}

	case crcMessagesBotResults:
		var flags int32
		r = MessagesBotResults{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
			m.Vector(),
			m.Int(),
			m.Vector(),
		}

	case crcExportedMessageLink:
		r = ExportedMessageLink{
			m.String(),
			m.String(),
		}

	case crcMessageFwdHeader:
		var flags int32
		r = MessageFwdHeader{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.FlaggedString(flags, 5),
			m.Int(),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedObject(flags, 4),
			m.FlaggedInt(flags, 4),
			m.FlaggedString(flags, 6),
		}

	case crcAuthCodeTypeSms:
		r = AuthCodeTypeSms{}

	case crcAuthCodeTypeCall:
		r = AuthCodeTypeCall{}

	case crcAuthCodeTypeFlashCall:
		r = AuthCodeTypeFlashCall{}

	case crcAuthSentCodeTypeApp:
		r = AuthSentCodeTypeApp{
			m.Int(),
		}

	case crcAuthSentCodeTypeSms:
		r = AuthSentCodeTypeSms{
			m.Int(),
		}

	case crcAuthSentCodeTypeCall:
		r = AuthSentCodeTypeCall{
			m.Int(),
		}

	case crcAuthSentCodeTypeFlashCall:
		r = AuthSentCodeTypeFlashCall{
			m.String(),
		}

	case crcMessagesBotCallbackAnswer:
		var flags int32
		r = MessagesBotCallbackAnswer{
			readFlags(m, &flags),
			flags&2 != 0,  //flag #1
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crcMessagesMessageEditData:
		var flags int32
		r = MessagesMessageEditData{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case crcInputBotInlineMessageID:
		r = InputBotInlineMessageID{
			m.Int(),
			m.Long(),
			m.Long(),
		}

	case crcInlineBotSwitchPM:
		r = InlineBotSwitchPM{
			m.String(),
			m.String(),
		}

	case crcMessagesPeerDialogs:
		r = MessagesPeerDialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crcTopPeer:
		r = TopPeer{
			m.Object(),
			m.Double(),
		}

	case crcTopPeerCategoryBotsPM:
		r = TopPeerCategoryBotsPM{}

	case crcTopPeerCategoryBotsInline:
		r = TopPeerCategoryBotsInline{}

	case crcTopPeerCategoryCorrespondents:
		r = TopPeerCategoryCorrespondents{}

	case crcTopPeerCategoryGroups:
		r = TopPeerCategoryGroups{}

	case crcTopPeerCategoryChannels:
		r = TopPeerCategoryChannels{}

	case crcTopPeerCategoryPhoneCalls:
		r = TopPeerCategoryPhoneCalls{}

	case crcTopPeerCategoryForwardUsers:
		r = TopPeerCategoryForwardUsers{}

	case crcTopPeerCategoryForwardChats:
		r = TopPeerCategoryForwardChats{}

	case crcTopPeerCategoryPeers:
		r = TopPeerCategoryPeers{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crcContactsTopPeersNotModified:
		r = ContactsTopPeersNotModified{}

	case crcContactsTopPeers:
		r = ContactsTopPeers{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcContactsTopPeersDisabled:
		r = ContactsTopPeersDisabled{}

	case crcDraftMessageEmpty:
		var flags int32
		r = DraftMessageEmpty{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
		}

	case crcDraftMessage:
		var flags int32
		r = DraftMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedInt(flags, 0),
			m.String(),
			m.FlaggedVector(flags, 3),
			m.Int(),
		}

	case crcMessagesFeaturedStickersNotModified:
		r = MessagesFeaturedStickersNotModified{
			m.Int(),
		}

	case crcMessagesFeaturedStickers:
		r = MessagesFeaturedStickers{
			m.Int(),
			m.Int(),
			m.Vector(),
			m.VectorLong(),
		}

	case crcMessagesRecentStickersNotModified:
		r = MessagesRecentStickersNotModified{}

	case crcMessagesRecentStickers:
		r = MessagesRecentStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.VectorInt(),
		}

	case crcMessagesArchivedStickers:
		r = MessagesArchivedStickers{
			m.Int(),
			m.Vector(),
		}

	case crcMessagesStickerSetInstallResultSuccess:
		r = MessagesStickerSetInstallResultSuccess{}

	case crcMessagesStickerSetInstallResultArchive:
		r = MessagesStickerSetInstallResultArchive{
			m.Vector(),
		}

	case crcStickerSetCovered:
		r = StickerSetCovered{
			m.Object(),
			m.Object(),
		}

	case crcStickerSetMultiCovered:
		r = StickerSetMultiCovered{
			m.Object(),
			m.Vector(),
		}

	case crcMaskCoords:
		r = MaskCoords{
			m.Int(),
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crcInputStickeredMediaPhoto:
		r = InputStickeredMediaPhoto{
			m.Object(),
		}

	case crcInputStickeredMediaDocument:
		r = InputStickeredMediaDocument{
			m.Object(),
		}

	case crcGame:
		var flags int32
		r = Game{
			readFlags(m, &flags),
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.FlaggedObject(flags, 0),
		}

	case crcInputGameID:
		r = InputGameID{
			m.Long(),
			m.Long(),
		}

	case crcInputGameShortName:
		r = InputGameShortName{
			m.Object(),
			m.String(),
		}

	case crcHighScore:
		r = HighScore{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesHighScores:
		r = MessagesHighScores{
			m.Vector(),
			m.Vector(),
		}

	case crcTextEmpty:
		r = TextEmpty{}

	case crcTextPlain:
		r = TextPlain{
			m.String(),
		}

	case crcTextBold:
		r = TextBold{
			m.Object(),
		}

	case crcTextItalic:
		r = TextItalic{
			m.Object(),
		}

	case crcTextUnderline:
		r = TextUnderline{
			m.Object(),
		}

	case crcTextStrike:
		r = TextStrike{
			m.Object(),
		}

	case crcTextFixed:
		r = TextFixed{
			m.Object(),
		}

	case crcTextUrl:
		r = TextUrl{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crcTextEmail:
		r = TextEmail{
			m.Object(),
			m.String(),
		}

	case crcTextConcat:
		r = TextConcat{
			m.Vector(),
		}

	case crcTextSubscript:
		r = TextSubscript{
			m.Object(),
		}

	case crcTextSuperscript:
		r = TextSuperscript{
			m.Object(),
		}

	case crcTextMarked:
		r = TextMarked{
			m.Object(),
		}

	case crcTextPhone:
		r = TextPhone{
			m.Object(),
			m.String(),
		}

	case crcTextImage:
		r = TextImage{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crcTextAnchor:
		r = TextAnchor{
			m.Object(),
			m.String(),
		}

	case crcPageBlockUnsupported:
		r = PageBlockUnsupported{}

	case crcPageBlockTitle:
		r = PageBlockTitle{
			m.Object(),
		}

	case crcPageBlockSubtitle:
		r = PageBlockSubtitle{
			m.Object(),
		}

	case crcPageBlockAuthorDate:
		r = PageBlockAuthorDate{
			m.Object(),
			m.Int(),
		}

	case crcPageBlockHeader:
		r = PageBlockHeader{
			m.Object(),
		}

	case crcPageBlockSubheader:
		r = PageBlockSubheader{
			m.Object(),
		}

	case crcPageBlockParagraph:
		r = PageBlockParagraph{
			m.Object(),
		}

	case crcPageBlockPreformatted:
		r = PageBlockPreformatted{
			m.Object(),
			m.String(),
		}

	case crcPageBlockFooter:
		r = PageBlockFooter{
			m.Object(),
		}

	case crcPageBlockDivider:
		r = PageBlockDivider{}

	case crcPageBlockAnchor:
		r = PageBlockAnchor{
			m.String(),
		}

	case crcPageBlockList:
		r = PageBlockList{
			m.Vector(),
		}

	case crcPageBlockBlockquote:
		r = PageBlockBlockquote{
			m.Object(),
			m.Object(),
		}

	case crcPageBlockPullquote:
		r = PageBlockPullquote{
			m.Object(),
			m.Object(),
		}

	case crcPageBlockPhoto:
		var flags int32
		r = PageBlockPhoto{
			readFlags(m, &flags),
			m.Long(),
			m.Object(),
			m.FlaggedString(flags, 0),
			m.FlaggedLong(flags, 0),
		}

	case crcPageBlockVideo:
		var flags int32
		r = PageBlockVideo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Long(),
			m.Object(),
		}

	case crcPageBlockCover:
		r = PageBlockCover{
			m.Object(),
		}

	case crcPageBlockEmbed:
		var flags int32
		r = PageBlockEmbed{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&8 != 0, //flag #3
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedLong(flags, 4),
			m.FlaggedInt(flags, 5),
			m.FlaggedInt(flags, 5),
			m.Object(),
		}

	case crcPageBlockEmbedPost:
		r = PageBlockEmbedPost{
			m.String(),
			m.Long(),
			m.Long(),
			m.String(),
			m.Int(),
			m.Vector(),
			m.Object(),
		}

	case crcPageBlockCollage:
		r = PageBlockCollage{
			m.Vector(),
			m.Object(),
		}

	case crcPageBlockSlideshow:
		r = PageBlockSlideshow{
			m.Vector(),
			m.Object(),
		}

	case crcPageBlockChannel:
		r = PageBlockChannel{
			m.Object(),
		}

	case crcPageBlockAudio:
		r = PageBlockAudio{
			m.Long(),
			m.Object(),
		}

	case crcPageBlockKicker:
		r = PageBlockKicker{
			m.Object(),
		}

	case crcPageBlockTable:
		var flags int32
		r = PageBlockTable{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Vector(),
		}

	case crcPageBlockOrderedList:
		r = PageBlockOrderedList{
			m.Vector(),
		}

	case crcPageBlockDetails:
		var flags int32
		r = PageBlockDetails{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Vector(),
			m.Object(),
		}

	case crcPageBlockRelatedArticles:
		r = PageBlockRelatedArticles{
			m.Object(),
			m.Vector(),
		}

	case crcPageBlockMap:
		r = PageBlockMap{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crcPhoneCallDiscardReasonMissed:
		r = PhoneCallDiscardReasonMissed{}

	case crcPhoneCallDiscardReasonDisconnect:
		r = PhoneCallDiscardReasonDisconnect{}

	case crcPhoneCallDiscardReasonHangup:
		r = PhoneCallDiscardReasonHangup{}

	case crcPhoneCallDiscardReasonBusy:
		r = PhoneCallDiscardReasonBusy{}

	case crcDataJSON:
		r = DataJSON{
			m.String(),
		}

	case crcLabeledPrice:
		r = LabeledPrice{
			m.String(),
			m.Long(),
		}

	case crcInvoice:
		var flags int32
		r = Invoice{
			readFlags(m, &flags),
			flags&1 != 0,   //flag #0
			flags&2 != 0,   //flag #1
			flags&4 != 0,   //flag #2
			flags&8 != 0,   //flag #3
			flags&16 != 0,  //flag #4
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.String(),
			m.Vector(),
		}

	case crcPaymentCharge:
		r = PaymentCharge{
			m.String(),
			m.String(),
		}

	case crcPostAddress:
		r = PostAddress{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcPaymentRequestedInfo:
		var flags int32
		r = PaymentRequestedInfo{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crcPaymentSavedCredentialsCard:
		r = PaymentSavedCredentialsCard{
			m.String(),
			m.String(),
		}

	case crcWebDocument:
		r = WebDocument{
			m.String(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crcWebDocumentNoProxy:
		r = WebDocumentNoProxy{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crcInputWebDocument:
		r = InputWebDocument{
			m.String(),
			m.Int(),
			m.String(),
			m.Vector(),
		}

	case crcInputWebFileLocation:
		r = InputWebFileLocation{
			m.String(),
			m.Long(),
		}

	case crcInputWebFileGeoPointLocation:
		r = InputWebFileGeoPointLocation{
			m.Object(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUploadWebFile:
		r = UploadWebFile{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crcPaymentsPaymentForm:
		var flags int32
		r = PaymentsPaymentForm{
			readFlags(m, &flags),
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 4),
			m.FlaggedObject(flags, 4),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Vector(),
		}

	case crcPaymentsValidatedRequestedInfo:
		var flags int32
		r = PaymentsValidatedRequestedInfo{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case crcPaymentsPaymentResult:
		r = PaymentsPaymentResult{
			m.Object(),
		}

	case crcPaymentsPaymentVerificationNeeded:
		r = PaymentsPaymentVerificationNeeded{
			m.String(),
		}

	case crcPaymentsPaymentReceipt:
		var flags int32
		r = PaymentsPaymentReceipt{
			readFlags(m, &flags),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.String(),
			m.Long(),
			m.String(),
			m.Vector(),
		}

	case crcPaymentsSavedInfo:
		var flags int32
		r = PaymentsSavedInfo{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedObject(flags, 0),
		}

	case crcInputPaymentCredentialsSaved:
		r = InputPaymentCredentialsSaved{
			m.String(),
			m.StringBytes(),
		}

	case crcInputPaymentCredentials:
		var flags int32
		r = InputPaymentCredentials{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case crcInputPaymentCredentialsApplePay:
		r = InputPaymentCredentialsApplePay{
			m.Object(),
		}

	case crcInputPaymentCredentialsAndroidPay:
		r = InputPaymentCredentialsAndroidPay{
			m.Object(),
			m.String(),
		}

	case crcAccountTmpPassword:
		r = AccountTmpPassword{
			m.StringBytes(),
			m.Int(),
		}

	case crcShippingOption:
		r = ShippingOption{
			m.String(),
			m.String(),
			m.Vector(),
		}

	case crcInputStickerSetItem:
		var flags int32
		r = InputStickerSetItem{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
		}

	case crcInputPhoneCall:
		r = InputPhoneCall{
			m.Long(),
			m.Long(),
		}

	case crcPhoneCallEmpty:
		r = PhoneCallEmpty{
			m.Long(),
		}

	case crcPhoneCallWaiting:
		var flags int32
		r = PhoneCallWaiting{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case crcPhoneCallRequested:
		var flags int32
		r = PhoneCallRequested{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crcPhoneCallAccepted:
		var flags int32
		r = PhoneCallAccepted{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crcPhoneCall:
		var flags int32
		r = PhoneCall{
			readFlags(m, &flags),
			flags&32 != 0, //flag #5
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
			m.Vector(),
			m.Int(),
		}

	case crcPhoneCallDiscarded:
		var flags int32
		r = PhoneCallDiscarded{
			readFlags(m, &flags),
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&32 != 0, //flag #5
			m.Long(),
			m.FlaggedObject(flags, 0),
			m.FlaggedInt(flags, 1),
		}

	case crcPhoneConnection:
		r = PhoneConnection{
			m.Long(),
			m.String(),
			m.String(),
			m.Int(),
			m.StringBytes(),
		}

	case crcPhoneCallProtocol:
		var flags int32
		r = PhoneCallProtocol{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Int(),
			m.Int(),
			m.VectorString(),
		}

	case crcPhonePhoneCall:
		r = PhonePhoneCall{
			m.Object(),
			m.Vector(),
		}

	case crcUploadCdnFileReuploadNeeded:
		r = UploadCdnFileReuploadNeeded{
			m.StringBytes(),
		}

	case crcUploadCdnFile:
		r = UploadCdnFile{
			m.StringBytes(),
		}

	case crcCdnPublicKey:
		r = CdnPublicKey{
			m.Int(),
			m.String(),
		}

	case crcCdnConfig:
		r = CdnConfig{
			m.Vector(),
		}

	case crcLangPackString:
		r = LangPackString{
			m.String(),
			m.String(),
		}

	case crcLangPackStringPluralized:
		var flags int32
		r = LangPackStringPluralized{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedString(flags, 4),
			m.String(),
		}

	case crcLangPackStringDeleted:
		r = LangPackStringDeleted{
			m.String(),
		}

	case crcLangPackDifference:
		r = LangPackDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crcLangPackLanguage:
		var flags int32
		r = LangPackLanguage{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedString(flags, 1),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crcChannelAdminLogEventActionChangeTitle:
		r = ChannelAdminLogEventActionChangeTitle{
			m.String(),
			m.String(),
		}

	case crcChannelAdminLogEventActionChangeAbout:
		r = ChannelAdminLogEventActionChangeAbout{
			m.String(),
			m.String(),
		}

	case crcChannelAdminLogEventActionChangeUsername:
		r = ChannelAdminLogEventActionChangeUsername{
			m.String(),
			m.String(),
		}

	case crcChannelAdminLogEventActionChangePhoto:
		r = ChannelAdminLogEventActionChangePhoto{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionToggleInvites:
		r = ChannelAdminLogEventActionToggleInvites{
			m.Object(),
		}

	case crcChannelAdminLogEventActionToggleSignatures:
		r = ChannelAdminLogEventActionToggleSignatures{
			m.Object(),
		}

	case crcChannelAdminLogEventActionUpdatePinned:
		r = ChannelAdminLogEventActionUpdatePinned{
			m.Object(),
		}

	case crcChannelAdminLogEventActionEditMessage:
		r = ChannelAdminLogEventActionEditMessage{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionDeleteMessage:
		r = ChannelAdminLogEventActionDeleteMessage{
			m.Object(),
		}

	case crcChannelAdminLogEventActionParticipantJoin:
		r = ChannelAdminLogEventActionParticipantJoin{}

	case crcChannelAdminLogEventActionParticipantLeave:
		r = ChannelAdminLogEventActionParticipantLeave{}

	case crcChannelAdminLogEventActionParticipantInvite:
		r = ChannelAdminLogEventActionParticipantInvite{
			m.Object(),
		}

	case crcChannelAdminLogEventActionParticipantToggleBan:
		r = ChannelAdminLogEventActionParticipantToggleBan{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionParticipantToggleAdmin:
		r = ChannelAdminLogEventActionParticipantToggleAdmin{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionChangeStickerSet:
		r = ChannelAdminLogEventActionChangeStickerSet{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionTogglePreHistoryHidden:
		r = ChannelAdminLogEventActionTogglePreHistoryHidden{
			m.Object(),
		}

	case crcChannelAdminLogEventActionDefaultBannedRights:
		r = ChannelAdminLogEventActionDefaultBannedRights{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionStopPoll:
		r = ChannelAdminLogEventActionStopPoll{
			m.Object(),
		}

	case crcChannelAdminLogEventActionChangeLinkedChat:
		r = ChannelAdminLogEventActionChangeLinkedChat{
			m.Int(),
			m.Int(),
		}

	case crcChannelAdminLogEventActionChangeLocation:
		r = ChannelAdminLogEventActionChangeLocation{
			m.Object(),
			m.Object(),
		}

	case crcChannelAdminLogEventActionToggleSlowMode:
		r = ChannelAdminLogEventActionToggleSlowMode{
			m.Int(),
			m.Int(),
		}

	case crcChannelAdminLogEvent:
		r = ChannelAdminLogEvent{
			m.Long(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crcChannelsAdminLogResults:
		r = ChannelsAdminLogResults{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcChannelAdminLogEventsFilter:
		var flags int32
		r = ChannelAdminLogEventsFilter{
			readFlags(m, &flags),
			flags&1 != 0,    //flag #0
			flags&2 != 0,    //flag #1
			flags&4 != 0,    //flag #2
			flags&8 != 0,    //flag #3
			flags&16 != 0,   //flag #4
			flags&32 != 0,   //flag #5
			flags&64 != 0,   //flag #6
			flags&128 != 0,  //flag #7
			flags&256 != 0,  //flag #8
			flags&512 != 0,  //flag #9
			flags&1024 != 0, //flag #10
			flags&2048 != 0, //flag #11
			flags&4096 != 0, //flag #12
			flags&8192 != 0, //flag #13
		}

	case crcPopularContact:
		r = PopularContact{
			m.Long(),
			m.Int(),
		}

	case crcMessagesFavedStickersNotModified:
		r = MessagesFavedStickersNotModified{}

	case crcMessagesFavedStickers:
		r = MessagesFavedStickers{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crcRecentMeUrlUnknown:
		r = RecentMeUrlUnknown{
			m.String(),
		}

	case crcRecentMeUrlUser:
		r = RecentMeUrlUser{
			m.String(),
			m.Int(),
		}

	case crcRecentMeUrlChat:
		r = RecentMeUrlChat{
			m.String(),
			m.Int(),
		}

	case crcRecentMeUrlChatInvite:
		r = RecentMeUrlChatInvite{
			m.String(),
			m.Object(),
		}

	case crcRecentMeUrlStickerSet:
		r = RecentMeUrlStickerSet{
			m.String(),
			m.Object(),
		}

	case crcHelpRecentMeUrls:
		r = HelpRecentMeUrls{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcInputSingleMedia:
		var flags int32
		r = InputSingleMedia{
			readFlags(m, &flags),
			m.Object(),
			m.Long(),
			m.String(),
			m.FlaggedVector(flags, 0),
		}

	case crcWebAuthorization:
		r = WebAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crcAccountWebAuthorizations:
		r = AccountWebAuthorizations{
			m.Vector(),
			m.Vector(),
		}

	case crcInputMessageID:
		r = InputMessageID{
			m.Int(),
		}

	case crcInputMessageReplyTo:
		r = InputMessageReplyTo{
			m.Int(),
		}

	case crcInputMessagePinned:
		r = InputMessagePinned{}

	case crcInputDialogPeer:
		r = InputDialogPeer{
			m.Object(),
		}

	case crcInputDialogPeerFolder:
		r = InputDialogPeerFolder{
			m.Int(),
		}

	case crcDialogPeer:
		r = DialogPeer{
			m.Object(),
		}

	case crcDialogPeerFolder:
		r = DialogPeerFolder{
			m.Int(),
		}

	case crcMessagesFoundStickerSetsNotModified:
		r = MessagesFoundStickerSetsNotModified{}

	case crcMessagesFoundStickerSets:
		r = MessagesFoundStickerSets{
			m.Int(),
			m.Vector(),
		}

	case crcFileHash:
		r = FileHash{
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crcInputClientProxy:
		r = InputClientProxy{
			m.String(),
			m.Int(),
		}

	case crcHelpTermsOfServiceUpdateEmpty:
		r = HelpTermsOfServiceUpdateEmpty{
			m.Int(),
		}

	case crcHelpTermsOfServiceUpdate:
		r = HelpTermsOfServiceUpdate{
			m.Int(),
			m.Object(),
		}

	case crcInputSecureFileUploaded:
		r = InputSecureFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crcInputSecureFile:
		r = InputSecureFile{
			m.Long(),
			m.Long(),
		}

	case crcSecureFileEmpty:
		r = SecureFileEmpty{}

	case crcSecureFile:
		r = SecureFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crcSecureData:
		r = SecureData{
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crcSecurePlainPhone:
		r = SecurePlainPhone{
			m.String(),
		}

	case crcSecurePlainEmail:
		r = SecurePlainEmail{
			m.String(),
		}

	case crcSecureValueTypePersonalDetails:
		r = SecureValueTypePersonalDetails{}

	case crcSecureValueTypePassport:
		r = SecureValueTypePassport{}

	case crcSecureValueTypeDriverLicense:
		r = SecureValueTypeDriverLicense{}

	case crcSecureValueTypeIdentityCard:
		r = SecureValueTypeIdentityCard{}

	case crcSecureValueTypeInternalPassport:
		r = SecureValueTypeInternalPassport{}

	case crcSecureValueTypeAddress:
		r = SecureValueTypeAddress{}

	case crcSecureValueTypeUtilityBill:
		r = SecureValueTypeUtilityBill{}

	case crcSecureValueTypeBankStatement:
		r = SecureValueTypeBankStatement{}

	case crcSecureValueTypeRentalAgreement:
		r = SecureValueTypeRentalAgreement{}

	case crcSecureValueTypePassportRegistration:
		r = SecureValueTypePassportRegistration{}

	case crcSecureValueTypeTemporaryRegistration:
		r = SecureValueTypeTemporaryRegistration{}

	case crcSecureValueTypePhone:
		r = SecureValueTypePhone{}

	case crcSecureValueTypeEmail:
		r = SecureValueTypeEmail{}

	case crcSecureValue:
		var flags int32
		r = SecureValue{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
			m.FlaggedVector(flags, 6),
			m.FlaggedVector(flags, 4),
			m.FlaggedObject(flags, 5),
			m.StringBytes(),
		}

	case crcInputSecureValue:
		var flags int32
		r = InputSecureValue{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
			m.FlaggedVector(flags, 6),
			m.FlaggedVector(flags, 4),
			m.FlaggedObject(flags, 5),
		}

	case crcSecureValueHash:
		r = SecureValueHash{
			m.Object(),
			m.StringBytes(),
		}

	case crcSecureValueErrorData:
		r = SecureValueErrorData{
			m.Object(),
			m.StringBytes(),
			m.String(),
			m.String(),
		}

	case crcSecureValueErrorFrontSide:
		r = SecureValueErrorFrontSide{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crcSecureValueErrorReverseSide:
		r = SecureValueErrorReverseSide{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crcSecureValueErrorSelfie:
		r = SecureValueErrorSelfie{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crcSecureValueErrorFile:
		r = SecureValueErrorFile{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crcSecureValueErrorFiles:
		r = SecureValueErrorFiles{
			m.Object(),
			m.Vector(),
			m.String(),
		}

	case crcSecureValueError:
		r = SecureValueError{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crcSecureValueErrorTranslationFile:
		r = SecureValueErrorTranslationFile{
			m.Object(),
			m.StringBytes(),
			m.String(),
		}

	case crcSecureValueErrorTranslationFiles:
		r = SecureValueErrorTranslationFiles{
			m.Object(),
			m.Vector(),
			m.String(),
		}

	case crcSecureCredentialsEncrypted:
		r = SecureCredentialsEncrypted{
			m.StringBytes(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crcAccountAuthorizationForm:
		var flags int32
		r = AccountAuthorizationForm{
			readFlags(m, &flags),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.FlaggedString(flags, 0),
		}

	case crcAccountSentEmailCode:
		r = AccountSentEmailCode{
			m.String(),
			m.Int(),
		}

	case crcHelpDeepLinkInfoEmpty:
		r = HelpDeepLinkInfoEmpty{}

	case crcHelpDeepLinkInfo:
		var flags int32
		r = HelpDeepLinkInfo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.FlaggedVector(flags, 1),
		}

	case crcSavedPhoneContact:
		r = SavedPhoneContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crcAccountTakeout:
		r = AccountTakeout{
			m.Long(),
		}

	case crcPasswordKdfAlgoUnknown:
		r = PasswordKdfAlgoUnknown{}

	case crcPasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow:
		r = PasswordKdfAlgoSHA256SHA256PBKDF2HMACSHA512iter100000SHA256ModPow{
			m.StringBytes(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crcSecurePasswordKdfAlgoUnknown:
		r = SecurePasswordKdfAlgoUnknown{}

	case crcSecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000:
		r = SecurePasswordKdfAlgoPBKDF2HMACSHA512iter100000{
			m.StringBytes(),
		}

	case crcSecurePasswordKdfAlgoSHA512:
		r = SecurePasswordKdfAlgoSHA512{
			m.StringBytes(),
		}

	case crcSecureSecretSettings:
		r = SecureSecretSettings{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crcInputCheckPasswordEmpty:
		r = InputCheckPasswordEmpty{}

	case crcInputCheckPasswordSRP:
		r = InputCheckPasswordSRP{
			m.Long(),
			m.StringBytes(),
			m.StringBytes(),
		}

	case crcSecureRequiredType:
		var flags int32
		r = SecureRequiredType{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Object(),
		}

	case crcSecureRequiredTypeOneOf:
		r = SecureRequiredTypeOneOf{
			m.Vector(),
		}

	case crcHelpPassportConfigNotModified:
		r = HelpPassportConfigNotModified{}

	case crcHelpPassportConfig:
		r = HelpPassportConfig{
			m.Int(),
			m.Object(),
		}

	case crcInputAppEvent:
		r = InputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.Object(),
		}

	case crcJsonObjectValue:
		r = JsonObjectValue{
			m.String(),
			m.Object(),
		}

	case crcJsonNull:
		r = JsonNull{}

	case crcJsonBool:
		r = JsonBool{
			m.Object(),
		}

	case crcJsonNumber:
		r = JsonNumber{
			m.Double(),
		}

	case crcJsonString:
		r = JsonString{
			m.String(),
		}

	case crcJsonArray:
		r = JsonArray{
			m.Vector(),
		}

	case crcJsonObject:
		r = JsonObject{
			m.Vector(),
		}

	case crcPageTableCell:
		var flags int32
		r = PageTableCell{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			flags&32 != 0, //flag #5
			flags&64 != 0, //flag #6
			m.FlaggedObject(flags, 7),
			m.FlaggedInt(flags, 1),
			m.FlaggedInt(flags, 2),
		}

	case crcPageTableRow:
		r = PageTableRow{
			m.Vector(),
		}

	case crcPageCaption:
		r = PageCaption{
			m.Object(),
			m.Object(),
		}

	case crcPageListItemText:
		r = PageListItemText{
			m.Object(),
		}

	case crcPageListItemBlocks:
		r = PageListItemBlocks{
			m.Vector(),
		}

	case crcPageListOrderedItemText:
		r = PageListOrderedItemText{
			m.String(),
			m.Object(),
		}

	case crcPageListOrderedItemBlocks:
		r = PageListOrderedItemBlocks{
			m.String(),
			m.Vector(),
		}

	case crcPageRelatedArticle:
		var flags int32
		r = PageRelatedArticle{
			readFlags(m, &flags),
			m.String(),
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedLong(flags, 2),
			m.FlaggedString(flags, 3),
			m.FlaggedInt(flags, 4),
		}

	case crcPage:
		var flags int32
		r = Page{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.String(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.FlaggedInt(flags, 3),
		}

	case crcHelpSupportName:
		r = HelpSupportName{
			m.String(),
		}

	case crcHelpUserInfoEmpty:
		r = HelpUserInfoEmpty{}

	case crcHelpUserInfo:
		r = HelpUserInfo{
			m.String(),
			m.Vector(),
			m.String(),
			m.Int(),
		}

	case crcPollAnswer:
		r = PollAnswer{
			m.String(),
			m.StringBytes(),
		}

	case crcPoll:
		var flags int32
		r = Poll{
			m.Long(),
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.String(),
			m.Vector(),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 5),
		}

	case crcPollAnswerVoters:
		var flags int32
		r = PollAnswerVoters{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.StringBytes(),
			m.Int(),
		}

	case crcPollResults:
		var flags int32
		r = PollResults{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedVector(flags, 1),
			m.FlaggedInt(flags, 2),
			m.FlaggedVectorInt(flags, 3),
			m.FlaggedString(flags, 4),
			m.FlaggedVector(flags, 4),
		}

	case crcChatOnlines:
		r = ChatOnlines{
			m.Int(),
		}

	case crcStatsURL:
		r = StatsURL{
			m.String(),
		}

	case crcChatAdminRights:
		var flags int32
		r = ChatAdminRights{
			readFlags(m, &flags),
			flags&1 != 0,   //flag #0
			flags&2 != 0,   //flag #1
			flags&4 != 0,   //flag #2
			flags&8 != 0,   //flag #3
			flags&16 != 0,  //flag #4
			flags&32 != 0,  //flag #5
			flags&128 != 0, //flag #7
			flags&512 != 0, //flag #9
		}

	case crcChatBannedRights:
		var flags int32
		r = ChatBannedRights{
			readFlags(m, &flags),
			flags&1 != 0,      //flag #0
			flags&2 != 0,      //flag #1
			flags&4 != 0,      //flag #2
			flags&8 != 0,      //flag #3
			flags&16 != 0,     //flag #4
			flags&32 != 0,     //flag #5
			flags&64 != 0,     //flag #6
			flags&128 != 0,    //flag #7
			flags&256 != 0,    //flag #8
			flags&1024 != 0,   //flag #10
			flags&32768 != 0,  //flag #15
			flags&131072 != 0, //flag #17
			m.Int(),
		}

	case crcInputWallPaper:
		r = InputWallPaper{
			m.Long(),
			m.Long(),
		}

	case crcInputWallPaperSlug:
		r = InputWallPaperSlug{
			m.String(),
		}

	case crcInputWallPaperNoFile:
		r = InputWallPaperNoFile{}

	case crcAccountWallPapersNotModified:
		r = AccountWallPapersNotModified{}

	case crcAccountWallPapers:
		r = AccountWallPapers{
			m.Int(),
			m.Vector(),
		}

	case crcCodeSettings:
		var flags int32
		r = CodeSettings{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&16 != 0, //flag #4
		}

	case crcWallPaperSettings:
		var flags int32
		r = WallPaperSettings{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 4),
			m.FlaggedInt(flags, 3),
			m.FlaggedInt(flags, 4),
		}

	case crcAutoDownloadSettings:
		var flags int32
		r = AutoDownloadSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			flags&8 != 0, //flag #3
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcAccountAutoDownloadSettings:
		r = AccountAutoDownloadSettings{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crcEmojiKeyword:
		r = EmojiKeyword{
			m.String(),
			m.VectorString(),
		}

	case crcEmojiKeywordDeleted:
		r = EmojiKeywordDeleted{
			m.String(),
			m.VectorString(),
		}

	case crcEmojiKeywordsDifference:
		r = EmojiKeywordsDifference{
			m.String(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crcEmojiURL:
		r = EmojiURL{
			m.String(),
		}

	case crcEmojiLanguage:
		r = EmojiLanguage{
			m.String(),
		}

	case crcFileLocationToBeDeprecated:
		r = FileLocationToBeDeprecated{
			m.Long(),
			m.Int(),
		}

	case crcFolder:
		var flags int32
		r = Folder{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Int(),
			m.String(),
			m.FlaggedObject(flags, 3),
		}

	case crcInputFolderPeer:
		r = InputFolderPeer{
			m.Object(),
			m.Int(),
		}

	case crcFolderPeer:
		r = FolderPeer{
			m.Object(),
			m.Int(),
		}

	case crcMessagesSearchCounter:
		var flags int32
		r = MessagesSearchCounter{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
		}

	case crcUrlAuthResultRequest:
		var flags int32
		r = UrlAuthResultRequest{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
		}

	case crcUrlAuthResultAccepted:
		r = UrlAuthResultAccepted{
			m.String(),
		}

	case crcUrlAuthResultDefault:
		r = UrlAuthResultDefault{}

	case crcChannelLocationEmpty:
		r = ChannelLocationEmpty{}

	case crcChannelLocation:
		r = ChannelLocation{
			m.Object(),
			m.String(),
		}

	case crcPeerLocated:
		r = PeerLocated{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcPeerSelfLocated:
		r = PeerSelfLocated{
			m.Int(),
		}

	case crcRestrictionReason:
		r = RestrictionReason{
			m.String(),
			m.String(),
			m.String(),
		}

	case crcInputTheme:
		r = InputTheme{
			m.Long(),
			m.Long(),
		}

	case crcInputThemeSlug:
		r = InputThemeSlug{
			m.String(),
		}

	case crcTheme:
		var flags int32
		r = Theme{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Long(),
			m.Long(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
			m.Int(),
		}

	case crcAccountThemesNotModified:
		r = AccountThemesNotModified{}

	case crcAccountThemes:
		r = AccountThemes{
			m.Int(),
			m.Vector(),
		}

	case crcAuthLoginToken:
		r = AuthLoginToken{
			m.Int(),
			m.StringBytes(),
		}

	case crcAuthLoginTokenMigrateTo:
		r = AuthLoginTokenMigrateTo{
			m.Int(),
			m.StringBytes(),
		}

	case crcAuthLoginTokenSuccess:
		r = AuthLoginTokenSuccess{
			m.Object(),
		}

	case crcAccountContentSettings:
		var flags int32
		r = AccountContentSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
		}

	case crcMessagesInactiveChats:
		r = MessagesInactiveChats{
			m.VectorInt(),
			m.Vector(),
			m.Vector(),
		}

	case crcBaseThemeClassic:
		r = BaseThemeClassic{}

	case crcBaseThemeDay:
		r = BaseThemeDay{}

	case crcBaseThemeNight:
		r = BaseThemeNight{}

	case crcBaseThemeTinted:
		r = BaseThemeTinted{}

	case crcBaseThemeArctic:
		r = BaseThemeArctic{}

	case crcInputThemeSettings:
		var flags int32
		r = InputThemeSettings{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedObject(flags, 1),
		}

	case crcThemeSettings:
		var flags int32
		r = ThemeSettings{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.FlaggedInt(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case crcWebPageAttributeTheme:
		var flags int32
		r = WebPageAttributeTheme{
			readFlags(m, &flags),
			m.FlaggedVector(flags, 0),
			m.FlaggedObject(flags, 1),
		}

	case crcMessageUserVote:
		r = MessageUserVote{
			m.Int(),
			m.StringBytes(),
			m.Int(),
		}

	case crcMessageUserVoteInputOption:
		r = MessageUserVoteInputOption{
			m.Int(),
			m.Int(),
		}

	case crcMessageUserVoteMultiple:
		r = MessageUserVoteMultiple{
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crcMessagesVotesList:
		var flags int32
		r = MessagesVotesList{
			readFlags(m, &flags),
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.FlaggedString(flags, 0),
		}

	case crcBankCardOpenUrl:
		r = BankCardOpenUrl{
			m.String(),
			m.String(),
		}

	case crcPaymentsBankCardData:
		r = PaymentsBankCardData{
			m.String(),
			m.Vector(),
		}

	case crcDialogFilter:
		var flags int32
		r = DialogFilter{
			readFlags(m, &flags),
			flags&1 != 0,    //flag #0
			flags&2 != 0,    //flag #1
			flags&4 != 0,    //flag #2
			flags&8 != 0,    //flag #3
			flags&16 != 0,   //flag #4
			flags&2048 != 0, //flag #11
			flags&4096 != 0, //flag #12
			flags&8192 != 0, //flag #13
			m.Int(),
			m.String(),
			m.FlaggedString(flags, 25),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcDialogFilterSuggested:
		r = DialogFilterSuggested{
			m.Object(),
			m.String(),
		}

	case crcStatsDateRangeDays:
		r = StatsDateRangeDays{
			m.Int(),
			m.Int(),
		}

	case crcStatsAbsValueAndPrev:
		r = StatsAbsValueAndPrev{
			m.Double(),
			m.Double(),
		}

	case crcStatsPercentValue:
		r = StatsPercentValue{
			m.Double(),
			m.Double(),
		}

	case crcStatsGraphAsync:
		r = StatsGraphAsync{
			m.String(),
		}

	case crcStatsGraphError:
		r = StatsGraphError{
			m.String(),
		}

	case crcStatsGraph:
		var flags int32
		r = StatsGraph{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedString(flags, 0),
		}

	case crcMessageInteractionCounters:
		r = MessageInteractionCounters{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcStatsBroadcastStats:
		r = StatsBroadcastStats{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
		}

	case crcHelpPromoDataEmpty:
		r = HelpPromoDataEmpty{
			m.Int(),
		}

	case crcHelpPromoData:
		var flags int32
		r = HelpPromoData{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case crcVideoSize:
		var flags int32
		r = VideoSize{
			readFlags(m, &flags),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.FlaggedDouble(flags, 0),
		}

	case crcStatsGroupTopPoster:
		r = StatsGroupTopPoster{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcStatsGroupTopAdmin:
		r = StatsGroupTopAdmin{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcStatsGroupTopInviter:
		r = StatsGroupTopInviter{
			m.Int(),
			m.Int(),
		}

	case crcStatsMegagroupStats:
		r = StatsMegagroupStats{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crcGlobalPrivacySettings:
		var flags int32
		r = GlobalPrivacySettings{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
		}

	case crcInvokeAfterMsg:
		r = InvokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case crcInvokeAfterMsgs:
		r = InvokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case crcInitConnection:
		var flags int32
		r = InitConnection{
			readFlags(m, &flags),
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.Object(),
		}

	case crcInvokeWithLayer:
		r = InvokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case crcInvokeWithoutUpdates:
		r = InvokeWithoutUpdates{
			m.Object(),
		}

	case crcInvokeWithMessagesRange:
		r = InvokeWithMessagesRange{
			m.Object(),
			m.Object(),
		}

	case crcInvokeWithTakeout:
		r = InvokeWithTakeout{
			m.Long(),
			m.Object(),
		}

	case crcAuthSendCode:
		r = AuthSendCode{
			m.String(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case crcAuthSignUp:
		r = AuthSignUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcAuthSignIn:
		r = AuthSignIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crcAuthLogOut:
		r = AuthLogOut{}

	case crcAuthResetAuthorizations:
		r = AuthResetAuthorizations{}

	case crcAuthExportAuthorization:
		r = AuthExportAuthorization{
			m.Int(),
		}

	case crcAuthImportAuthorization:
		r = AuthImportAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crcAuthBindTempAuthKey:
		r = AuthBindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crcAuthImportBotAuthorization:
		r = AuthImportBotAuthorization{
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crcAuthCheckPassword:
		r = AuthCheckPassword{
			m.Object(),
		}

	case crcAuthRequestPasswordRecovery:
		r = AuthRequestPasswordRecovery{}

	case crcAuthRecoverPassword:
		r = AuthRecoverPassword{
			m.String(),
		}

	case crcAuthResendCode:
		r = AuthResendCode{
			m.String(),
			m.String(),
		}

	case crcAuthCancelCode:
		r = AuthCancelCode{
			m.String(),
			m.String(),
		}

	case crcAuthDropTempAuthKeys:
		r = AuthDropTempAuthKeys{
			m.VectorLong(),
		}

	case crcAuthExportLoginToken:
		r = AuthExportLoginToken{
			m.Int(),
			m.String(),
			m.VectorInt(),
		}

	case crcAuthImportLoginToken:
		r = AuthImportLoginToken{
			m.StringBytes(),
		}

	case crcAuthAcceptLoginToken:
		r = AuthAcceptLoginToken{
			m.StringBytes(),
		}

	case crcAccountRegisterDevice:
		var flags int32
		r = AccountRegisterDevice{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.String(),
			m.Object(),
			m.StringBytes(),
			m.VectorInt(),
		}

	case crcAccountUnregisterDevice:
		r = AccountUnregisterDevice{
			m.Int(),
			m.String(),
			m.VectorInt(),
		}

	case crcAccountUpdateNotifySettings:
		r = AccountUpdateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crcAccountGetNotifySettings:
		r = AccountGetNotifySettings{
			m.Object(),
		}

	case crcAccountResetNotifySettings:
		r = AccountResetNotifySettings{}

	case crcAccountUpdateProfile:
		var flags int32
		r = AccountUpdateProfile{
			readFlags(m, &flags),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedString(flags, 2),
		}

	case crcAccountUpdateStatus:
		r = AccountUpdateStatus{
			m.Object(),
		}

	case crcAccountGetWallPapers:
		r = AccountGetWallPapers{
			m.Int(),
		}

	case crcAccountReportPeer:
		r = AccountReportPeer{
			m.Object(),
			m.Object(),
		}

	case crcAccountCheckUsername:
		r = AccountCheckUsername{
			m.String(),
		}

	case crcAccountUpdateUsername:
		r = AccountUpdateUsername{
			m.String(),
		}

	case crcAccountGetPrivacy:
		r = AccountGetPrivacy{
			m.Object(),
		}

	case crcAccountSetPrivacy:
		r = AccountSetPrivacy{
			m.Object(),
			m.Vector(),
		}

	case crcAccountDeleteAccount:
		r = AccountDeleteAccount{
			m.String(),
		}

	case crcAccountGetAccountTTL:
		r = AccountGetAccountTTL{}

	case crcAccountSetAccountTTL:
		r = AccountSetAccountTTL{
			m.Object(),
		}

	case crcAccountSendChangePhoneCode:
		r = AccountSendChangePhoneCode{
			m.String(),
			m.Object(),
		}

	case crcAccountChangePhone:
		r = AccountChangePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crcAccountUpdateDeviceLocked:
		r = AccountUpdateDeviceLocked{
			m.Int(),
		}

	case crcAccountGetAuthorizations:
		r = AccountGetAuthorizations{}

	case crcAccountResetAuthorization:
		r = AccountResetAuthorization{
			m.Long(),
		}

	case crcAccountGetPassword:
		r = AccountGetPassword{}

	case crcAccountGetPasswordSettings:
		r = AccountGetPasswordSettings{
			m.Object(),
		}

	case crcAccountUpdatePasswordSettings:
		r = AccountUpdatePasswordSettings{
			m.Object(),
			m.Object(),
		}

	case crcAccountSendConfirmPhoneCode:
		r = AccountSendConfirmPhoneCode{
			m.String(),
			m.Object(),
		}

	case crcAccountConfirmPhone:
		r = AccountConfirmPhone{
			m.String(),
			m.String(),
		}

	case crcAccountGetTmpPassword:
		r = AccountGetTmpPassword{
			m.Object(),
			m.Int(),
		}

	case crcAccountGetWebAuthorizations:
		r = AccountGetWebAuthorizations{}

	case crcAccountResetWebAuthorization:
		r = AccountResetWebAuthorization{
			m.Long(),
		}

	case crcAccountResetWebAuthorizations:
		r = AccountResetWebAuthorizations{}

	case crcAccountGetAllSecureValues:
		r = AccountGetAllSecureValues{}

	case crcAccountGetSecureValue:
		r = AccountGetSecureValue{
			m.Vector(),
		}

	case crcAccountSaveSecureValue:
		r = AccountSaveSecureValue{
			m.Object(),
			m.Long(),
		}

	case crcAccountDeleteSecureValue:
		r = AccountDeleteSecureValue{
			m.Vector(),
		}

	case crcAccountGetAuthorizationForm:
		r = AccountGetAuthorizationForm{
			m.Int(),
			m.String(),
			m.String(),
		}

	case crcAccountAcceptAuthorization:
		r = AccountAcceptAuthorization{
			m.Int(),
			m.String(),
			m.String(),
			m.Vector(),
			m.Object(),
		}

	case crcAccountSendVerifyPhoneCode:
		r = AccountSendVerifyPhoneCode{
			m.String(),
			m.Object(),
		}

	case crcAccountVerifyPhone:
		r = AccountVerifyPhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crcAccountSendVerifyEmailCode:
		r = AccountSendVerifyEmailCode{
			m.String(),
		}

	case crcAccountVerifyEmail:
		r = AccountVerifyEmail{
			m.String(),
			m.String(),
		}

	case crcAccountInitTakeoutSession:
		var flags int32
		r = AccountInitTakeoutSession{
			readFlags(m, &flags),
			flags&1 != 0,  //flag #0
			flags&2 != 0,  //flag #1
			flags&4 != 0,  //flag #2
			flags&8 != 0,  //flag #3
			flags&16 != 0, //flag #4
			flags&32 != 0, //flag #5
			m.FlaggedInt(flags, 5),
		}

	case crcAccountFinishTakeoutSession:
		var flags int32
		r = AccountFinishTakeoutSession{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case crcAccountConfirmPasswordEmail:
		r = AccountConfirmPasswordEmail{
			m.String(),
		}

	case crcAccountResendPasswordEmail:
		r = AccountResendPasswordEmail{}

	case crcAccountCancelPasswordEmail:
		r = AccountCancelPasswordEmail{}

	case crcAccountGetContactSignUpNotification:
		r = AccountGetContactSignUpNotification{}

	case crcAccountSetContactSignUpNotification:
		r = AccountSetContactSignUpNotification{
			m.Object(),
		}

	case crcAccountGetNotifyExceptions:
		var flags int32
		r = AccountGetNotifyExceptions{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedObject(flags, 0),
		}

	case crcAccountGetWallPaper:
		r = AccountGetWallPaper{
			m.Object(),
		}

	case crcAccountUploadWallPaper:
		r = AccountUploadWallPaper{
			m.Object(),
			m.String(),
			m.Object(),
		}

	case crcAccountSaveWallPaper:
		r = AccountSaveWallPaper{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crcAccountInstallWallPaper:
		r = AccountInstallWallPaper{
			m.Object(),
			m.Object(),
		}

	case crcAccountResetWallPapers:
		r = AccountResetWallPapers{}

	case crcAccountGetAutoDownloadSettings:
		r = AccountGetAutoDownloadSettings{}

	case crcAccountSaveAutoDownloadSettings:
		var flags int32
		r = AccountSaveAutoDownloadSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
		}

	case crcAccountUploadTheme:
		var flags int32
		r = AccountUploadTheme{
			readFlags(m, &flags),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.String(),
		}

	case crcAccountCreateTheme:
		var flags int32
		r = AccountCreateTheme{
			readFlags(m, &flags),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crcAccountUpdateTheme:
		var flags int32
		r = AccountUpdateTheme{
			readFlags(m, &flags),
			m.String(),
			m.Object(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crcAccountSaveTheme:
		r = AccountSaveTheme{
			m.Object(),
			m.Object(),
		}

	case crcAccountInstallTheme:
		var flags int32
		r = AccountInstallTheme{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedString(flags, 1),
			m.FlaggedObject(flags, 1),
		}

	case crcAccountGetTheme:
		r = AccountGetTheme{
			m.String(),
			m.Object(),
			m.Long(),
		}

	case crcAccountGetThemes:
		r = AccountGetThemes{
			m.String(),
			m.Int(),
		}

	case crcAccountSetContentSettings:
		var flags int32
		r = AccountSetContentSettings{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case crcAccountGetContentSettings:
		r = AccountGetContentSettings{}

	case crcAccountGetMultiWallPapers:
		r = AccountGetMultiWallPapers{
			m.Vector(),
		}

	case crcAccountGetGlobalPrivacySettings:
		r = AccountGetGlobalPrivacySettings{}

	case crcAccountSetGlobalPrivacySettings:
		r = AccountSetGlobalPrivacySettings{
			m.Object(),
		}

	case crcUsersGetUsers:
		r = UsersGetUsers{
			m.Vector(),
		}

	case crcUsersGetFullUser:
		r = UsersGetFullUser{
			m.Object(),
		}

	case crcUsersSetSecureValueErrors:
		r = UsersSetSecureValueErrors{
			m.Object(),
			m.Vector(),
		}

	case crcContactsGetContactIDs:
		r = ContactsGetContactIDs{
			m.Int(),
		}

	case crcContactsGetStatuses:
		r = ContactsGetStatuses{}

	case crcContactsGetContacts:
		r = ContactsGetContacts{
			m.Int(),
		}

	case crcContactsImportContacts:
		r = ContactsImportContacts{
			m.Vector(),
		}

	case crcContactsDeleteContacts:
		r = ContactsDeleteContacts{
			m.Vector(),
		}

	case crcContactsDeleteByPhones:
		r = ContactsDeleteByPhones{
			m.VectorString(),
		}

	case crcContactsBlock:
		r = ContactsBlock{
			m.Object(),
		}

	case crcContactsUnblock:
		r = ContactsUnblock{
			m.Object(),
		}

	case crcContactsGetBlocked:
		r = ContactsGetBlocked{
			m.Int(),
			m.Int(),
		}

	case crcContactsSearch:
		r = ContactsSearch{
			m.String(),
			m.Int(),
		}

	case crcContactsResolveUsername:
		r = ContactsResolveUsername{
			m.String(),
		}

	case crcContactsGetTopPeers:
		var flags int32
		r = ContactsGetTopPeers{
			readFlags(m, &flags),
			flags&1 != 0,     //flag #0
			flags&2 != 0,     //flag #1
			flags&4 != 0,     //flag #2
			flags&8 != 0,     //flag #3
			flags&16 != 0,    //flag #4
			flags&32 != 0,    //flag #5
			flags&1024 != 0,  //flag #10
			flags&32768 != 0, //flag #15
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcContactsResetTopPeerRating:
		r = ContactsResetTopPeerRating{
			m.Object(),
			m.Object(),
		}

	case crcContactsResetSaved:
		r = ContactsResetSaved{}

	case crcContactsGetSaved:
		r = ContactsGetSaved{}

	case crcContactsToggleTopPeers:
		r = ContactsToggleTopPeers{
			m.Object(),
		}

	case crcContactsAddContact:
		var flags int32
		r = ContactsAddContact{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crcContactsAcceptContact:
		r = ContactsAcceptContact{
			m.Object(),
		}

	case crcContactsGetLocated:
		var flags int32
		r = ContactsGetLocated{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.FlaggedInt(flags, 0),
		}

	case crcMessagesGetMessages:
		r = MessagesGetMessages{
			m.Vector(),
		}

	case crcMessagesGetDialogs:
		var flags int32
		r = MessagesGetDialogs{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.FlaggedInt(flags, 1),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesGetHistory:
		r = MessagesGetHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesSearch:
		var flags int32
		r = MessagesSearch{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesReadHistory:
		r = MessagesReadHistory{
			m.Object(),
			m.Int(),
		}

	case crcMessagesDeleteHistory:
		var flags int32
		r = MessagesDeleteHistory{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
		}

	case crcMessagesDeleteMessages:
		var flags int32
		r = MessagesDeleteMessages{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.VectorInt(),
		}

	case crcMessagesReceivedMessages:
		r = MessagesReceivedMessages{
			m.Int(),
		}

	case crcMessagesSetTyping:
		r = MessagesSetTyping{
			m.Object(),
			m.Object(),
		}

	case crcMessagesSendMessage:
		var flags int32
		r = MessagesSendMessage{
			readFlags(m, &flags),
			flags&2 != 0,   //flag #1
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Long(),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 10),
		}

	case crcMessagesSendMedia:
		var flags int32
		r = MessagesSendMedia{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.String(),
			m.Long(),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 10),
		}

	case crcMessagesForwardMessages:
		var flags int32
		r = MessagesForwardMessages{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&256 != 0, //flag #8
			flags&512 != 0, //flag #9
			m.Object(),
			m.VectorInt(),
			m.VectorLong(),
			m.Object(),
			m.FlaggedInt(flags, 10),
		}

	case crcMessagesReportSpam:
		r = MessagesReportSpam{
			m.Object(),
		}

	case crcMessagesGetPeerSettings:
		r = MessagesGetPeerSettings{
			m.Object(),
		}

	case crcMessagesReport:
		r = MessagesReport{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crcMessagesGetChats:
		r = MessagesGetChats{
			m.VectorInt(),
		}

	case crcMessagesGetFullChat:
		r = MessagesGetFullChat{
			m.Int(),
		}

	case crcMessagesEditChatTitle:
		r = MessagesEditChatTitle{
			m.Int(),
			m.String(),
		}

	case crcMessagesEditChatPhoto:
		r = MessagesEditChatPhoto{
			m.Int(),
			m.Object(),
		}

	case crcMessagesAddChatUser:
		r = MessagesAddChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crcMessagesDeleteChatUser:
		r = MessagesDeleteChatUser{
			m.Int(),
			m.Object(),
		}

	case crcMessagesCreateChat:
		r = MessagesCreateChat{
			m.Vector(),
			m.String(),
		}

	case crcMessagesGetDhConfig:
		r = MessagesGetDhConfig{
			m.Int(),
			m.Int(),
		}

	case crcMessagesRequestEncryption:
		r = MessagesRequestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crcMessagesAcceptEncryption:
		r = MessagesAcceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crcMessagesDiscardEncryption:
		r = MessagesDiscardEncryption{
			m.Int(),
		}

	case crcMessagesSetEncryptedTyping:
		r = MessagesSetEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case crcMessagesReadEncryptedHistory:
		r = MessagesReadEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case crcMessagesSendEncrypted:
		r = MessagesSendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crcMessagesSendEncryptedFile:
		r = MessagesSendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case crcMessagesSendEncryptedService:
		r = MessagesSendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crcMessagesReceivedQueue:
		r = MessagesReceivedQueue{
			m.Int(),
		}

	case crcMessagesReportEncryptedSpam:
		r = MessagesReportEncryptedSpam{
			m.Object(),
		}

	case crcMessagesReadMessageContents:
		r = MessagesReadMessageContents{
			m.VectorInt(),
		}

	case crcMessagesGetStickers:
		r = MessagesGetStickers{
			m.String(),
			m.Int(),
		}

	case crcMessagesGetAllStickers:
		r = MessagesGetAllStickers{
			m.Int(),
		}

	case crcMessagesGetWebPagePreview:
		var flags int32
		r = MessagesGetWebPagePreview{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedVector(flags, 3),
		}

	case crcMessagesExportChatInvite:
		r = MessagesExportChatInvite{
			m.Object(),
		}

	case crcMessagesCheckChatInvite:
		r = MessagesCheckChatInvite{
			m.String(),
		}

	case crcMessagesImportChatInvite:
		r = MessagesImportChatInvite{
			m.String(),
		}

	case crcMessagesGetStickerSet:
		r = MessagesGetStickerSet{
			m.Object(),
		}

	case crcMessagesInstallStickerSet:
		r = MessagesInstallStickerSet{
			m.Object(),
			m.Object(),
		}

	case crcMessagesUninstallStickerSet:
		r = MessagesUninstallStickerSet{
			m.Object(),
		}

	case crcMessagesStartBot:
		r = MessagesStartBot{
			m.Object(),
			m.Object(),
			m.Long(),
			m.String(),
		}

	case crcMessagesGetMessagesViews:
		r = MessagesGetMessagesViews{
			m.Object(),
			m.VectorInt(),
			m.Object(),
		}

	case crcMessagesEditChatAdmin:
		r = MessagesEditChatAdmin{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crcMessagesMigrateChat:
		r = MessagesMigrateChat{
			m.Int(),
		}

	case crcMessagesSearchGlobal:
		var flags int32
		r = MessagesSearchGlobal{
			readFlags(m, &flags),
			m.FlaggedInt(flags, 0),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesReorderStickerSets:
		var flags int32
		r = MessagesReorderStickerSets{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.VectorLong(),
		}

	case crcMessagesGetDocumentByHash:
		r = MessagesGetDocumentByHash{
			m.StringBytes(),
			m.Int(),
			m.String(),
		}

	case crcMessagesGetSavedGifs:
		r = MessagesGetSavedGifs{
			m.Int(),
		}

	case crcMessagesSaveGif:
		r = MessagesSaveGif{
			m.Object(),
			m.Object(),
		}

	case crcMessagesGetInlineBotResults:
		var flags int32
		r = MessagesGetInlineBotResults{
			readFlags(m, &flags),
			m.Object(),
			m.Object(),
			m.FlaggedObject(flags, 0),
			m.String(),
			m.String(),
		}

	case crcMessagesSetInlineBotResults:
		var flags int32
		r = MessagesSetInlineBotResults{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Long(),
			m.Vector(),
			m.Int(),
			m.FlaggedString(flags, 2),
			m.FlaggedObject(flags, 3),
		}

	case crcMessagesSendInlineBotResult:
		var flags int32
		r = MessagesSendInlineBotResult{
			readFlags(m, &flags),
			flags&32 != 0,   //flag #5
			flags&64 != 0,   //flag #6
			flags&128 != 0,  //flag #7
			flags&2048 != 0, //flag #11
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Long(),
			m.Long(),
			m.String(),
			m.FlaggedInt(flags, 10),
		}

	case crcMessagesGetMessageEditData:
		r = MessagesGetMessageEditData{
			m.Object(),
			m.Int(),
		}

	case crcMessagesEditMessage:
		var flags int32
		r = MessagesEditMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
			m.FlaggedInt(flags, 15),
		}

	case crcMessagesEditInlineBotMessage:
		var flags int32
		r = MessagesEditInlineBotMessage{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.FlaggedString(flags, 11),
			m.FlaggedObject(flags, 14),
			m.FlaggedObject(flags, 2),
			m.FlaggedVector(flags, 3),
		}

	case crcMessagesGetBotCallbackAnswer:
		var flags int32
		r = MessagesGetBotCallbackAnswer{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
		}

	case crcMessagesSetBotCallbackAnswer:
		var flags int32
		r = MessagesSetBotCallbackAnswer{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 2),
			m.Int(),
		}

	case crcMessagesGetPeerDialogs:
		r = MessagesGetPeerDialogs{
			m.Vector(),
		}

	case crcMessagesSaveDraft:
		var flags int32
		r = MessagesSaveDraft{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.FlaggedInt(flags, 0),
			m.Object(),
			m.String(),
			m.FlaggedVector(flags, 3),
		}

	case crcMessagesGetAllDrafts:
		r = MessagesGetAllDrafts{}

	case crcMessagesGetFeaturedStickers:
		r = MessagesGetFeaturedStickers{
			m.Int(),
		}

	case crcMessagesReadFeaturedStickers:
		r = MessagesReadFeaturedStickers{
			m.VectorLong(),
		}

	case crcMessagesGetRecentStickers:
		var flags int32
		r = MessagesGetRecentStickers{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
		}

	case crcMessagesSaveRecentSticker:
		var flags int32
		r = MessagesSaveRecentSticker{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Object(),
		}

	case crcMessagesClearRecentStickers:
		var flags int32
		r = MessagesClearRecentStickers{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
		}

	case crcMessagesGetArchivedStickers:
		var flags int32
		r = MessagesGetArchivedStickers{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Long(),
			m.Int(),
		}

	case crcMessagesGetMaskStickers:
		r = MessagesGetMaskStickers{
			m.Int(),
		}

	case crcMessagesGetAttachedStickers:
		r = MessagesGetAttachedStickers{
			m.Object(),
		}

	case crcMessagesSetGameScore:
		var flags int32
		r = MessagesSetGameScore{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crcMessagesSetInlineGameScore:
		var flags int32
		r = MessagesSetInlineGameScore{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Object(),
			m.Int(),
		}

	case crcMessagesGetGameHighScores:
		r = MessagesGetGameHighScores{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crcMessagesGetInlineGameHighScores:
		r = MessagesGetInlineGameHighScores{
			m.Object(),
			m.Object(),
		}

	case crcMessagesGetCommonChats:
		r = MessagesGetCommonChats{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesGetAllChats:
		r = MessagesGetAllChats{
			m.VectorInt(),
		}

	case crcMessagesGetWebPage:
		r = MessagesGetWebPage{
			m.String(),
			m.Int(),
		}

	case crcMessagesToggleDialogPin:
		var flags int32
		r = MessagesToggleDialogPin{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case crcMessagesReorderPinnedDialogs:
		var flags int32
		r = MessagesReorderPinnedDialogs{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Vector(),
		}

	case crcMessagesGetPinnedDialogs:
		r = MessagesGetPinnedDialogs{
			m.Int(),
		}

	case crcMessagesSetBotShippingResults:
		var flags int32
		r = MessagesSetBotShippingResults{
			readFlags(m, &flags),
			m.Long(),
			m.FlaggedString(flags, 0),
			m.FlaggedVector(flags, 1),
		}

	case crcMessagesSetBotPrecheckoutResults:
		var flags int32
		r = MessagesSetBotPrecheckoutResults{
			readFlags(m, &flags),
			flags&2 != 0, //flag #1
			m.Long(),
			m.FlaggedString(flags, 0),
		}

	case crcMessagesUploadMedia:
		r = MessagesUploadMedia{
			m.Object(),
			m.Object(),
		}

	case crcMessagesSendScreenshotNotification:
		r = MessagesSendScreenshotNotification{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crcMessagesGetFavedStickers:
		r = MessagesGetFavedStickers{
			m.Int(),
		}

	case crcMessagesFaveSticker:
		r = MessagesFaveSticker{
			m.Object(),
			m.Object(),
		}

	case crcMessagesGetUnreadMentions:
		r = MessagesGetUnreadMentions{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesReadMentions:
		r = MessagesReadMentions{
			m.Object(),
		}

	case crcMessagesGetRecentLocations:
		r = MessagesGetRecentLocations{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesSendMultiMedia:
		var flags int32
		r = MessagesSendMultiMedia{
			readFlags(m, &flags),
			flags&32 != 0,  //flag #5
			flags&64 != 0,  //flag #6
			flags&128 != 0, //flag #7
			m.Object(),
			m.FlaggedInt(flags, 0),
			m.Vector(),
			m.FlaggedInt(flags, 10),
		}

	case crcMessagesUploadEncryptedFile:
		r = MessagesUploadEncryptedFile{
			m.Object(),
			m.Object(),
		}

	case crcMessagesSearchStickerSets:
		var flags int32
		r = MessagesSearchStickerSets{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.String(),
			m.Int(),
		}

	case crcMessagesGetSplitRanges:
		r = MessagesGetSplitRanges{}

	case crcMessagesMarkDialogUnread:
		var flags int32
		r = MessagesMarkDialogUnread{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case crcMessagesGetDialogUnreadMarks:
		r = MessagesGetDialogUnreadMarks{}

	case crcMessagesClearAllDrafts:
		r = MessagesClearAllDrafts{}

	case crcMessagesUpdatePinnedMessage:
		var flags int32
		r = MessagesUpdatePinnedMessage{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
		}

	case crcMessagesSendVote:
		r = MessagesSendVote{
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crcMessagesGetPollResults:
		r = MessagesGetPollResults{
			m.Object(),
			m.Int(),
		}

	case crcMessagesGetOnlines:
		r = MessagesGetOnlines{
			m.Object(),
		}

	case crcMessagesGetStatsURL:
		var flags int32
		r = MessagesGetStatsURL{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.String(),
		}

	case crcMessagesEditChatAbout:
		r = MessagesEditChatAbout{
			m.Object(),
			m.String(),
		}

	case crcMessagesEditChatDefaultBannedRights:
		r = MessagesEditChatDefaultBannedRights{
			m.Object(),
			m.Object(),
		}

	case crcMessagesGetEmojiKeywords:
		r = MessagesGetEmojiKeywords{
			m.String(),
		}

	case crcMessagesGetEmojiKeywordsDifference:
		r = MessagesGetEmojiKeywordsDifference{
			m.String(),
			m.Int(),
		}

	case crcMessagesGetEmojiKeywordsLanguages:
		r = MessagesGetEmojiKeywordsLanguages{
			m.VectorString(),
		}

	case crcMessagesGetEmojiURL:
		r = MessagesGetEmojiURL{
			m.String(),
		}

	case crcMessagesGetSearchCounters:
		r = MessagesGetSearchCounters{
			m.Object(),
			m.Vector(),
		}

	case crcMessagesRequestUrlAuth:
		r = MessagesRequestUrlAuth{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesAcceptUrlAuth:
		var flags int32
		r = MessagesAcceptUrlAuth{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcMessagesHidePeerSettingsBar:
		r = MessagesHidePeerSettingsBar{
			m.Object(),
		}

	case crcMessagesGetScheduledHistory:
		r = MessagesGetScheduledHistory{
			m.Object(),
			m.Int(),
		}

	case crcMessagesGetScheduledMessages:
		r = MessagesGetScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crcMessagesSendScheduledMessages:
		r = MessagesSendScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crcMessagesDeleteScheduledMessages:
		r = MessagesDeleteScheduledMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crcMessagesGetPollVotes:
		var flags int32
		r = MessagesGetPollVotes{
			readFlags(m, &flags),
			m.Object(),
			m.Int(),
			m.FlaggedStringBytes(flags, 0),
			m.FlaggedString(flags, 1),
			m.Int(),
		}

	case crcMessagesToggleStickerSets:
		var flags int32
		r = MessagesToggleStickerSets{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			flags&4 != 0, //flag #2
			m.Vector(),
		}

	case crcMessagesGetDialogFilters:
		r = MessagesGetDialogFilters{}

	case crcMessagesGetSuggestedDialogFilters:
		r = MessagesGetSuggestedDialogFilters{}

	case crcMessagesUpdateDialogFilter:
		var flags int32
		r = MessagesUpdateDialogFilter{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedObject(flags, 0),
		}

	case crcMessagesUpdateDialogFiltersOrder:
		r = MessagesUpdateDialogFiltersOrder{
			m.VectorInt(),
		}

	case crcMessagesGetOldFeaturedStickers:
		r = MessagesGetOldFeaturedStickers{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcUpdatesGetState:
		r = UpdatesGetState{}

	case crcUpdatesGetDifference:
		var flags int32
		r = UpdatesGetDifference{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedInt(flags, 0),
			m.Int(),
			m.Int(),
		}

	case crcUpdatesGetChannelDifference:
		var flags int32
		r = UpdatesGetChannelDifference{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcPhotosUpdateProfilePhoto:
		r = PhotosUpdateProfilePhoto{
			m.Object(),
		}

	case crcPhotosUploadProfilePhoto:
		var flags int32
		r = PhotosUploadProfilePhoto{
			readFlags(m, &flags),
			m.FlaggedObject(flags, 0),
			m.FlaggedObject(flags, 1),
			m.FlaggedDouble(flags, 2),
		}

	case crcPhotosDeletePhotos:
		r = PhotosDeletePhotos{
			m.Vector(),
		}

	case crcPhotosGetUserPhotos:
		r = PhotosGetUserPhotos{
			m.Object(),
			m.Int(),
			m.Long(),
			m.Int(),
		}

	case crcUploadSaveFilePart:
		r = UploadSaveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crcUploadGetFile:
		var flags int32
		r = UploadGetFile{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUploadSaveBigFilePart:
		r = UploadSaveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crcUploadGetWebFile:
		r = UploadGetWebFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crcUploadGetCdnFile:
		r = UploadGetCdnFile{
			m.StringBytes(),
			m.Int(),
			m.Int(),
		}

	case crcUploadReuploadCdnFile:
		r = UploadReuploadCdnFile{
			m.StringBytes(),
			m.StringBytes(),
		}

	case crcUploadGetCdnFileHashes:
		r = UploadGetCdnFileHashes{
			m.StringBytes(),
			m.Int(),
		}

	case crcUploadGetFileHashes:
		r = UploadGetFileHashes{
			m.Object(),
			m.Int(),
		}

	case crcHelpGetConfig:
		r = HelpGetConfig{}

	case crcHelpGetNearestDc:
		r = HelpGetNearestDc{}

	case crcHelpGetAppUpdate:
		r = HelpGetAppUpdate{
			m.String(),
		}

	case crcHelpGetInviteText:
		r = HelpGetInviteText{}

	case crcHelpGetSupport:
		r = HelpGetSupport{}

	case crcHelpGetAppChangelog:
		r = HelpGetAppChangelog{
			m.String(),
		}

	case crcHelpSetBotUpdatesStatus:
		r = HelpSetBotUpdatesStatus{
			m.Int(),
			m.String(),
		}

	case crcHelpGetCdnConfig:
		r = HelpGetCdnConfig{}

	case crcHelpGetRecentMeUrls:
		r = HelpGetRecentMeUrls{
			m.String(),
		}

	case crcHelpGetTermsOfServiceUpdate:
		r = HelpGetTermsOfServiceUpdate{}

	case crcHelpAcceptTermsOfService:
		r = HelpAcceptTermsOfService{
			m.Object(),
		}

	case crcHelpGetDeepLinkInfo:
		r = HelpGetDeepLinkInfo{
			m.String(),
		}

	case crcHelpGetAppConfig:
		r = HelpGetAppConfig{}

	case crcHelpSaveAppLog:
		r = HelpSaveAppLog{
			m.Vector(),
		}

	case crcHelpGetPassportConfig:
		r = HelpGetPassportConfig{
			m.Int(),
		}

	case crcHelpGetSupportName:
		r = HelpGetSupportName{}

	case crcHelpGetUserInfo:
		r = HelpGetUserInfo{
			m.Object(),
		}

	case crcHelpEditUserInfo:
		r = HelpEditUserInfo{
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case crcHelpGetPromoData:
		r = HelpGetPromoData{}

	case crcHelpHidePromoData:
		r = HelpHidePromoData{
			m.Object(),
		}

	case crcHelpDismissSuggestion:
		r = HelpDismissSuggestion{
			m.String(),
		}

	case crcChannelsReadHistory:
		r = ChannelsReadHistory{
			m.Object(),
			m.Int(),
		}

	case crcChannelsDeleteMessages:
		r = ChannelsDeleteMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crcChannelsDeleteUserHistory:
		r = ChannelsDeleteUserHistory{
			m.Object(),
			m.Object(),
		}

	case crcChannelsReportSpam:
		r = ChannelsReportSpam{
			m.Object(),
			m.Object(),
			m.VectorInt(),
		}

	case crcChannelsGetMessages:
		r = ChannelsGetMessages{
			m.Object(),
			m.Vector(),
		}

	case crcChannelsGetParticipants:
		r = ChannelsGetParticipants{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crcChannelsGetParticipant:
		r = ChannelsGetParticipant{
			m.Object(),
			m.Object(),
		}

	case crcChannelsGetChannels:
		r = ChannelsGetChannels{
			m.Vector(),
		}

	case crcChannelsGetFullChannel:
		r = ChannelsGetFullChannel{
			m.Object(),
		}

	case crcChannelsCreateChannel:
		var flags int32
		r = ChannelsCreateChannel{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.FlaggedString(flags, 2),
		}

	case crcChannelsEditAdmin:
		r = ChannelsEditAdmin{
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crcChannelsEditTitle:
		r = ChannelsEditTitle{
			m.Object(),
			m.String(),
		}

	case crcChannelsEditPhoto:
		r = ChannelsEditPhoto{
			m.Object(),
			m.Object(),
		}

	case crcChannelsCheckUsername:
		r = ChannelsCheckUsername{
			m.Object(),
			m.String(),
		}

	case crcChannelsUpdateUsername:
		r = ChannelsUpdateUsername{
			m.Object(),
			m.String(),
		}

	case crcChannelsJoinChannel:
		r = ChannelsJoinChannel{
			m.Object(),
		}

	case crcChannelsLeaveChannel:
		r = ChannelsLeaveChannel{
			m.Object(),
		}

	case crcChannelsInviteToChannel:
		r = ChannelsInviteToChannel{
			m.Object(),
			m.Vector(),
		}

	case crcChannelsDeleteChannel:
		r = ChannelsDeleteChannel{
			m.Object(),
		}

	case crcChannelsExportMessageLink:
		r = ChannelsExportMessageLink{
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crcChannelsToggleSignatures:
		r = ChannelsToggleSignatures{
			m.Object(),
			m.Object(),
		}

	case crcChannelsGetAdminedPublicChannels:
		var flags int32
		r = ChannelsGetAdminedPublicChannels{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
		}

	case crcChannelsEditBanned:
		r = ChannelsEditBanned{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crcChannelsGetAdminLog:
		var flags int32
		r = ChannelsGetAdminLog{
			readFlags(m, &flags),
			m.Object(),
			m.String(),
			m.FlaggedObject(flags, 0),
			m.FlaggedVector(flags, 1),
			m.Long(),
			m.Long(),
			m.Int(),
		}

	case crcChannelsSetStickers:
		r = ChannelsSetStickers{
			m.Object(),
			m.Object(),
		}

	case crcChannelsReadMessageContents:
		r = ChannelsReadMessageContents{
			m.Object(),
			m.VectorInt(),
		}

	case crcChannelsDeleteHistory:
		r = ChannelsDeleteHistory{
			m.Object(),
			m.Int(),
		}

	case crcChannelsTogglePreHistoryHidden:
		r = ChannelsTogglePreHistoryHidden{
			m.Object(),
			m.Object(),
		}

	case crcChannelsGetLeftChannels:
		r = ChannelsGetLeftChannels{
			m.Int(),
		}

	case crcChannelsGetGroupsForDiscussion:
		r = ChannelsGetGroupsForDiscussion{}

	case crcChannelsSetDiscussionGroup:
		r = ChannelsSetDiscussionGroup{
			m.Object(),
			m.Object(),
		}

	case crcChannelsEditCreator:
		r = ChannelsEditCreator{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crcChannelsEditLocation:
		r = ChannelsEditLocation{
			m.Object(),
			m.Object(),
			m.String(),
		}

	case crcChannelsToggleSlowMode:
		r = ChannelsToggleSlowMode{
			m.Object(),
			m.Int(),
		}

	case crcChannelsGetInactiveChannels:
		r = ChannelsGetInactiveChannels{}

	case crcBotsSendCustomRequest:
		r = BotsSendCustomRequest{
			m.String(),
			m.Object(),
		}

	case crcBotsAnswerWebhookJSONQuery:
		r = BotsAnswerWebhookJSONQuery{
			m.Long(),
			m.Object(),
		}

	case crcBotsSetBotCommands:
		r = BotsSetBotCommands{
			m.Vector(),
		}

	case crcPaymentsGetPaymentForm:
		r = PaymentsGetPaymentForm{
			m.Int(),
		}

	case crcPaymentsGetPaymentReceipt:
		r = PaymentsGetPaymentReceipt{
			m.Int(),
		}

	case crcPaymentsValidateRequestedInfo:
		var flags int32
		r = PaymentsValidateRequestedInfo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Int(),
			m.Object(),
		}

	case crcPaymentsSendPaymentForm:
		var flags int32
		r = PaymentsSendPaymentForm{
			readFlags(m, &flags),
			m.Int(),
			m.FlaggedString(flags, 0),
			m.FlaggedString(flags, 1),
			m.Object(),
		}

	case crcPaymentsGetSavedInfo:
		r = PaymentsGetSavedInfo{}

	case crcPaymentsClearSavedInfo:
		var flags int32
		r = PaymentsClearSavedInfo{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
		}

	case crcPaymentsGetBankCardData:
		r = PaymentsGetBankCardData{
			m.String(),
		}

	case crcStickersCreateStickerSet:
		var flags int32
		r = StickersCreateStickerSet{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			flags&2 != 0, //flag #1
			m.Object(),
			m.String(),
			m.String(),
			m.FlaggedObject(flags, 2),
			m.Vector(),
		}

	case crcStickersRemoveStickerFromSet:
		r = StickersRemoveStickerFromSet{
			m.Object(),
		}

	case crcStickersChangeStickerPosition:
		r = StickersChangeStickerPosition{
			m.Object(),
			m.Int(),
		}

	case crcStickersAddStickerToSet:
		r = StickersAddStickerToSet{
			m.Object(),
			m.Object(),
		}

	case crcStickersSetStickerSetThumb:
		r = StickersSetStickerSetThumb{
			m.Object(),
			m.Object(),
		}

	case crcPhoneGetCallConfig:
		r = PhoneGetCallConfig{}

	case crcPhoneRequestCall:
		var flags int32
		r = PhoneRequestCall{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crcPhoneAcceptCall:
		r = PhoneAcceptCall{
			m.Object(),
			m.StringBytes(),
			m.Object(),
		}

	case crcPhoneConfirmCall:
		r = PhoneConfirmCall{
			m.Object(),
			m.StringBytes(),
			m.Long(),
			m.Object(),
		}

	case crcPhoneReceivedCall:
		r = PhoneReceivedCall{
			m.Object(),
		}

	case crcPhoneDiscardCall:
		var flags int32
		r = PhoneDiscardCall{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.Object(),
			m.Long(),
		}

	case crcPhoneSetCallRating:
		var flags int32
		r = PhoneSetCallRating{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
			m.Int(),
			m.String(),
		}

	case crcPhoneSaveCallDebug:
		r = PhoneSaveCallDebug{
			m.Object(),
			m.Object(),
		}

	case crcPhoneSendSignalingData:
		r = PhoneSendSignalingData{
			m.Object(),
			m.StringBytes(),
		}

	case crcLangpackGetLangPack:
		r = LangpackGetLangPack{
			m.String(),
			m.String(),
		}

	case crcLangpackGetStrings:
		r = LangpackGetStrings{
			m.String(),
			m.String(),
			m.VectorString(),
		}

	case crcLangpackGetDifference:
		r = LangpackGetDifference{
			m.String(),
			m.String(),
			m.Int(),
		}

	case crcLangpackGetLanguages:
		r = LangpackGetLanguages{
			m.String(),
		}

	case crcLangpackGetLanguage:
		r = LangpackGetLanguage{
			m.String(),
			m.String(),
		}

	case crcFoldersEditPeerFolders:
		r = FoldersEditPeerFolders{
			m.Vector(),
		}

	case crcFoldersDeleteFolder:
		r = FoldersDeleteFolder{
			m.Int(),
		}

	case crcStatsGetBroadcastStats:
		var flags int32
		r = StatsGetBroadcastStats{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	case crcStatsLoadAsyncGraph:
		var flags int32
		r = StatsLoadAsyncGraph{
			readFlags(m, &flags),
			m.String(),
			m.FlaggedLong(flags, 0),
		}

	case crcStatsGetMegagroupStats:
		var flags int32
		r = StatsGetMegagroupStats{
			readFlags(m, &flags),
			flags&1 != 0, //flag #0
			m.Object(),
		}

	default:
		m.err = merry.Errorf("Unknown constructor: \u002508x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}
