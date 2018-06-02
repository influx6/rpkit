package history

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

// level constants
const (
	RedLvl Level = iota
	YellowLvl
	ErrorLvl
	InfoLvl
)

// errors ...
var (
	ErrYellowAlert = errors.New("warning: error or invalid state occured")
	ErrRedAlert    = errors.New("very bad: earth damaging error occured, check now")
)

var (
	hl = struct {
		ml  sync.Mutex
		hls Handler
	}{}
)

// SetDefaultHandlers sets default Handlers to be included
// in all instances of Sources to be used to process provided
// BugLogs.
func SetDefaultHandlers(hs ...Handler) bool {
	if len(hs) == 0 {
		return false
	}

	hl.ml.Lock()
	defer hl.ml.Unlock()
	hl.hls = Handlers(hs)
	return true
}

//**************************************************************
// Ctx Interface
//**************************************************************

var empty BugLog

// Attrs defines a map type.
type Attrs map[string]interface{}

// Ctx exposes an interface which provides a means to collate giving
// data for a giving data.
type Ctx interface {
	Free()
	Red(string, ...interface{}) Ctx
	Info(string, ...interface{}) Ctx
	Yellow(string, ...interface{}) Ctx
	Error(error, string, ...interface{}) Ctx

	// Logging methods that skip retrieving function
	// caller details (i.e function name,line number, etc).
	IRed(string, ...interface{}) Ctx
	IInfo(string, ...interface{}) Ctx
	IYellow(string, ...interface{}) Ctx
	IError(error, string, ...interface{}) Ctx

	WithTags(...string) Ctx
	WithFields(Attrs, ...string) Ctx
	WithHandler([]string, ...Handler) Ctx
	WithTitle(string, ...interface{}) Ctx
	With(string, interface{}, ...string) Ctx
}

// WithHandlers  returns a new Source with giving Handlers as receivers of
// all provided instances of B.
func WithHandlers(hs ...Handler) Ctx {
	return empty.WithHandler(nil, hs...)
}

// WithTags returns a new Source with giving tags as default tags
// added to all B instances submitted through Source's Ctx instances.
func WithTags(hs ...string) Ctx {
	return empty.WithTags(hs...)
}

// With returns a new Ctx adding the giving key-value.
func With(k string, v interface{}, tags ...string) Ctx {
	return empty.With(k, v, tags...)
}

// WithFields returns a new Source with giving fields as default tags
// added to all B instances submitted through Source's Ctx instances.
func WithFields(attr Attrs, tags ...string) Ctx {
	return empty.WithFields(attr, tags...)
}

// WithTitle returns a new Ctx setting the BugLog.Title.
func WithTitle(title string, v ...interface{}) Ctx {
	return empty.WithTitle(title, v...)
}

//**************************************************************
// Handler Interface
//**************************************************************

// Recv defines a function type which receives a pointer of B.
type Recv func(BugLog) error

// FilterRecv defines a function type which returns true/false for
// a given BugLog.
type FilterRecv func(BugLog) bool

// Handler exposes a single method to deliver giving B value to
type Handler interface {
	Recv(BugLog) error
}

// HandlerFunc returns a new Handler using provided function for
// calls to the Handler.Recv method.
func HandlerFunc(recv Recv) Handler {
	return fnHandler{rc: recv}
}

// FilterByLevel returns a Handler which will only allow BugLogs
// with Status.Level equal or above provided level.
func FilterByLevel(recv Recv, level Level) Handler {
	return FilterFunc(recv, func(b BugLog) bool {
		if b.Status.Level < level {
			return false
		}
		return true
	})
}

// FilterFunc returns a new Handler which will filter all BugLog received
// by the recv function.
func FilterFunc(recv Recv, filter FilterRecv) Handler {
	return HandlerFunc(func(log BugLog) error {
		if filter(log) {
			return recv(log)
		}
		return nil
	})
}

//Handlers defines a slice of Handlers as a type.
type Handlers []Handler

// Recv calls individual Handler.Recv in slice with BugLog instance.
func (h Handlers) Recv(b BugLog) error {
	for _, hl := range h {
		if hl == nil {
			continue
		}
		if err := hl.Recv(b); err != nil {
			return err
		}
	}
	return nil
}

type fnHandler struct {
	rc Recv
}

// Recv implements the Handler and calls the underline Recv
// function provided with provided B pointer.
func (fn fnHandler) Recv(b BugLog) error {
	return fn.rc(b)
}

//**************************************************************
// Level
//**************************************************************

// Level defines a int type which represent the a giving level of entry for a giving entry.
type Level int

// GetLevel returns Level value for the giving string.
// It returns -1 if it does not know the level string.
func GetLevel(lvl string) Level {
	switch strings.ToLower(lvl) {
	case "red":
		return RedLvl
	case "yellow":
		return YellowLvl
	case "error":
		return ErrorLvl
	case "info":
		return InfoLvl
	}

	return -1
}

// String returns the string version of the Level.
func (l Level) String() string {
	switch l {
	case RedLvl:
		return "red"
	case YellowLvl:
		return "yellow"
	case ErrorLvl:
		return "error"
	case InfoLvl:
		return "info"
	}

	return "UNKNOWN"
}

