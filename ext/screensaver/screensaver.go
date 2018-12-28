package screensaver

import "github.com/linuxdeepin/go-x11-client"

type NotifyEvent struct {
	State    uint8
	Sequence uint16
	Time     x.Timestamp
	Root     x.Window
	Window   x.Window
	Kind     uint8
	Forced   bool
}

func readNotifyEvent(r *x.Reader, v *NotifyEvent) error {
	// code
	r.ReadPad(1)
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	v.Sequence = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Time = x.Timestamp(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Root = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Window = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.Kind = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Forced = x.Uint8ToBool(r.Read1b())
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(14)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeQueryVersion(clientMajorVersion, clientMinorVersion uint8) (b x.RequestBody) {
	b.AddBlock(1).
		Write1b(clientMajorVersion).
		Write1b(clientMinorVersion).
		WritePad(2).
		End()
	return
}

type QueryVersionReply struct {
	ServerMajorVersion uint16
	ServerMinorVersion uint16
}

func readQueryVersionReply(r *x.Reader, v *QueryVersionReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMajorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	v.ServerMinorVersion = r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	r.ReadPad(20)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeQueryInfo(drawable x.Drawable) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(drawable)).
		End()
	return
}

type QueryInfoReply struct {
	State            uint8
	SaverWindow      x.Window
	MsUntilServer    uint32
	MsSinceUserInput uint32
	EventMask        uint32
	Kind             uint8
}

func readQueryInfoReply(r *x.Reader, v *QueryInfoReply) error {
	r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	v.State = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// seq
	r.Read2b()
	if r.Err() != nil {
		return r.Err()
	}

	// length
	r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.SaverWindow = x.Window(r.Read4b())
	if r.Err() != nil {
		return r.Err()
	}

	v.MsUntilServer = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.MsSinceUserInput = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.EventMask = r.Read4b()
	if r.Err() != nil {
		return r.Err()
	}

	v.Kind = r.Read1b()
	if r.Err() != nil {
		return r.Err()
	}

	// read field Pad0
	r.ReadPad(7)
	if r.Err() != nil {
		return r.Err()
	}

	return nil
}

// #WREQ
func encodeSelectInput(drawable x.Drawable, eventMask uint32) (b x.RequestBody) {
	b.AddBlock(2).
		Write4b(uint32(drawable)).
		Write4b(eventMask).
		End()
	return
}

// #WREQ
func encodeSetAttributes(drawable x.Drawable, X, y int16, width, height,
	boardWidth uint16, class, depth uint8, visual x.VisualID, valueMask uint32,
	valueList []uint32) (b x.RequestBody) {

	b0 := b.AddBlock(6 + len(valueList)).
		Write4b(uint32(drawable)).
		Write2b(uint16(X)).
		Write2b(uint16(y)).
		Write2b(width).
		Write2b(height).
		Write2b(boardWidth).
		Write1b(class).
		Write1b(depth).
		Write4b(uint32(visual)).
		Write4b(valueMask)

	for _, value := range valueList {
		b0.Write4b(value)
	}
	b0.End()
	return
}

// #WREQ
func encodeUnsetAttributes(drawable x.Drawable) (b x.RequestBody) {
	b.AddBlock(1).
		Write4b(uint32(drawable)).
		End()
	return
}

// #WREQ
func encodeSuspend(suspend bool) (b x.RequestBody) {
	b.AddBlock(1).
		Write1b(x.BoolToUint8(suspend)).
		WritePad(3).
		End()
	return
}
