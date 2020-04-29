package engine

var (
	effectBlacklist = [...]int32{
		132, // skillEffect (converts skillpoints to skill level, but we modify skillLevel ourselves)
	}
)
