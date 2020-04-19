package defs

//- 电影图片
//- 电影名字
//- 电影导演
//- 电影评价
//- 电影评语
type Movie struct {
	Url        string
	Name       string `json:"name"`
	Img        string `json:"img"`
	Director   string `json:"director"`
	Evaluation string `json:"evaluation"`
	Comments   string `json:"comments"`
}
//接口 方法的合集
//type MovieInterface interface {
//
//	Add()
//}
