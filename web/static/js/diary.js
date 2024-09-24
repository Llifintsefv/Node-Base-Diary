let cy

document.addEventListener('DOMContentLoaded', function () {
	cy = cytoscape({
		container: document.getElementById('diary-canvas'),
		elements: [],
		style: [
			{
				selector: 'node',
				style: {
					'background-color': '#ff0000',
					'border-color': '#ff0000',
					'border-width': 3,

					label: 'data(title)',
				},
			},
			{
				selector: 'edge',
				style: {
					width: 3,
					'line-color': '#ccc',
					'target-arrow-color': '#ccc',
					'target-arrow-shape': 'triangle',
					'curve-style': 'bezier',
				},
			},
		],

		zoom: 1,

		pan: { x: 0, y: 0 },
		minZoom: 0.1,
		maxZoom: 10,
		wheelSensitivity: 0.2,
	})

	loadNodes()

	cy.on('tap', function (evt) {
		if (evt.target === cy) {
			const pos = evt.position
			createNode(pos.x, pos.y)
		}
	})

	cy.on('tap', 'node', function (evt) {
		const node = evt.target
		updateNodeContent(node)
	})

	cy.on('drag', 'node', function (evt) {
		const node = evt.target
		updateNodePosition(node)
	})
})

function loadNodes() {
	fetch('/api/nodes')
		.then(response => response.json())
		.then(nodes => {
			nodes.forEach(node => {
				console.log('Node X:', node.x) // Вывод значения X в консоль
				console.log('Node Y:', node.y) // Вывод значения Y в консоль
				cy.add({
					group: 'nodes',
					data: {
						id: node.ID.toString(),
						title: node.Title,
						content: node.Content,
					},
					position: { x: node.x, y: node.y },
				})
				if (node.ParentID) {
					cy.add({
						group: 'edges',
						data: {
							source: node.ParentID.toString(),
							target: node.ID.toString(),
						},
					})
				}
			})
			cy.layout({ name: 'preset' })
		})
}

function createNode(x, y) {
	const title = prompt('Введите заголовок ноды:')
	if (title) {
		fetch('/api/nodes', {
			method: 'POST',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ title, x, y }),
		})
			.then(response => response.json())
			.then(node => {
				cy.add({
					group: 'nodes',
					data: {
						id: node.ID.toString(),
						title: node.Title,
						content: node.Content,
						// Сохраняем координаты в data
						X: x,
						Y: y,
					},
					// Устанавливаем координаты непосредственно здесь
					position: { x: x, y: y },
				})
			})
	}
}

function updateNodeContent(node) {
	const newContent = prompt(
		'Введите новое содержимое ноды:',
		node.data('content')
	)
	if (newContent !== null) {
		fetch(`/api/nodes/${node.id()}`, {
			method: 'PUT',
			headers: { 'Content-Type': 'application/json' },
			body: JSON.stringify({ content: newContent }),
		})
			.then(response => response.json())
			.then(updatedNode => {
				node.data('content', updatedNode.Content)
			})
	}
}

function updateNodePosition(node) {
	const position = node.position()
	fetch(`/api/nodes/${node.id()}`, {
		method: 'PUT',
		headers: { 'Content-Type': 'application/json' },
		body: JSON.stringify({ x: position.x, y: position.y }),
	})
}
