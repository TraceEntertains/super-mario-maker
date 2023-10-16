package nex

import (
	datastore_super_mario_maker "github.com/PretendoNetwork/nex-protocols-go/datastore/super-mario-maker"
	"github.com/PretendoNetwork/super-mario-maker-secure/globals"
	nex_datastore "github.com/PretendoNetwork/super-mario-maker-secure/nex/datastore"
	nex_datastore_super_mario_maker "github.com/PretendoNetwork/super-mario-maker-secure/nex/datastore/super-mario-maker"
)

func registerNEXProtocols() {
	datastoreSMMProtocol := datastore_super_mario_maker.NewProtocol(globals.SecureServer)

	datastoreSMMProtocol.GetMeta(nex_datastore.GetMeta)
	datastoreSMMProtocol.PreparePostObject(nex_datastore.PreparePostObject)
	datastoreSMMProtocol.PrepareGetObject(nex_datastore.PrepareGetObject)
	datastoreSMMProtocol.CompletePostObject(nex_datastore.CompletePostObject)
	datastoreSMMProtocol.GetMetasMultipleParam(nex_datastore.GetMetasMultipleParam)
	datastoreSMMProtocol.ChangeMeta(nex_datastore.ChangeMeta)
	datastoreSMMProtocol.RateObjects(nex_datastore.RateObjects)
	datastoreSMMProtocol.GetObjectInfos(nex_datastore_super_mario_maker.GetObjectInfos)
	datastoreSMMProtocol.RateCustomRanking(nex_datastore_super_mario_maker.RateCustomRanking)
	datastoreSMMProtocol.GetCustomRankingByDataID(nex_datastore_super_mario_maker.GetCustomRankingByDataID)
	datastoreSMMProtocol.AddToBufferQueues(nex_datastore_super_mario_maker.AddToBufferQueues)
	datastoreSMMProtocol.GetBufferQueue(nex_datastore_super_mario_maker.GetBufferQueue)
	datastoreSMMProtocol.CompleteAttachFile(nex_datastore_super_mario_maker.CompleteAttachFile)
	datastoreSMMProtocol.PrepareAttachFile(nex_datastore_super_mario_maker.PrepareAttachFile)
	datastoreSMMProtocol.GetApplicationConfig(nex_datastore_super_mario_maker.GetApplicationConfig)
	datastoreSMMProtocol.FollowingsLatestCourseSearchObject(nex_datastore_super_mario_maker.FollowingsLatestCourseSearchObject)
	datastoreSMMProtocol.RecommendedCourseSearchObject(nex_datastore_super_mario_maker.RecommendedCourseSearchObject)
	datastoreSMMProtocol.SuggestedCourseSearchObject(nex_datastore_super_mario_maker.SuggestedCourseSearchObject)
	datastoreSMMProtocol.UploadCourseRecord(nex_datastore_super_mario_maker.UploadCourseRecord)
	datastoreSMMProtocol.GetCourseRecord(nex_datastore_super_mario_maker.GetCourseRecord)
	datastoreSMMProtocol.GetApplicationConfigString(nex_datastore_super_mario_maker.GetApplicationConfigString)
	datastoreSMMProtocol.GetDeletionReason(nex_datastore_super_mario_maker.GetDeletionReason)
	datastoreSMMProtocol.GetMetasWithCourseRecord(nex_datastore_super_mario_maker.GetMetasWithCourseRecord)
	datastoreSMMProtocol.CheckRateCustomRankingCounter(nex_datastore_super_mario_maker.CheckRateCustomRankingCounter)
	datastoreSMMProtocol.CTRPickUpCourseSearchObject(nex_datastore_super_mario_maker.CTRPickUpCourseSearchObject)
}
