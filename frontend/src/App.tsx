import * as React from 'react'
import { useState, useEffect } from 'react'
import './App.css'

interface GroceryItem {
  id: number
  name: string
  completed: boolean
  created_at: string
}

function App() {
  const [items, setItems] = useState<GroceryItem[]>([])
  const [newItemName, setNewItemName] = useState('')
  const [loading, setLoading] = useState(true)
  const [error, setError] = useState<string | null>(null)

  useEffect(() => {
    let mounted = true

    const loadItems = async () => {
      try {
        setLoading(true)
        setError(null)
        const response = await fetch('/api/items')
        
        if (!mounted) return
        
        if (!response.ok) {
          setError(`Failed to fetch items: ${response.status} ${response.statusText}`)
          return
        }
        
        const data = await response.json()
        if (!mounted) return
        setItems(data)
      } catch (error) {
        if (!mounted) return
        const errorMessage = error instanceof Error ? error.message : 'An error occurred while fetching items'
        setError(errorMessage)
        console.error('Error fetching items:', error)
      } finally {
        if (mounted) {
          setLoading(false)
        }
      }
    }

    loadItems()

    return () => {
      mounted = false
    }
  }, [])

  const addItem = async (e: React.FormEvent) => {
    e.preventDefault()
    if (!newItemName.trim()) return

    try {
      const response = await fetch('/api/items', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          name: newItemName.trim(),
          completed: false,
        }),
      })

      if (!response.ok) {
        throw new Error(`Failed to add item: ${response.status} ${response.statusText}`)
      }

      const newItem = await response.json()
      setItems([newItem, ...items])
      setNewItemName('')
    } catch (error) {
      console.error('Error adding item:', error)
    }
  }

  const toggleItem = async (id: number) => {
    const item = items.find(i => i.id === id)
    if (!item) return

    try {
      const response = await fetch(`/api/items/${id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          ...item,
          completed: !item.completed,
        }),
      })

      if (!response.ok) {
        throw new Error(`Failed to toggle item: ${response.status} ${response.statusText}`)
      }

      setItems(items.map(item =>
        item.id === id ? { ...item, completed: !item.completed } : item
      ))
    } catch (error) {
      console.error('Error updating item:', error)
    }
  }

  const removeItem = async (id: number) => {
    try {
      const response = await fetch(`/api/items/${id}`, {
        method: 'DELETE',
      })

      if (!response.ok) {
        throw new Error(`Failed to remove item: ${response.status} ${response.statusText}`)
      }

      setItems(items.filter(item => item.id !== id))
    } catch (error) {
      console.error('Error removing item:', error)
    }
  }

  if (loading) {
    return <div className="container"><div className="content">Loading...</div></div>
  }

  if (error) {
    return <div className="container"><div className="content">Error: {error}</div></div>
  }

  return (
    <div className="container">
      <div className="content">
        <h1>Grocery List</h1>
        
        <form onSubmit={addItem} className="add-item-form">
          <input
            type="text"
            value={newItemName}
            onChange={(e) => setNewItemName(e.target.value)}
            placeholder="Add new item"
            autoFocus
          />
          <button type="submit">Add</button>
        </form>

        <ul className="grocery-list">
          {items.length === 0 ? (
            <li className="empty-list">No items in the list</li>
          ) : (
            items.map(item => (
              <li key={item.id} className={item.completed ? 'completed' : ''}>
                <input
                  type="checkbox"
                  checked={item.completed}
                  onChange={() => toggleItem(item.id)}
                />
                <span className="item-name">{item.name}</span>
                <button 
                  onClick={() => removeItem(item.id)} 
                  className="remove-btn"
                  type="button"
                >
                  ✕
                </button>
              </li>
            ))
          )}
        </ul>
      </div>
    </div>
  )
}

export default App
