import axios from 'axios'
import React, { useEffect } from 'react'

export default function Account() {
    useEffect(() => {
        axios.get('/me').then(({data}) => console.log(data))
    }, [])

  return (
    <div>
        Users


    </div>
  )
}
