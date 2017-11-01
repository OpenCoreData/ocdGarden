### Tile 38 test

#### About
This is an experiment with Tile 38.  

http://tile38.com

The client to this (since we are geohashing) is just Redigo (Redis API complient client).

```
set test hilo point 19.705627232977267 -155.093994140625 
set test hono point 21.300570216749353 -157.8680419921875 
set test aroundHaw point 20.514981807048372 -155.9893798828125 
set test MauiPoint point 20.546329665198517 -156.0552978515625 
set test Kula point 20.698436036336485 -156.29837036132812 
set test SouthBigIsl point 19.005970464828987 -155.9454345703125 
```



Test geoJson

geoJSON
```
{
  "type": "FeatureCollection",
  "features": [
    {
      "type": "Feature",
      "properties": {},
      "geometry": {
        "type": "Polygon",
        "coordinates": [
          [
            [
              -158.060302734375,
              20.601936194281016
            ],
            [
              -157.313232421875,
              19.621892180319374
            ],
            [
              -154.87426757812497,
              20.673905264672843
            ],
            [
              -155.85205078125,
              21.85130210558968
            ],
            [
              -158.060302734375,
              20.601936194281016
            ]
          ]
        ]
      }
    }
  ]
}
```