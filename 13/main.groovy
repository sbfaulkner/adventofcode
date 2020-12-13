import groovy.transform.SourceURI
import java.nio.file.Path
import java.nio.file.Paths

@SourceURI
final URI sourceUri
final Path input = Paths.get(sourceUri).resolveSibling("input")

Track track = new Track(input)

while (track.isMultipleCarts()) {
  // track.print()
  track.tick()
}

Cart c = track.carts.first()

println "Last cart is at ${c.x},${c.y}"
