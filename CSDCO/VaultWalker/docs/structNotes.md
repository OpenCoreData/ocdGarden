# Data structure notes

## About
comparing DO parameters

### FDP resource
```json
{
    "name": "solar-system",
    "path": "http://example.com/solar-system.csv",
    "title": "The Solar System",
    "description": "My favourite data about the solar system.",
    "format": "csv",
    "mediatype": "text/csv",
    "encoding": "utf-8",
    "bytes": 1,
    "hash": "",
    "schema": "",
    "sources": "",
    "licenses": ""
}
```

### DID 
Note, DID may only have a single context https://w3c-ccg.github.io/did-spec/contexts/did-v1.jsonld
However, I don't completely follow that since one can do custom context.

```json
{
  "@context": "https://example.org/example-method/v1",
  "id": "did:example:123456789abcdefghi",
  "publicKey": [{ ...  }],
  "authentication": [{ ...  }],
  "service": [{ ...  }]
}
```

Reference https://w3c-ccg.github.io/did-spec/#example-8-various-service-endpoints for services.


### PID Kernel (RDA)
PID,Handle,1..n,
Global identifier for the object; external to the PID Kernel Information

KernelInformationProfile,Handle,1,
Handle to the Kernel Information type profile; serves as pointer to profile in DTR. Address of DTR federation expected to be global (common) knowledge.

digitalObjectType,Handle,1,
"Handle points to type definition in DTR for this type of object. Distinguishing metadata from data objects is a client decision within a particular usage context, which may to some extent rely on the digitalObjectType value provided."

digitalObjectLocation,URL,1..n,
Pointer to the content object location (pointer to the DO). This may be in addition to a pointer to a human-readable landing page for the object.

digitalObjectPolicy,Handle,1,
"Pointer to a policy object which specifies a model for managing changes to the object or its Kernel Information record, including particularly object access and modification policies. A caller should be able to determine the expected future changes to the object from the policy, which are based on managed processes the object owner maintains."

etag,Hex string,1,
Checksum of object contents. Checksum format determined via attribute type referenced in a Kernel Information record.

dateModified,ISO 8601 Date,0..1,
Last date/time of object modification. Mandatory if applicable.

dateCreated,ISO 8601 Date,1,
Date/time of object creation

version,String,0..1,
"If tracked, a version for the object, which must follow a total order. Mandatory for all objects with at least one predecessor version."

wasDerivedFrom,Handle,0..n,
"PROV-DM: Transformation of an entity into another, an update of an entity resulting in a new one, or the construction of a new entity based on a pre-existing entity."

specializationOf,Handle,0..n,
"PROV-DM: Entity is a specialization of another that shares all aspects of the latter, and additionally presents more specific aspects of the same thing as the latter."

wasRevisionOf,Handle,0..n,
PROV-DM: A derivation for which the resulting entity is a revised version of some original.

hadPrimarySource,Handle,0..n,
"PROV-DM: A primary source for a topic refers to something produced by some agent with direct experience and knowledge about the topic, at the time of the topic's study, without benefit from hindsight."

wasQuotedFrom,Handle,0..n,
"PROV-DM: Used for the repeat of (some or all of) an entity, such as text or image, by someone who may or may not be its original author."

alternateOf,Handle,0..n,
"PROV-DM: Entities present aspects of the same thing. These aspects may be the same or different, and the alternate entities may or may not overlap in time."







PID,Handle,1..n,
KernelInformationProfile,Handle,1,
digitalObjectType,Handle,1,
digitalObjectLocation,URL,1..n,
digitalObjectPolicy,Handle,1,
etag,Hex string,1,
dateModified,ISO 8601 Date,0..1,
dateCreated,ISO 8601 Date,1,
version,String,0..1,
wasDerivedFrom,Handle,0..n,
specializationOf,Handle,0..n,
wasRevisionOf,Handle,0..n,
hadPrimarySource,Handle,0..n,
wasQuotedFrom,Handle,0..n,
alternateOf,Handle,0..n,
