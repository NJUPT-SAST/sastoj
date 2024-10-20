package gojudge

import (
	"sastoj/pkg/util"

	"github.com/criyle/go-judge/pb"
)

func Convert(state pb.Response_Result_StatusType) (res int16) {
	switch state {
	case pb.Response_Result_Invalid:
		res = util.Invalid
	case pb.Response_Result_Accepted:
		res = util.Accepted
	case pb.Response_Result_WrongAnswer:
		res = util.WrongAnswer
	case pb.Response_Result_PartiallyCorrect:
		res = util.Invalid
	case pb.Response_Result_MemoryLimitExceeded:
		res = util.MemoryLimitExceeded
	case pb.Response_Result_TimeLimitExceeded:
		res = util.TimeLimitExceeded
	case pb.Response_Result_OutputLimitExceeded:
		res = util.OutputLimitExceeded
	case pb.Response_Result_FileError:
		res = util.Invalid
	case pb.Response_Result_NonZeroExitStatus:
		res = util.RuntimeError
	case pb.Response_Result_Signalled:
		res = util.Invalid
	case pb.Response_Result_DangerousSyscall:
		res = util.Invalid
	case pb.Response_Result_JudgementFailed:
		res = util.Invalid
	case pb.Response_Result_InvalidInteraction:
		res = util.Invalid
	case pb.Response_Result_InternalError:
		res = util.Invalid
	default:
		res = util.Invalid
	}
	return
}
