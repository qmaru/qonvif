import { useState, useCallback } from "react"

import Snackbar from "@mui/joy/Snackbar"

type notifyLevel = "primary" | "neutral" | "danger" | "success" | "warning"

export const useNotify = () => {
    const [open, setOpen] = useState<boolean>(false)
    const [message, setMessage] = useState<string>("")
    const [color, setColor] = useState<notifyLevel>("primary")

    const notify = useCallback((msg: string, color: notifyLevel = "primary") => {
        setMessage(msg)
        setColor(color)
        setOpen(true)
    }, [])

    const handleClose = () => {
        setOpen(false)
    }

    const Notifybar = (
        <Snackbar
            anchorOrigin={{ vertical: "top", horizontal: "center" }}
            open={open}
            color={color}
            autoHideDuration={3000}
            onClose={handleClose}
            variant="solid"
        >
            {message}
        </Snackbar>
    )

    return { notify, Notifybar }
}