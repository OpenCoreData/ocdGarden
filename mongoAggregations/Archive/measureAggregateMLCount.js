db.schemaorg.aggregate(

  // Pipeline
  [
    // Stage 1
    {
      $group: { 
          "_id" : {"measure" : "$opencoremeasurement", 
              "leg" : "$opencoreleg"},
          
          "legcount" : {
              "$sum" : 1
          }
      }
    },

    // Stage 2
    {
      $out: "aggregation_janusMLCount"
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
