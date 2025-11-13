function generateColorFromString(str) {
  if (!str) return '#C6E5FF'
  
  let hash = 0
  for (let i = 0; i < str.length; i++) {
    hash = str.charCodeAt(i) + ((hash << 5) - hash)
  }
  
  const hue = Math.abs(hash % 360)
  const saturation = 60 + (Math.abs(hash) % 20)
  const lightness = 50 + (Math.abs(hash) % 20)
  
  return `hsl(${hue}, ${saturation}%, ${lightness}%)`
}

export function transformGraphData(backendData) {
  if (!backendData) {
    return { nodes: [], edges: [] }
  }

  const nodes = Array.isArray(backendData.nodes) ? backendData.nodes : []
  const edges = Array.isArray(backendData.edges) ? backendData.edges : []
  
  const typeColorMap = {}
  nodes.forEach(node => {
    const entityType = node.entity_type || ''
    if (entityType && !typeColorMap[entityType]) {
      typeColorMap[entityType] = generateColorFromString(entityType)
    }
  })
  
  const nodeIdMap = new Map()

  const transformedNodes = nodes.map((node, index) => {
    const nodeId = node.entity_name || `node_${index}`
    const nodeLabel = node.entity_name || `node_${index}`
    const nodeSize = node.pagerank ? Math.max(15, Math.min(30, node.pagerank * 100)) : 20
    const entityType = node.entity_type || ''
    const nodeColor = typeColorMap[entityType] || '#C6E5FF'

    if (node && node.entity_name) {
      nodeIdMap.set(node.entity_name, nodeId)
    }
    if (node && node.entity_id) {
      nodeIdMap.set(String(node.entity_id), nodeId)
    }
    if (node && node.id) {
      nodeIdMap.set(String(node.id), nodeId)
    }

    return {
      ...node,
      id: nodeId,
      label: nodeLabel,
      originalLabel: nodeLabel,
      type: 'circle',
      size: nodeSize,
      style: {
        fill: nodeColor
      }
    }
  })

  const transformedEdges = edges.map((edge, index) => {
    const edgeId = `e${index}`

    const source =
      nodeIdMap.get(edge && edge.source_entity) ||
      nodeIdMap.get(edge && edge.source) ||
      nodeIdMap.get(
        edge && edge.source_id ? String(edge.source_id) : undefined
      ) ||
      (edge && edge.source_entity) ||
      (edge && edge.source) ||
      (edge && edge.source_id ? String(edge.source_id) : undefined) ||
      `source_${index}`

    const target =
      nodeIdMap.get(edge && edge.target_entity) ||
      nodeIdMap.get(edge && edge.target) ||
      nodeIdMap.get(
        edge && edge.target_id ? String(edge.target_id) : undefined
      ) ||
      (edge && edge.target_entity) ||
      (edge && edge.target) ||
      (edge && edge.target_id ? String(edge.target_id) : undefined) ||
      `target_${index}`

    return {
      id: edgeId,
      source,
      target,
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
