package nex_datastore_super_mario_maker

import (
	"github.com/PretendoNetwork/nex-go/v2"
	"github.com/PretendoNetwork/nex-go/v2/types"
	datastore_super_mario_maker "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/super-mario-maker"
	datastore_types "github.com/PretendoNetwork/nex-protocols-go/v2/datastore/types"
	datastore_smm_db "github.com/PretendoNetwork/super-mario-maker/database/datastore/super-mario-maker"
	"github.com/PretendoNetwork/super-mario-maker/globals"
)

func RecommendedCourseSearchObject(err error, packet nex.PacketInterface, callID uint32, param datastore_types.DataStoreSearchParam, extraData types.List[types.String]) (*nex.RMCMessage, *nex.Error) {
	if err != nil {
		globals.Logger.Error(err.Error())
		return nil, nex.NewError(nex.ResultCodes.DataStore.Unknown, err.Error())
	}

	// * This method is used in 100 Mario and Course World
	// *
	// * extraData seems to be a set of filters defining a
	// * range for a courses success rate
	// *
	// * Course World (All)          ["",  "",   "",    "0", "0"]
	// * Course World (Easy)         ["1", "0",  "34",  "0", "0"]
	// * Course World (Normal)       ["1", "35", "74",  "0", "0"]
	// * Course World (Expert)       ["1", "75", "95",  "0", "0"]
	// * Course World (Super Expert) ["1", "96", "100", "0", "0"]
	// *
	// * Indexes 1 and 2 seem to be a min and max for the *failure*
	// * rate of the courses. This is not taken into account yet,
	// * as the SQL query for this would need to be rather complex.
	// * The last 2 values always seem to be 0, and the first seems
	// * to always be 1 besides filtering for "All"
	// *
	// ! All requests are treated as filtering for "All" right now
	// TODO - Use these ranges to properly filter by difficulty

	// HACK The database load is exponential here and
	length := int(param.ResultRange.Length)
	maxLength := 25
	if length < 0 || length > maxLength {
		globals.Logger.Warningf("Limiting request to %d courses (was %d)", maxLength, length)
		length = maxLength
	}

	// TODO - Use the offet? Real client never uses it, but might be nice for completeness sake?
	pRankingResults, nexError := datastore_smm_db.GetRandomCoursesWithLimit(length)
	if nexError != nil {
		return nil, nexError
	}

	rmcResponseStream := nex.NewByteStreamOut(globals.SecureServer.LibraryVersions, globals.SecureServer.ByteStreamSettings)

	pRankingResults.WriteTo(rmcResponseStream)

	rmcResponse := nex.NewRMCSuccess(globals.SecureEndpoint, rmcResponseStream.Bytes())
	rmcResponse.ProtocolID = datastore_super_mario_maker.ProtocolID
	rmcResponse.MethodID = datastore_super_mario_maker.MethodRecommendedCourseSearchObject
	rmcResponse.CallID = callID

	return rmcResponse, nil
}
