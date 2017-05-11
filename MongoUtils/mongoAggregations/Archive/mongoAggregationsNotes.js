// http://stackoverflow.com/questions/22932364/mongodb-group-values-by-multiple-fields


{
   _id : "$opencoremeasurement", legs: { $addToSet: "$opencoreleg" } , legcount: {$sum:1}
}


// this is one (half) of what I want
{
   _id :{"measure": "$opencoremeasurement", "leg":"$opencoreleg"}, legcount: {$sum:1}
}


// collect the URL's for the L+M=datasets landing page
{
   _id :{"measure": "$opencoremeasurement", "leg":"$opencoreleg"}, urls: { $addToSet: "$url" }
}