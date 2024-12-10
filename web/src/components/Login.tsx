import { useState } from 'react'

import { useNavigate } from "react-router-dom"

import Box from '@mui/joy/Box'
import Button from '@mui/joy/Button'
import Input from '@mui/joy/Input'

import { useNotify } from './common'

import { ApiAuthCheck } from "../../wailsjs/go/ui/App"


interface AuthResponse {
  status: number,
  message: string,
}


export default function Login() {
  const navigate = useNavigate()
  const { notify, Notifybar } = useNotify()

  const [apiKey, setApiKey] = useState<string>("")

  const ApiKeyChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setApiKey(e.target.value)
  }

  const EnterLogin = (e: React.KeyboardEvent<HTMLInputElement>) => {
    e.key === "Enter" && Login()
  }

  const Login = async () => {
    const response = await ApiAuthCheck(apiKey)
    if (response.status === 1) {
      localStorage.setItem("apikey", apiKey)
      navigate("/index", { replace: true })
    } else {
      notify("认证失败", "danger")
    }
  }

  return (
    <Box sx={{ height: "100vh" }}>
      {Notifybar}
      <Box sx={{
        maxWidth: 320,
        m: "0 auto",
        position: "relative",
        top: 200,
      }}>
        <Input
          placeholder="API Key"
          type="password"
          onChange={ApiKeyChange}
          onKeyUp={EnterLogin}
          endDecorator={
            <Button variant="solid" onClick={Login}>
              确定
            </Button>
          }
        />
      </Box>
    </Box>
  )
}