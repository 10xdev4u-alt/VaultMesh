import React, { useEffect, useRef } from 'react'
import * as d3 from 'd3'

const NetworkGraph = () => {
  const svgRef = useRef<SVGSVGElement>(null)

  useEffect(() => {
    if (!svgRef.current) return

    const svg = d3.select(svgRef.current)
    svg.selectAll('*').remove()

    // Mock data for nodes and links
    const nodes = [
      { id: 'you', color: '#7C3AED' },
      { id: 'p1', color: '#06B6D4' },
      { id: 'p2', color: '#06B6D4' },
      { id: 'p3', color: '#06B6D4' },
    ]

    svg.selectAll('circle')
      .data(nodes)
      .enter()
      .append('circle')
      .attr('cx', (_, i) => 50 + i * 60)
      .attr('cy', 100)
      .attr('r', 12)
      .attr('fill', d => d.color)
  }, [])

  return (
    <div className="glass rounded-2xl p-6 mt-8 h-64 overflow-hidden">
      <h2 className="text-xl font-semibold mb-4 text-white">Network Mesh</h2>
      <svg ref={svgRef} className="w-full h-full" />
    </div>
  )
}

export default NetworkGraph
