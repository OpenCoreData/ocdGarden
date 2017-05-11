db.jsonld.aggregate(

  // Pipeline
  [
    // Stage 1
    {
      $unwind: "$tables"
    },

    // Stage 2
    {
      $project: {
        "url": "$tables.url",
        rowcount: {$size: "$tables.row"}
      }
    },

    // Stage 3
    {
      $out: "aggregation_rowCount"
    }
  ],

  // Options
  {
    cursor: {
      batchSize: 50
    }
  }

  // Created with 3T MongoChef, the GUI for MongoDB - http://3t.io/mongochef

);
