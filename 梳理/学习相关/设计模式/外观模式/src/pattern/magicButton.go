package pattern

type MagicButton struct {
	tv          TV
	light       Light
	waterHeater WaterHeater
}

func NewMagicButton(tv TV,light Light,heater WaterHeater) *MagicButton{
	return &MagicButton{
		tv,light,heater,
	}
}

func (mb MagicButton) On() {
	mb.tv.Open()
	mb.light.TurnOn()
	mb.waterHeater.On()
}

func (mb MagicButton) Off() {
	mb.tv.Close()
	mb.light.TurnOff()
	mb.waterHeater.Off()
}
