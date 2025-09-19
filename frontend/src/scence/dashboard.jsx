import React, { useState, useEffect } from 'react';
import { Container, Typography, Box, Grid, Card, CardContent, TextField, Button, createTheme, ThemeProvider } from '@mui/material';
import axios from 'axios';
import { jwtDecode } from 'jwt-decode';
import { LineChart, Line, XAxis, YAxis, CartesianGrid, Tooltip, Legend, ResponsiveContainer } from 'recharts';
import { useNavigate } from 'react-router-dom';

const theme = createTheme({
    palette: {
        background: {
            default: '#f0f0f0',
        },
    },
});

const Dashboard = () => {
    const [country, setCountry] = useState('');
    const [city, setCity] = useState('');
    const [selectedDate, setSelectedDate] = useState(new Date());
    const [weatherData, setWeatherData] = useState({
        day: {
            datetime: "",
            tempmin: 0,
            tempmax: 0,
            humidity: 0,
            precip: 0,
            windspeed: 0
        },
        hourly: []
    });

    const navigate = useNavigate();

    useEffect(() => {
        const token = localStorage.getItem('auth');
        if (token) {
            try {
                const decodedToken = jwtDecode(token);
                setCountry(decodedToken.country || '');
                setCity(decodedToken.city || '');
            } catch (error) {
                console.error('Error decoding token:', error);
            }
        }  
    }, []);

    const handleCountryChange = (event) => {
        setCountry(event.target.value);
    };

    const handleCityChange = (event) => {
        setCity(event.target.value);
    };

    const handleDateChange = (date) => {
        setSelectedDate(date);
        fetchWeatherData();
    };

    const fetchWeatherData = async () => {
        try {
            const response = await axios.post(
                'http://localhost:8080/api/weather',
                {
                    location: `${city},${country}`,
                    datetime: selectedDate.toISOString().split('T')[0]
                },
                {
                    headers: {
                        'Content-Type': 'application/json',
                        'Authorization': `Bearer ${localStorage.getItem('auth')}`
                    }
                }
            );
            setWeatherData(response.data);
        } catch (error) {
            console.error('Error fetching weather data:', error);
        }
    };

    const handleLogout = () => {
        localStorage.removeItem('auth');
        navigate('/');
    };

    return (
        <ThemeProvider theme={theme}>
            <Box sx={{ height: '100vh', display: 'flex', flexDirection: 'column', backgroundColor: theme.palette.background.default }}>
                <Box sx={{ padding: '20px', flex: '0 1 auto', boxShadow: '0 4px 8px rgba(0,0,0,0.1)' }}>
                    <Typography variant="h4" component="h1" gutterBottom>
                        Weather Dashboard
                    </Typography>
                    <Button variant="outlined" color="secondary" onClick={handleLogout} sx={{ float: 'right' }}>
                        Logout
                    </Button>
                </Box>
                <Grid container sx={{ flex: '1 1 auto' }}>
                    <Grid item xs={12} md={3} sx={{ padding: '20px' }}>
                        <Card sx={{ boxShadow: '0 4px 8px rgba(0,0,0,0.1)' }}>
                            <CardContent>
                                <Box sx={{ display: 'flex', flexDirection: 'column', gap: '20px' }}>
                                    <TextField
                                        label="Country"
                                        value={country}
                                        onChange={handleCountryChange}
                                        sx={{ width: '100%' }}
                                        variant="outlined"
                                    />
                                    <TextField
                                        label="City"
                                        value={city}
                                        onChange={handleCityChange}
                                        sx={{ width: '100%' }}
                                        variant="outlined"
                                    />
                                    <TextField
                                        id="date"
                                        label="Date"
                                        type="date"
                                        value={selectedDate.toISOString().split('T')[0]}
                                        onChange={(e) => handleDateChange(new Date(e.target.value))}
                                        InputLabelProps={{
                                            shrink: true,
                                        }}
                                        sx={{ width: '100%' }}
                                        variant="outlined"
                                    />
                                    <Button variant="contained" color="primary" onClick={fetchWeatherData} sx={{ width: '100%' }}>
                                        Fetch Weather
                                    </Button>
                                </Box>
                            </CardContent>
                        </Card>
                    </Grid>
                    <Grid item xs={12} md={9} sx={{ padding: '20px' }}>
                        <Container maxWidth="lg" sx={{ paddingTop: '20px', backgroundColor: theme.palette.background.default, boxShadow: '0 4px 8px rgba(0,0,0,0.1)' }}>
                            <Typography variant="h6" component="h2" gutterBottom>
                                Daily Summary
                            </Typography>
                            <Grid container spacing={3} sx={{ marginBottom: '20px' }}>
                                <Grid item xs={12}>
                                    <Card sx={{ boxShadow: '0 4px 8px rgba(0,0,0,0.1)' }}>
                                        <CardContent>
                                            <Typography variant="body1">Date: {weatherData.day.datetime}</Typography>
                                            <Typography variant="body1">Min Temp: {weatherData.day.tempmin}°F</Typography>
                                            <Typography variant="body1">Max Temp: {weatherData.day.tempmax}°F</Typography>
                                            <Typography variant="body1">Humidity: {weatherData.day.humidity}%</Typography>
                                            <Typography variant="body1">Precipitation: {weatherData.day.precip} inches</Typography>
                                            <Typography variant="body1">Wind Speed: {weatherData.day.windspeed} mph</Typography>
                                        </CardContent>
                                    </Card>
                                </Grid>
                            </Grid>
                            <Typography variant="h6" component="h2" gutterBottom>
                                Hourly Data
                            </Typography>
                            <ResponsiveContainer width="100%" height={400}>
                                <LineChart
                                    data={weatherData.hourly}
                                    margin={{ top: 20, right: 30, left: 20, bottom: 5 }}
                                >
                                    <CartesianGrid strokeDasharray="3 3" />
                                    <XAxis dataKey="datetime" />
                                    <YAxis />
                                    <Tooltip />
                                    <Legend />
                                    <Line type="monotone" dataKey="temp" stroke="#8884d8" activeDot={{ r: 8 }} />
                                    <Line type="monotone" dataKey="humidity" stroke="#82ca9d" />
                                    <Line type="monotone" dataKey="windspeed" stroke="#ffc658" />
                                </LineChart>
                            </ResponsiveContainer>
                        </Container>
                    </Grid>
                </Grid>
            </Box>
        </ThemeProvider>
    );
};

export default Dashboard;
