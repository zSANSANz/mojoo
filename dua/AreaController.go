package dua

import( 

)

func (_r *AreaRepository) InsertArea(param1 int32, param2 int64, type []string, ar *Model.Area)
(err error) {
inst := _r.DB.Model(ar)
Var area int
area = 0
switch type {
case "persegi panjang":
var area := param1 * param2
ar.AreaValue = area
ar.AreaType = "persegi panjang"
err = _r.DB.create(&ar).Error
if err != nil {
return err
}
case "persegi":
var area = param1 * param2
ar.AreaValue = area
ar.AreaType = "persegi"
err = _r.DB.create(&ar).Error
if err != nil {
return err
}
case "segitiga":
area = 0.5 * (param1 * param2)
ar.AreaValue = area
ar.AreaType = "segitiga"
err = _r.DB.create(&ar).Error
if err != nil {
return err
}
default:
ar.AreaValue = 0
ar.AreaType = "undefined data"
err = _r.DB.create(&ar).Error
if err != nil {
return err
}
}
}

