var mapFunction1 = function() {
     emit(this.opencoreleg, this.opencoremeasurement);
};
                   
var reduceFunction1 = function(leg, measurement) {
     reducedVal = {measurement:measurement, count: 0 };
     
    // need to loop on leg and 
     for (var idx = 0; idx < measurement.length; idx++) {
            reducedVal.count++;                     
     }

     return reducedVal;
};
                          
db.schemaorg.mapReduce(
    mapFunction1,
    reduceFunction1,
    { out: "map_reduce_example" }
)