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
          "legcount" : {
              "$sum" : 1
          }
      }
    },

    // Stage 2
    {
      $project: { 
          "_id" : 0, 
          "measure" : "$_id.measure", 
          "leg" : "$_id.leg", 
          "count" : "$legcount"
      }
    },

    // Stage 3
    {
      $out: "aggregation_janusMLCountv2"
      
    }

  ]

  // Created with 3T MongoChef, the GUI for MongoDB - http://3t.io/mongochef

);
