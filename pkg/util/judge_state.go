package util

// judge
const (
	Invalid = iota
	Accepted
	CompileError
	WrongAnswer
	PresentationError
	RuntimeError
	TimeLimitExceeded
	MemoryLimitExceeded
	OutputLimitExceeded
	Waiting
	Judging
	SystemError
)

// rsjudge
const (
	_JUDGE_RESULT_UNSPECIFIED = iota
	_JUDGE_RESULT_ACCEPTED
	_JUDGE_RESULT_COMPILE_ERROR
	_JUDGE_RESULT_WRONG_ANSWER
	_JUDGE_RESULT_PRESENTATION_ERROR
	_JUDGE_RESULT_RUNTIME_ERROR
	_JUDGE_RESULT_TIME_LIMIT_EXCEEDED
	_JUDGE_RESULT_MEMORY_LIMIT_EXCEEDED
	_JUDGE_RESULT_OUTPUT_LIMIT_EXCEEDED
)

// go-judge
const (
	_Invalid = iota
	_Accepted
	_WrongAnswer      // Not used
	_PartiallyCorrect // Not used
	_MemoryLimitExceeded
	_TimeLimitExceeded
	_OutputLimitExceeded
	_FileError
	_NonZeroExitStatus
	_Signalled
	_DangerousSyscall
	_JudgementFailed    // Not used
	_InvalidInteraction // Not used
	_InternalError
)

func FromGoJudgeState(state int32) (res int16) {
	switch state {
	case _Invalid:
		res = Invalid
	case _Accepted:
		res = Accepted
	case _WrongAnswer:
		res = WrongAnswer
	case _PartiallyCorrect:
		res = Invalid
	case _MemoryLimitExceeded:
		res = MemoryLimitExceeded
	case _TimeLimitExceeded:
		res = TimeLimitExceeded
	case _OutputLimitExceeded:
		res = OutputLimitExceeded
	case _FileError:
		res = Invalid
	case _NonZeroExitStatus:
		res = RuntimeError
	case _Signalled:
		res = Invalid
	case _DangerousSyscall:
		res = Invalid
	case _JudgementFailed:
		res = Invalid
	case _InvalidInteraction:
		res = Invalid
	case _InternalError:
		res = Invalid
	default:
		res = Invalid
	}
	return
}

func FromRsjudgeState(state int) (res int16) {
	switch state {
	case _JUDGE_RESULT_UNSPECIFIED:
		res = Invalid
	case _JUDGE_RESULT_ACCEPTED:
		res = Accepted
	case _JUDGE_RESULT_COMPILE_ERROR:
		res = CompileError
	case _JUDGE_RESULT_WRONG_ANSWER:
		res = WrongAnswer
	case _JUDGE_RESULT_PRESENTATION_ERROR:
		res = PresentationError
	case _JUDGE_RESULT_RUNTIME_ERROR:
		res = RuntimeError
	case _JUDGE_RESULT_TIME_LIMIT_EXCEEDED:
		res = TimeLimitExceeded
	case _JUDGE_RESULT_MEMORY_LIMIT_EXCEEDED:
		res = MemoryLimitExceeded
	case _JUDGE_RESULT_OUTPUT_LIMIT_EXCEEDED:
		res = OutputLimitExceeded
	default:
		res = Invalid
	}
	return
}
