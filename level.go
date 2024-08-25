package slslog

import "log/slog"

var levelCritical = level{val: int(slog.LevelError) + 4}

type level struct {
	val int
}

func (v *level) Level() slog.Level {
	return slog.Level(v.val)
}
