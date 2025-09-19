import { useState } from 'react'
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom'
import './App.css'
import  Login from '../src/scence/Login'
import Dashboard from './scence/dashboard'


function App() {

  return (
   
      <Routes>
        <Route path="/" element={<Login />} />
        <Route path='/dashboard' element={<Dashboard />} />
      </Routes>
  )
}

export default App
