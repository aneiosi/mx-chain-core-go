package indexer

const txsBulkSizeThreshold = 5242880 // 5MBs

const maxNumberOfDocumentsGet = 1000
const txIndex = "transactions"
const blockIndex = "blocks"
const miniblocksIndex = "miniblocks"
const tpsIndex = "tps"
const validatorsIndex = "validators"
const roundIndex = "rounds"
const ratingIndex = "rating"

const metachainTpsDocID = "meta"
const shardTpsDocIDPrefix = "shard"
