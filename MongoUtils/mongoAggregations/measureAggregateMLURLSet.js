db.schemaorg.aggregate(

  // Pipeline
  [
    // Stage 1
    {
      $group: { 
          "_id" : {
              "measure" : "$opencoremeasurement", 
              "leg" : "$opencoreleg"
          }, 
          refdata : {"$push" : {url: "$url", lat: "$spatial.geo.latitude", long : "$spatial.geo.longitude"}}
      }
    },

    // Stage 2
    {
      $project: { 
          "_id" : 0, 
          "measure" : "$_id.measure", 
          "leg" : "$_id.leg", 
          "refdata" : "$refdata"
      }
    },

    // Stage 3
    {
      $out: "aggregation_janusURLSet"
    }
  ],

  // Options
  {
    cursor: {
      batchSize: 50
    },

    allowDiskUse: true
  }

  // Created with 3T MongoChef, the GUI for MongoDB - http://3t.io/mongochef

);
