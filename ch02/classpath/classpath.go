package classpath


type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}

func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {

}

func (self *Classpath) String() string {
}

func (classpath *Classpath) parseBootAndExtClasspath(jreOption string) {

}
func (classpath *Classpath) parseUserClasspath(cpOption string) {

}


