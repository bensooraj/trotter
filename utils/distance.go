package utils

import "math"

// function getDistanceFromLatLonInKm(lat1,lon1,lat2,lon2) {
// 	var R = 6371; // Radius of the earth in km
// 	var dLat = Deg2Rad(lat2-lat1);  // Deg2Rad below
// 	var dLon = Deg2Rad(lon2-lon1);
// 	var a =
// 	  Math.sin(dLat/2) * Math.sin(dLat/2) +
// 	  Math.cos(Deg2Rad(lat1)) * Math.cos(Deg2Rad(lat2)) *
// 	  Math.sin(dLon/2) * Math.sin(dLon/2)
// 	  ;
// 	var c = 2 * Math.atan2(Math.sqrt(a), Math.sqrt(1-a));
// 	var d = R * c; // Distance in km
// 	return d;
//   }

//   function Deg2Rad(deg) {
// 	return deg * (Math.PI/180)
//   }

func GetDistanceFromLatLonInKm(lat1, lon1, lat2, lon2 float64) float64 {
	var R float64 = 6371.0
	var dLat float64 = Deg2Rad(lat2 - lat1)
	var dLon float64 = Deg2Rad(lon2 - lon1)

	var a float64 = ((math.Sin(dLat/2) * math.Sin(dLat/2)) +
		(math.Cos(Deg2Rad(lat1))*math.Cos(Deg2Rad(lat2)))*
			(math.Sin(dLon/2)*math.Sin(dLon/2)))

	var c float64 = math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	var d = R * c

	return d
}

func Deg2Rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
