package iferr

import (
	"fmt"
	"net/http"
	"os"
	"reflect"

	"github.com/go-dockly/utility/xlogger"
	"github.com/pkg/errors"

	"github.com/imdario/mergo"
)

var Default *IfErr

func init() {
	var err error

	log, err := xlogger.New(&xlogger.Config{
		CallerDepth: 4,
	})

	if err != nil {
		log.Error("Failed to initialize the default IfErr %+v\n", err)
	}

	Default, err = New(WithLogger(log))

	if err != nil {
		log.Error("Failed to initialize the default IfErr %+v\n", err)
	}
}

type IfErr struct {
	*Options
}

type Options struct {
	Log     *xlogger.Logger
	Verbose *bool
}

type OptionFunc func(*Options)

func New(optFns ...OptionFunc) (*IfErr, error) {
	opts := &Options{}

	for _, optFn := range optFns {
		optFn(opts)
	}

	if opts.Verbose == nil {
		true := true
		opts.Verbose = &true
	}

	if opts.Log == nil {
		log, err := xlogger.New(&xlogger.Config{
			CallerDepth: 3,
		})

		if err != nil {
			return nil, err
		}

		opts.Log = log
	}

	return &IfErr{
		Options: opts,
	}, nil
}

func WithLogger(log *xlogger.Logger) OptionFunc {
	return func(o *Options) {
		o.Log = log
	}
}

func Verbose(flag bool) OptionFunc {
	return func(o *Options) {
		o.Verbose = &flag
	}
}

func Warn(err error) { Default.Warn(err) }
func (ie *IfErr) Warn(err error) {
	if err != nil {
		if *ie.Verbose {
			ie.Log.Warnf("%+v\n", err)
		} else {
			ie.Log.Warnf("%v\n", err)
		}
	}
}

func Exit(err error) { Default.Exit(err) }
func (ie *IfErr) Exit(err error) {
	if err != nil {
		if *ie.Verbose {
			ie.Log.Errorf("%+v\n", err)
		} else {
			ie.Log.Errorf("%v\n", err)
		}
		os.Exit(-1)
	}
}

func Panic(err error) { Default.Panic(err) }
func (ie *IfErr) Panic(err error) {
	if err != nil {
		var message string

		if *ie.Verbose {
			message = fmt.Sprintf("%+v\n", err)
		} else {
			message = fmt.Sprintf("%v\n", err)
		}

		panic(message)
	}
}

type Fataler interface {
	Fatalf(format string, args ...interface{})
}

func Fail(f Fataler, err error) { Default.Fail(f, err) }
func (ie *IfErr) Fail(f Fataler, err error) {
	if err != nil {
		f.Fatalf("%+v\n", err)
	}
}

type ResponseOpts struct {
	Code    int
	Message string
	Depth   int
}

func Respond(w http.ResponseWriter, err error, opts ...*ResponseOpts) bool {
	return Default.Respond(w, err, opts...)
}
func (ie *IfErr) Respond(w http.ResponseWriter, err error, opts ...*ResponseOpts) bool {
	if err == nil {
		return false
	}

	opt := &ResponseOpts{
		Code: http.StatusInternalServerError,
	}

	if len(opts) > 0 {
		err := Merge(opt, opts[0])
		ie.Warn(err)
	}

	if opt.Message == "" {
		opt.Message = errors.Message(err, opt.Depth)
	}

	if opt.Message == "" {
		opt.Message = err.Error()
	}

	if *ie.Verbose {
		ie.Log.Error("HTTP Error: ", err)
	} else {
		ie.Log.Error("HTTP Error: ", err)
	}

	http.Error(w, opt.Message, opt.Code)
	return true
}

func Merge(dst, src interface{}) error {
	if src != nil && !reflect.ValueOf(src).IsNil() {
		return mergo.Merge(dst, src, mergo.WithOverride)
	}

	return nil
}
