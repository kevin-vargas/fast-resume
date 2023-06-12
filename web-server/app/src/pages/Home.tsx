import { Box, Button, Container, Typography, useTheme } from "@mui/material";
import { useEffect } from "react";
import Config from "../config"
import Cookies from 'js-cookie';
import { Navigate } from "react-router-dom";

const makeRef = (url: string) => {
  const from = encodeURI(window.location.href)
  const to = encodeURI(url)
  return `https://slack.com/oauth/v2/authorize?scope=&amp;user_scope=channels%3Aread%2Cchannels%3Ahistory%2Cusers%3Aread&amp;redirect_uri=${to}&amp;state=${from}&amp;client_id=5388910517927.5416109352385`
}
const slackButton = (href: string) => (dark: boolean) => `
<a
        href="${href}" 
        style="${
            dark ? 
            "align-items:center;color:#fff;background-color:#4A154B;border:0;border-radius:4px;display:inline-flex;font-family:Lato" 
            : 
            "align-items:center;color:#000;background-color:#fff;border:1px solid #ddd;border-radius:4px;display:inline-flex;font-family:Lato"
        }, sans-serif;font-size:18px;font-weight:600;height:56px;justify-content:center;text-decoration:none;width:276px">
            <svg 
            xmlns="http://www.w3.org/2000/svg" 
            style="height:24px;width:24px;margin-right:12px" 
            viewBox="0 0 122.8 122.8">
                <path 
                d="M25.8 77.6c0 7.1-5.8 12.9-12.9 12.9S0 84.7 0 77.6s5.8-12.9 12.9-12.9h12.9v12.9zm6.5 0c0-7.1 5.8-12.9 12.9-12.9s12.9 5.8 12.9 12.9v32.3c0 7.1-5.8 12.9-12.9 12.9s-12.9-5.8-12.9-12.9V77.6z" 
                fill="#e01e5a">
                </path>
                <path 
                d="M45.2 25.8c-7.1 0-12.9-5.8-12.9-12.9S38.1 0 45.2 0s12.9 5.8 12.9 12.9v12.9H45.2zm0 6.5c7.1 0 12.9 5.8 12.9 12.9s-5.8 12.9-12.9 12.9H12.9C5.8 58.1 0 52.3 0 45.2s5.8-12.9 12.9-12.9h32.3z" 
                fill="#36c5f0">
                </path>
                <path 
                d="M97 45.2c0-7.1 5.8-12.9 12.9-12.9s12.9 5.8 12.9 12.9-5.8 12.9-12.9 12.9H97V45.2zm-6.5 0c0 7.1-5.8 12.9-12.9 12.9s-12.9-5.8-12.9-12.9V12.9C64.7 5.8 70.5 0 77.6 0s12.9 5.8 12.9 12.9v32.3z" 
                fill="#2eb67d">
                </path>
                <path 
                d="M77.6 97c7.1 0 12.9 5.8 12.9 12.9s-5.8 12.9-12.9 12.9-12.9-5.8-12.9-12.9V97h12.9zm0-6.5c-7.1 0-12.9-5.8-12.9-12.9s5.8-12.9 12.9-12.9h32.3c7.1 0 12.9 5.8 12.9 12.9s-5.8 12.9-12.9 12.9H77.6z" 
                fill="#ecb22e">
                </path>
            </svg>
            Add to Slack
            </a>
`

export default function HomePage() {
    const theme = useTheme();
    const isLogged = Cookies.get("logged")
      if (isLogged && isLogged == "true" ) {
        return <Navigate to="/channels" />
      }
    useEffect(() => {
        const ele = document.getElementById("slackconnect")
        if (ele){
            const ref = makeRef(Config.authorizeUri)
            ele.innerHTML = slackButton(ref)(theme.palette.mode == "dark")
        }        
    },[])

    return (
        <Container
        maxWidth="lg"
        sx={{
          display: "flex",
          alignItems: "center",
          justifyContent: "center",
          flexDirection: "column",
          height: "100vh",
        }}
      >
        <Box sx={{ mb: 5, mt: -10 }}>
        </Box>
        <Typography
          sx={{
            textAlign: "center",
            marginTop: "-4rem",
            fontSize: "5rem",
            fontWeight: 700,
            letterSpacing: "-0.5rem",
            display: "inline-block",
            whiteSpace: "nowrap",
            [theme.breakpoints.down("sm")]: {
              fontSize: "4rem",
              letterSpacing: "-0.4rem",
            },
          }}
          gutterBottom
        >
          Fast Resume
        </Typography>
  
        <div id="slackconnect">
        </div>
      </Container>
    );
  }