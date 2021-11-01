# Object Walker

## About

This is the start of a re-implementation of the "valutwalker" code made for CSDCO 
into a generic tool for taking file systems or object stores and building ot the 
associated initial Schema.org based data graph for them.

It follows the RDA inspired Digital Object pattern and is designed from the start
to work with FAIR Digital Object patterns.

## Next steps

Review KV vs S3Select patterns for improved performance in source to sink synchronization. 

## Goals

- [ ] Build out the DO + DOmeta elements from a source to a sink 
  - should this follow SHA hash naming patterns only?  
  - Can this be configured?
- [ ] Build out the sitegraph pattern with this tool (is this the right place for it?)
- [ ] syndication approaches (sitemap.xml, RSS, others)
- [ ] Sync patterns ala restic?




