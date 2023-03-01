# Object Walker

## About

This is the start of a re-implementation of the "valutwalker" code made for CSDCO 
into a generic tool for taking file systems or object stores and building ot the 
associated initial Schema.org based data graph for them.

ObjectWalker is a tool for a group to use that wants to sync a subset of existing objects
to a set for publishing following structured data on the web approaches.   It is an exploration
environment for the publishing side of the structurd data on the web pattern.

It follows the RDA inspired Digital Object pattern and is designed from the start
to work with FAIR Digital Object patterns.

## Next steps

Review KV vs S3Select patterns for improved performance in source to sink synchronization. 

## Goals

- [ ] heuristics as configuration via viper
- [ ] source and sink as configuration 
  - [ ] Filesystem to S3 first to support CSDCO
  - [ ] S3 to S3 second to support UFOKN  (review old hydrocode for dams)
- [ ] Build out the DO + DOmeta elements from a source to a sink 
  - should this follow SHA hash naming patterns only?  
  - Can this be configured?
- [ ] Build out the sitegraph pattern with this tool (is this the right place for it?)
- [ ] syndication approaches (sitemap.xml, RSS, others)
  - [  ] https://w3id.org/tree/specification
  - [  ] LDES
- [ ] Sync patterns ala restic?




