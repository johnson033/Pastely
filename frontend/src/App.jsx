import { useState, useEffect } from 'react'
import styles from './App.module.css'
import { EventsOn } from '../wailsjs/runtime/runtime'
import { ListItems } from '../wailsjs/go/bindings/ClipboardItem'

function App() {
  const [count, setCount] = useState(0)
  const [items, setItems] = useState([]); 

  useEffect(() => {
    const fetchItems = async () => {
      const itemsList = await ListItems(0, 50); 
      setItems(itemsList);
    }
    fetchItems(); 
  }, [])

  useEffect(() => {
    const unsubscribeItemCreated = EventsOn('clipboard.item.created', (item) => {
      console.log('New clipboard item created:', item);
      setItems((prevItems) => [item, ...prevItems]); // Add new item to the top of the list
    }); 

    const ubsubscribeItemUpdated = EventsOn('clipboard.item.updated', (item) => {
      console.log('Clipboard item updated:', item);
      setItems((prevItems) => {
        // Update the item in the list
        return prevItems.map((prevItem) => {
          if (prevItem.id === item.id) {
            return item; // Replace the old item with the updated one
          }
          return prevItem; // Keep the old item
        });
      });
    });

    return () => {
      unsubscribeItemCreated(); // Clean up the event listener on component unmount
      ubsubscribeItemUpdated(); // Clean up the event listener on component unmount
    }
  }, [])
  

  return (
    <>
      <div
        style={{
          display: 'flex', 
          flexWrap: 'wrap', 
          gap: '10px',
        }}
      >
        {items.map((item, index) => {
          return (
            <div class={styles.Item} onClick={() => {
              // Set the clipboard content to the item's content
              navigator.clipboard.writeText(item.content);
            }}>
              {item.type}
              {item.content}
              {item.times_used}
            </div>
          )
        })}
      </div>
    </>
  )
}

export default App
