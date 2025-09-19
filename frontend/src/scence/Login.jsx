import React, { useState, useEffect } from 'react';
import { TextField, Button, Container, Typography, Box } from '@mui/material';
import axios from 'axios';
import { ToastContainer, toast } from 'react-toastify';
import 'react-toastify/dist/ReactToastify.css';
import { useNavigate } from 'react-router-dom';

const Login = () => {
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  const navigate = useNavigate();

  useEffect(() => {
    // Check if auth token exists in localStorage
    const token = localStorage.getItem('auth');
    if (token) {
      // Redirect to dashboard if token exists
      navigate('/dashboard');
    }
  }, [navigate]);

  const handleLogin = async (event) => {
    event.preventDefault();
    const loginData = {
      email: email,
      password: password
    };
    try {
      const response = await axios.post('http://localhost:8080/api/user/login', loginData, {
        headers: {
          'Content-Type': 'application/json'
        }
      });
      if (response.status === 200) {
        // Extract token from response data
        const token = response.data.token;
  
        // Store token in localStorage
        localStorage.setItem('auth', token);
  
        toast.success('Login successful!');
        setTimeout(() => {
          navigate('/dashboard');
        }, 2000); // Delay to allow the toast message to be seen
      } else {
        toast.error(response.data.error || 'Login failed!');
      }
    } catch (error) {
      toast.error(error.response ? error.response.data.error : 'Error occurred');
    }
  };
  
  return (
    <Box
      sx={{
        width: "100vw",
        padding: "0px",
        position: "fixed",
        top: "0px",
        left: "0px",
        margin: "0px",
        height: "100vh",
        backgroundImage: "url('https://images.unsplash.com/photo-1601134467661-3d775b999c8b?q=80&w=2575&auto=format&fit=crop&ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D')",
      }}
    >
      <Container maxWidth="sm" sx={{ background: "#fff", padding: "80px", paddingTop: "0px" }}>
        <Box
          sx={{
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
            marginTop: "35vh",
          }}
        >
          <Typography sx={{ marginTop: "30px" }} variant="h4" component="h1" gutterBottom>
            Login
          </Typography>
          <Box
            component="form"
            onSubmit={handleLogin}
            sx={{
              mt: 1,
              width: '100%',
            }}
          >
            <TextField
              margin="normal"
              required
              type='email'
              fullWidth
              id="email"
              label="email"
              name="email"
              autoComplete="email"
              autoFocus
              value={email}
              onChange={(e) => setEmail(e.target.value)}
            />
            <TextField
              margin="normal"
              required
              fullWidth
              name="password"
              label="Password"
              type="password"
              id="password"
              autoComplete="current-password"
              value={password}
              onChange={(e) => setPassword(e.target.value)}
            />
            <Button
              type="submit"
              fullWidth
              variant="contained"
              color="primary"
              sx={{ mt: 3, mb: 2 }}
            >
              Login
            </Button>
          </Box>
        </Box>
      </Container>
      <ToastContainer />
    </Box>
  );
};

export default Login;
