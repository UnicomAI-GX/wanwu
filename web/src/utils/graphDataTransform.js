export function transformGraphData(backendData, options = {}) {
  if (!backendData) {
    return { nodes: [], edges: [] }
  }

  const nodes = backendData.nodes || []
  const edges = backendData.edges || []
  
  const {
    getNodeId = (node, index) => node.entity_name || `n${index}`,
    getNodeLabel = (node, index) => node.entity_name || '',
    getNodeSize = (node, index) => node.pagerank ? Math.max(15, Math.min(30, node.pagerank * 100)) : 20,
    getNodeColor = (node, index) => undefined,
    getEdgeId = (edge, index) => `e${index}`,
    getEdgeLabel = (edge, index) => edge.description || ''
  } = options

  const transformedNodes = nodes.map((node, index) => {
    const nodeId = getNodeId(node, index)
    const nodeLabel = getNodeLabel(node, index)
    const nodeSize = getNodeSize(node, index)
    const nodeColor = getNodeColor(node, index)

    const nodeStyle = {}
    if (nodeColor !== undefined) {
      nodeStyle.fill = nodeColor
    }

    return {
      id: nodeId,
      label: nodeLabel,
      type: 'circle',
      size: nodeSize,
      ...(Object.keys(nodeStyle).length > 0 && { style: nodeStyle }),
      ...node
    }
  })

  const transformedEdges = edges.map((edge, index) => {
    const edgeId = getEdgeId(edge, index)
    const edgeLabel = getEdgeLabel(edge, index)

    return {
      id: edgeId,
      source: edge.source_entity,
      target: edge.target_entity,
      label: edgeLabel,
      ...(edge.weight && {
        style: {
          lineWidth: Math.max(1, Math.min(5, edge.weight / 2))
        }
      }),
      ...edge
    }
  })

  return {
    nodes: transformedNodes,
    edges: transformedEdges
  }
}

export default {
  transformGraphData
}
