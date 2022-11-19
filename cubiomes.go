package main

// #include "finders.h"
import "C"

type CB_StructureConfig C.struct_StructureConfig

const (
	CB_Feature = iota
	CB_Desert_Pyramid
	CB_Jungle_Temple
	CB_Swamp_Hut
	CB_Igloo
	CB_Village
	CB_Ocean_Ruin
	CB_Shipwreck
	CB_Monument
	CB_Mansion
	CB_Outpost
	CB_Ruined_Portal
	CB_Ruined_Portal_N
	CB_Ancient_City
	CB_Treasure
	CB_Mineshaft
	CB_Fortress
	CB_Bastion
	CB_End_City
	CB_End_Gateway
	CB_FEATURE_NUM
	CB_Jungle_Pyramid = CB_Jungle_Temple
)

const (
	CB_MC_1_0 = iota
	CB_MC_1_1
	CB_MC_1_2
	CB_MC_1_3
	CB_MC_1_4
	CB_MC_1_5
	CB_MC_1_6
	CB_MC_1_7
	CB_MC_1_8
	CB_MC_1_9
	CB_MC_1_10
	CB_MC_1_11
	CB_MC_1_12
	CB_MC_1_13
	CB_MC_1_14
	CB_MC_1_15
	CB_MC_1_16
	CB_MC_1_17
	CB_MC_1_18
	CB_MC_1_19
	CB_MC_NEWEST = CB_MC_1_19
)
