package lifting

type structureConfig struct {
	salt       int32
	regionSize int8
	chunkRange int8
}

//go:generate stringer -type=Structure
type Structure int

const (
	BastionRemnant Structure = iota
	DesertPyramid
	Igloo
	JungleTemple
	OceanRuin
	PillagerOutpost
	RuinedPortalOverworld
	RuinedPortalNether
	Shipwreck
	SwampHut
	Village
	JunglePyramid = JungleTemple
)

var (
	Structures = []Structure{
		BastionRemnant,
		DesertPyramid,
		Igloo,
		JungleTemple,
		OceanRuin,
		PillagerOutpost,
		RuinedPortalOverworld,
		RuinedPortalNether,
		Shipwreck,
		SwampHut,
		Village,
	}
)

func getStructureConfig(structure Structure, version Version) *structureConfig {
	var (
		s_BastionRemnant_16     = inlineConfig(30, 4, 30084232)
		s_BastionRemnant_16_1   = inlineConfig(27, 4, 30084232)
		s_DesertPyramid         = inlineConfig(32, 8, 14357617)
		s_Igloo_9               = inlineConfig(32, 8, 14357617)
		s_Igloo_13              = inlineConfig(32, 8, 14357618)
		s_JungleTemple_8        = inlineConfig(32, 8, 14357617)
		s_JungleTemple_13       = inlineConfig(32, 8, 14357619)
		s_OceanRuin_13          = inlineConfig(16, 8, 14357621)
		s_OceanRuin_16          = inlineConfig(20, 8, 14357621)
		s_PillagerOutpost       = inlineConfig(32, 8, 165745296)
		s_RuinedPortalOverworld = inlineConfig(40, 15, 34222645)
		s_RuinedPortalNether    = inlineConfig(25, 15, 34222645)
		s_Shipwreck_13          = inlineConfig(15, 8, 165745295)
		s_Shipwreck_13_1        = inlineConfig(16, 8, 165745295)
		s_Shipwreck_16          = inlineConfig(24, 4, 165745295)
		s_SwampHut_8            = inlineConfig(32, 8, 14357617)
		s_SwampHut_13           = inlineConfig(32, 8, 14357620)
		s_Village               = inlineConfig(32, 8, 10387312)
	)
	switch structure {
	case BastionRemnant:
		if version == MC_1_16 {
			return &s_BastionRemnant_16
		} else if version > MC_1_16 {
			return &s_BastionRemnant_16_1
		}
	case DesertPyramid:
		if version >= MC_1_8 {
			return &s_DesertPyramid
		}
	case Igloo:
		if version >= MC_1_13 {
			return &s_Igloo_13
		} else if version >= MC_1_9 {
			return &s_Igloo_9
		}
	case JungleTemple:
		if version >= MC_1_13 {
			return &s_JungleTemple_13
		} else if version >= MC_1_8 {
			return &s_JungleTemple_8
		}
	case OceanRuin:
		if version >= MC_1_16 {
			return &s_OceanRuin_16
		} else if version >= MC_1_13 {
			return &s_OceanRuin_13
		}
	case PillagerOutpost:
		if version >= MC_1_14 {
			return &s_PillagerOutpost
		}
	case RuinedPortalOverworld:
		if version >= MC_1_16 {
			return &s_RuinedPortalOverworld
		}
	case RuinedPortalNether:
		if version >= MC_1_16 {
			return &s_RuinedPortalNether
		}
	case Shipwreck:
		if version >= MC_1_16 {
			return &s_Shipwreck_16
		} else if version >= MC_1_13_1 {
			return &s_Shipwreck_13_1
		} else if version == MC_1_13 {
			return &s_Shipwreck_13
		}
	case SwampHut:
		if version >= MC_1_13 {
			return &s_SwampHut_13
		} else if version >= MC_1_8 {
			return &s_SwampHut_8
		}
	case Village:
		if version >= MC_1_8 {
			return &s_Village
		}
	}
	return nil
}

func inlineConfig(spacing, separation, salt int) structureConfig {
	return structureConfig{
		salt:       int32(salt),
		regionSize: int8(spacing),
		chunkRange: int8(spacing) - int8(separation),
	}
}