//**************************************************************
// Location
//**************************************************************

// Location defines the location which an history occured in.
type Location struct {
	Function string `json:"function"`
	Line     int    `json:"line"`
	File     string `json:"file"`
}

// CallGraph embodies a graph representing the areas where a method
// call occured.
type CallGraph struct {
	In Location
	By Location
}

//**************************************************************
// Field
//**************************************************************

// Field represents a giving key-value pair with location details.
type Field struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}

//**************************************************************
// Status struct.
//**************************************************************

// Progress embodies giving messages where giving progress with
// associated message and level for status.
type Status struct {
	Message string      `json:"message"`
	Level   Level       `json:"level"`
	Err     interface{} `json:"err"`
	Graph   CallGraph   `json:"graph"`
}

//**************************************************************
// BugLog struct.
//**************************************************************

var bugPool = sync.Pool{
	New: func() interface{} {
		return new(BugLog)
	},
}

// BugLog represent a giving record of data at a giving period of time.
type BugLog struct {
	Title  string           `json:"title"`
	From   time.Time        `json:"from"`
	Tags   []string         `json:"tags"`
	Status Status           `json:"status"`
	Fields map[string]Field `json:"fields"`

	bugs  Handler
	mapp  mapp
	mtags list
}

// WithHandler adds giving handler into returned Ctx.
func (b *BugLog) WithHandler(tags []string, c ...Handler) Ctx {
	if len(c) == 0 {
		return b
	}

	var cr Handlers

	if len(c) == 1 {
		cr = Handlers{b.bugs, c[0]}
	} else {
		cr = Handlers{Handlers(c), b.bugs}
	}

	br := b.branch().addTags(tags...)
	br.bugs = cr
	return br
}

// WithFields returns a new instance of BugLog from source but unique
// and sets giving fields.
func (b *BugLog) WithFields(kv Attrs, tags ...string) Ctx {
	if len(kv) == 0 {
		return b
	}

	return b.branch().addKVS(kv).addTags(tags...)
}

func (b *BugLog) addKVS(kv Attrs) *BugLog {
	for k, v := range kv {
		b.mapp = addKV(b.mapp, k, Field{Key: k, Value: v})
	}
	return b
}

// With returns a new instance of BugLog from source but unique
// and adds key-value pair into Ctx.
func (b *BugLog) With(k string, v interface{}, tags ...string) Ctx {
	return b.branch().addKV(k, v).addTags(tags...)
}

func (b *BugLog) addKV(k string, v interface{}) *BugLog {
	b.mapp = addKV(b.mapp, k, Field{Key: k, Value: v})
	return b
}

// WithTags returns a new instance of BugLog from source but unique
// and sets giving tags.
func (b *BugLog) WithTags(tags ...string) Ctx {
	if len(tags) == 0 {
		return b
	}

	return b.branch().addTags(tags...)
}

func (b *BugLog) addTags(tags ...string) *BugLog {
	if len(tags) == 0 {
		return b
	}
	b.Tags = append(b.Tags, tags...)
	return b
}

// FromTitle returns a new instance of BugLog from source but unique
// and sets giving title.
func (b *BugLog) WithTitle(title string, v ...interface{}) Ctx {
	if len(v) != 0 {
		title = fmt.Sprintf(title, v...)
	}

	br := b.branch()
	br.Title = title
	return br
}

// IRed logs giving status message at giving time with RedLvl.
// This method will not attempt to get the function line, number and name (its an optimization need).
func (b *BugLog) IRed(msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Message: msg,
		Level:   RedLvl,
		Graph:   CallGraph{},
	})
}

// IYellow logs giving status message at giving time with YellowLvl.
// This method will not attempt to get the function line, number and name (its an optimization need).
func (b *BugLog) IYellow(msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Message: msg,
		Level:   YellowLvl,
		Graph:   CallGraph{},
	})
}

// IError logs giving status message at giving time with ErrorLvl.
// This method will not attempt to get the function line, number and name (its an optimization need).
func (b *BugLog) IError(err error, msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Err:     err,
		Message: msg,
		Level:   ErrorLvl,
		Graph:   CallGraph{},
	})
}

// IInfo logs giving status message at giving time with InfoLvl.
// This method will not attempt to get the function line, number and name (its an optimization need).
func (b *BugLog) IInfo(msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Message: msg,
		Level:   InfoLvl,
		Graph:   CallGraph{},
	})
}

// Red logs giving status message at giving time with RedLvl.
func (b *BugLog) Red(msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Message: msg,
		Level:   RedLvl,
		Graph:   GetMethodGraph(3),
	})
}

// Yellow logs giving status message at giving time with YellowLvl.
func (b *BugLog) Yellow(msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Message: msg,
		Level:   YellowLvl,
		Graph:   GetMethodGraph(3),
	})
}

// Error logs giving status message at giving time with ErrorLvl.
func (b *BugLog) Error(err error, msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Err:     err,
		Message: msg,
		Level:   ErrorLvl,
		Graph:   GetMethodGraph(3),
	})
}

