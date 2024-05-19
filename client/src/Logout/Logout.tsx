import React, { useEffect } from 'react'
import axios from 'axios'
import { useNavigate } from 'react-router-dom'

export default function Logout() {
    const navigate = useNavigate()

    useEffect(() => {
        axios.post('/logout').then(() => navigate('/'))
    },[])
    
  return (
    <div>Logout</div>
  )
}
