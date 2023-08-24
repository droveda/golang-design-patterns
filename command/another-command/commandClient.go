package anothercommand

func CommandClient() {
	tv := &Tv{}

	onCommand := &OnCommand{
		Device: tv,
	}

	offCommand := &OffCommand{
		Device: tv,
	}

	onButton := &Button{
		Command: onCommand,
	}

	onButton.execute()

	offButton := &Button{
		Command: offCommand,
	}

	offButton.execute()
}