// Info logs giving status message at giving time with InfoLvl.
func (b *BugLog) Info(msg string, vals ...interface{}) Ctx {
	if len(vals) != 0 {
		msg = fmt.Sprintf(msg, vals...)
	}

	return b.logStatus(Status{
		Message: msg,
		Level:   InfoLvl,
		Graph:   GetMethodGraph(3),
	})
}

// Free is an optimization function which returns the instance of
// BugLog into a sync pool. It is not safe to call concurrently and
// should be called after clear use of no more of this particular
// object.
func (b *BugLog) Free() {
	b.bugs = nil
	b.mapp = nil
	b.mtags = nil
	bugPool.Put(b)
}

func (b *BugLog) logStatus(s Status) Ctx {
	blog := *b.branch()
	blog.Status = s
	blog.bugs = nil

	if b.mapp != nil {
		blog.Fields = packMap(blog.mapp)
		blog.mapp = nil
	}

	if b.mtags != nil {
		blog.Tags = packList(blog.mtags)
		blog.mtags = nil
	}

	if blog.Title == "" {
		blog.Title = s.Graph.By.Function
	}

	if b.bugs != nil {
		if err := b.bugs.Recv(blog); err != nil {
			log.Printf("error logging: %+s", err)
		}
	}

	hl.ml.Lock()
	defer hl.ml.Unlock()
	if hl.hls != nil {
		if err := hl.hls.Recv(blog); err != nil {
			log.Printf("error logging: %+s", err)
		}
	}

	return b
}

// branch duplicates BugLog and copies appropriate
// dataset over to ensure uniqueness of values to source.
func (b *BugLog) branch() *BugLog {
	if b.mapp != nil {
		if b.mapp.Count() >= 10 {
			b.mapp = mmap(packMap(b.mapp))
		}
	}

	if b.Tags != nil {
		if b.mtags.Count() >= 10 {
			b.mtags = packedlist(packList(b.mtags))
		}
	}

	br := bugPool.Get().(*BugLog)
	br.Title = b.Title
	br.From = time.Now()

	if b.mapp == nil {
		br.mapp = mmap{}
	} else {
		br.mapp = b.mapp
	}

	if b.mtags == nil {
		br.mtags = packedlist{}
	} else {
		br.mtags = b.mtags
	}

	return br
}

//**************************************************************
// persistent map
//**************************************************************

type list interface {
	Count() int
	Pack([]string) []string
	Iterate(func(v string))
}

func packList(m list) []string {
	if mo, ok := m.(packedlist); ok {
		return []string(mo)
	}

	items := make([]string, 0, m.Count())
	return m.Pack(items)
}

type packedlist []string

// Pack provides a optimization, if its self can pack, then
// pack into provided list and return.
func (n packedlist) Pack(nx []string) []string {
	return append(nx, n...)
}

func (n packedlist) Count() int {
	return len(n)
}

func (n packedlist) Iterate(fn func(string)) {
	for _, k := range n {
		fn(k)
	}
}

type nlist struct {
	parent list
	k      string
	c      int
}

func (n *nlist) Count() int {
	return n.c
}

func (n *nlist) Pack(nx []string) []string {
	return append(nx, n.k)
}

func (n *nlist) Iterate(fn func(string)) {
	fn(n.k)
	if n.parent != nil {
		n.parent.Iterate(fn)
	}
}

func addList(m list, k string) list {
	next := new(nlist)
	next.parent = m
	next.k = k
	next.c = m.Count() + 1
	return next
}

type mapp interface {
	Count() int
	Iterate(func(k string, v Field))
}

func packMap(m mapp) map[string]Field {
	if mo, ok := m.(mmap); ok {
		return map[string]Field(mo)
	}

	fields := make(map[string]Field, m.Count())
	m.Iterate(func(k string, v Field) {
		if _, ok := fields[k]; ok {
			return
		}
		fields[k] = v
	})
	return fields
}

type mmap map[string]Field

func (n mmap) Count() int {
	return len(n)
}
func (n mmap) Iterate(fn func(string, Field)) {
	for k, v := range n {
		fn(k, v)
	}
}

type nmap struct {
	parent mapp
	k      string
	v      Field
	c      int
}

func (n *nmap) Count() int {
	return n.c
}

func (n *nmap) Iterate(fn func(string, Field)) {
	fn(n.k, n.v)
	if n.parent != nil {
		n.parent.Iterate(fn)
	}
}

func addKV(m mapp, k string, v Field) mapp {
	next := new(nmap)
	next.parent = m
	next.k = k
	next.v = v
	next.c = m.Count() + 1
	return next
}

//**************************************************************
// internal methods and impl
//**************************************************************

func makeLocation(d int) Location {
	var loc Location
	loc.Function, loc.File, loc.Line = GetMethod(d)
	return loc
}

type sortedFields []Field

func (s sortedFields) Len() int {
	return len(s)
}

func (s sortedFields) Less(i, j int) bool {
	return s[i].Key < s[j].Key
}

func (s sortedFields) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}
