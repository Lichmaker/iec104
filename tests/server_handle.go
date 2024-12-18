package tests

import (
	"github.com/wendy512/go-iecp5/asdu"
	"time"
)

const (
	commonAddr = 1
)

type myServerHandler struct {
}

func (ms *myServerHandler) OnInterrogation(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierOfInterrogation) error {
	//_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.Single(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.SinglePointInfo{
		Ioa:   100,
		Value: true,
		Qds:   asdu.QDSGood,
	})
	_ = asdu.Double(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.DoublePointInfo{
		Ioa:   200,
		Value: asdu.DPIDeterminedOn,
		Qds:   asdu.QDSGood,
	})
	//_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnCounterInterrogation(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierCountCall) error {
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.CounterInterrogationCmd(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, asdu.QualifierCountCall{asdu.QCCGroup1, asdu.QCCFrzRead})
	_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnRead(conn asdu.Connect, pack *asdu.ASDU, addr asdu.InfoObjAddr) error {
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.Single(conn, false, asdu.CauseOfTransmission{Cause: asdu.InterrogatedByStation}, commonAddr, asdu.SinglePointInfo{
		Ioa:   addr,
		Value: true,
		Qds:   asdu.QDSGood,
	})
	_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnClockSync(conn asdu.Connect, pack *asdu.ASDU, tm time.Time) error {
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	now := time.Now()
	_ = asdu.ClockSynchronizationCmd(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, now)
	_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnResetProcess(conn asdu.Connect, pack *asdu.ASDU, quality asdu.QualifierOfResetProcessCmd) error {
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.ResetProcessCmd(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, asdu.QPRGeneralRest)
	_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnDelayAcquisition(conn asdu.Connect, pack *asdu.ASDU, msec uint16) error {
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	_ = asdu.DelayAcquireCommand(conn, asdu.CauseOfTransmission{Cause: asdu.Activation}, commonAddr, msec)
	_ = pack.SendReplyMirror(conn, asdu.ActivationTerm)
	return nil
}

func (ms *myServerHandler) OnASDU(conn asdu.Connect, pack *asdu.ASDU) error {
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	// TODO
	cmd := pack.GetSingleCmd()
	_ = asdu.SingleCmd(conn, pack.Type, pack.Coa, pack.CommonAddr, asdu.SingleCommandInfo{
		Ioa:   cmd.Ioa,
		Value: cmd.Value,
		Qoc:   cmd.Qoc,
	})
	_ = pack.SendReplyMirror(conn, asdu.ActivationCon)
	return nil
}
