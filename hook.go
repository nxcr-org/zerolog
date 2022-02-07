package zerolog

import "context"

// Hook defines an interface to a log hook.
type Hook interface {
	// Run runs the hook with the event.
	Run(ctx context.Context, e *Event, level Level, message string)
}

// HookFunc is an adaptor to allow the use of an ordinary function
// as a Hook.
type HookFunc func(ctx context.Context, e *Event, level Level, message string)

// Run implements the Hook interface.
func (h HookFunc) Run(ctx context.Context, e *Event, level Level, message string) {
	h(ctx, e, level, message)
}

// LevelHook applies a different hook for each level.
type LevelHook struct {
	NoLevelHook, TraceHook, DebugHook, InfoHook, WarnHook, ErrorHook, FatalHook, PanicHook Hook
}

// Run implements the Hook interface.
func (h LevelHook) Run(ctx context.Context, e *Event, level Level, message string) {
	switch level {
	case TraceLevel:
		if h.TraceHook != nil {
			h.TraceHook.Run(ctx, e, level, message)
		}
	case DebugLevel:
		if h.DebugHook != nil {
			h.DebugHook.Run(ctx, e, level, message)
		}
	case InfoLevel:
		if h.InfoHook != nil {
			h.InfoHook.Run(ctx, e, level, message)
		}
	case WarnLevel:
		if h.WarnHook != nil {
			h.WarnHook.Run(ctx, e, level, message)
		}
	case ErrorLevel:
		if h.ErrorHook != nil {
			h.ErrorHook.Run(ctx, e, level, message)
		}
	case FatalLevel:
		if h.FatalHook != nil {
			h.FatalHook.Run(ctx, e, level, message)
		}
	case PanicLevel:
		if h.PanicHook != nil {
			h.PanicHook.Run(ctx, e, level, message)
		}
	case NoLevel:
		if h.NoLevelHook != nil {
			h.NoLevelHook.Run(ctx, e, level, message)
		}
	}
}

// NewLevelHook returns a new LevelHook.
func NewLevelHook() LevelHook {
	return LevelHook{}
}
