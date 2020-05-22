package nasaData

type SolDays struct {
	Day520  Day520   `json:"520"`
	Day521  Day521   `json:"521"`
	Day522  Day522   `json:"522"`
	Day523  Day523   `json:"523"`
	Day524  Day524   `json:"524"`
	Day525  Day525   `json:"525"`
	Day526  Day526   `json:"526"`
	SolKeys []string `json:"sol_keys"`
	// ValidityChecks ValidityChecks `json:"validity_checks"`
}
type AT struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}
type HWS struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}
type PRE struct {
	Av float64 `json:"av"`
	Ct int     `json:"ct"`
	Mn float64 `json:"mn"`
	Mx float64 `json:"mx"`
}
type Dev0 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev1 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev2 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev3 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev5 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev6 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev7 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev8 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev9 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev10 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev11 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev12 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev13 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev14 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type Dev15 struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type MostCommon struct {
	CompassDegrees float64 `json:"compass_degrees"`
	CompassPoint   string  `json:"compass_point"`
	CompassRight   float64 `json:"compass_right"`
	CompassUp      float64 `json:"compass_up"`
	Ct             int     `json:"ct"`
}
type WD struct {
	Dev0       Dev0       `json:"0"`
	Dev1       Dev1       `json:"1"`
	Dev2       Dev2       `json:"2"`
	Dev3       Dev3       `json:"3"`
	Dev5       Dev5       `json:"5"`
	Dev6       Dev6       `json:"6"`
	Dev7       Dev7       `json:"7"`
	Dev8       Dev8       `json:"8"`
	Dev9       Dev9       `json:"9"`
	Dev10      Dev10      `json:"10"`
	Dev11      Dev11      `json:"11"`
	Dev12      Dev12      `json:"12"`
	Dev13      Dev13      `json:"13"`
	Dev14      Dev14      `json:"14"`
	Dev15      Dev15      `json:"15"`
	MostCommon MostCommon `json:"most_common"`
}
type Day520 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}
type Day521 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}
type Day522 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}
type Day523 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}
type Day524 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}
type Day525 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}
type Day526 struct {
	AT     AT     `json:"AT"`
	HWS    HWS    `json:"HWS"`
	PRE    PRE    `json:"PRE"`
	Season string `json:"Season"`
	WD     WD     `json:"WD"`
}

/* type ATValid struct {
	SolHoursWithData []int `json:"sol_hours_with_data"`
	Valid            bool  `json:"valid"`
}
type HWSValid struct {
	SolHoursWithData []int `json:"sol_hours_with_data"`
	Valid            bool  `json:"valid"`
}
type PREValid struct {
	SolHoursWithData []int `json:"sol_hours_with_data"`
	Valid            bool  `json:"valid"`
}
type WDValid struct {
	SolHoursWithData []int `json:"sol_hours_with_data"`
	Valid            bool  `json:"valid"`
}
type Day519Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day520Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day521Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day522Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day523Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day524Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day525Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type Day526Val struct {
	AT  AT  `json:"AT"`
	HWS HWS `json:"HWS"`
	PRE PRE `json:"PRE"`
	WD  WD  `json:"WD"`
}
type ValidityChecks struct {
	Day519Val        Day519Val `json:"519"`
	Day520Val        Day520Val `json:"520"`
	Day521Val        Day521Val `json:"521"`
	Day522Val        Day522Val `json:"522"`
	Day523Val        Day523Val `json:"523"`
	Day524Val        Day524Val `json:"524"`
	Day525Val        Day525Val `json:"525"`
	Day526Val        Day526Val `json:"526"`
	SolHoursRequired int       `json:"sol_hours_required"`
	SolsChecked      []string  `json:"sols_checked"`
}
*/
