import { useState, useEffect, useCallback, useTransition } from "react"
import { useNavigate } from "react-router-dom"

import Box from "@mui/joy/Box"
import Button from "@mui/joy/Button"
import Card from "@mui/joy/Card"
import CardContent from "@mui/joy/CardContent"
import Chip from "@mui/joy/Chip"
import Input from '@mui/joy/Input'
import IconButton from '@mui/joy/IconButton'
import Modal from "@mui/joy/Modal"
import ModalDialog from "@mui/joy/ModalDialog"
import ModalClose from "@mui/joy/ModalClose"
import DialogTitle from "@mui/joy/DialogTitle"
import DialogContent from "@mui/joy/DialogContent"
import Typography from "@mui/joy/Typography"
import Stack from "@mui/joy/Stack"

import { useNotify } from "./common"
import { LeftArrowIcon, RightArrowIcon, UpArrowIcon, DownArrowIcon } from "./icons/svgs"

import { player, ui } from "../../wailsjs/go/models"
import { ApiAuthCheck, ApiOnvifDevices, ApiOnvifDeviceProfile, ApiOnvifDeviceStreamurl, ApiOnvifPlay, ApiOnvifDevicePtzStatus, ApiOnvifDevicePtzMoveRelative, ApiOnvifDevicePtzMoveAbsolute } from "../../wailsjs/go/ui/App"

interface CommonResponse {
  status: number,
  message: string,
}

interface DevicesDetailsInfo {
  manufacturer: string,
  model: string,
  firmware_version: string,
  serial_number: string,
  hardware_id: string,
}

interface DevicesDetailsNetwork {
  iface: string,
  addr: string,
  mac_addr: string,
}

interface DevicesDetails {
  info: DevicesDetailsInfo,
  network: DevicesDetailsNetwork,
}

interface DevicesProfileDevice {
  details: string,
  port: number,
  username: string,
  password: string
}

interface DevicesProfile {
  name: string,
  device: DevicesProfileDevice,
}

interface DevicesData {
  details: DevicesDetails,
  profile: DevicesProfile
}

interface DevicesResponse extends CommonResponse {
  data: DevicesData[]
}

interface ProfileData {
  name: string,
  token: string,
}

interface ProfileResponse extends CommonResponse {
  data: ProfileData[]
}

interface StreamData {
  url: string
}

interface StreamResponse extends CommonResponse {
  data: StreamData
}

interface PTZStatusData {
  x: number
  y: number
}

interface PTZStatusResponse extends CommonResponse {
  data: PTZStatusData
}

export default function Index() {
  const navigate = useNavigate()
  const { notify, Notifybar } = useNotify()
  const apikey = localStorage.getItem("apikey")

  const [devices, setDevices] = useState<DevicesData[]>([])
  const [deviceProfile, setDeviceProfile] = useState<ProfileData[]>([])
  const [isPendingForDeviceProfile, startDeviceProfileTransition] = useTransition()

  const [streamPanelOpen, setStreamPanelOpen] = useState<boolean>(false)
  const [streamName, setStreamName] = useState<string>("")
  const [streamTitle, setStreamTitle] = useState<string>("")
  const [streamUrl, setStreamUrl] = useState<string>("")

  const [XStepValue, SetXStepValue] = useState<number>(0.1)
  const [YStepValue, SetYStepValue] = useState<number>(0.1)
  const [XDefaultValue, SetXDefaultValue] = useState<number>(0)
  const [YDefaultValue, SetYDefaultValue] = useState<number>(0)

  const XStepValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    SetXStepValue(Number(e.target.value))
  }

  const YStepValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    SetYStepValue(Number(e.target.value))
  }

  const XDefaultValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    SetXDefaultValue(Number(e.target.value))
  }

  const YDefaultValueChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    SetYDefaultValue(Number(e.target.value))
  }

  const CheckAuth = useCallback(async () => {
    if (apikey) {
      const response = await ApiAuthCheck(apikey)
      if (response.status !== 1) {
        navigate("/login", { replace: true })
      }
    } else {
      navigate("/login", { replace: true })
    }
  }, [])

  const ListDevices = useCallback(async () => {
    if (apikey) {
      const jData = await ApiOnvifDevices(apikey)
      if (jData.status === 1) {
        setDevices(jData.data)
      } else {
        notify("获取设备信息失败", "danger")
      }
    }
  }, [])

  const GetMediaStreamInfo = async (name: string) => {
    setDeviceProfile([])
    if (apikey) {
      startDeviceProfileTransition(async () => {
        const jData = await ApiOnvifDeviceProfile(apikey, name)
        if (jData.status === 1) {
          setDeviceProfile(jData.data)
        } else {
          notify("获取媒体信息失败", "danger")
        }
      })
    }
  }

  const StremPanelClose = () => {
    setStreamPanelOpen(false)
  }

  const StremPanelOpen = async (profile: ProfileData, device: DevicesData) => {
    if (apikey) {
      const ptzStatus: PTZStatusData = await PTZStatus(device.profile.name)
      SetXDefaultValue(ptzStatus.x)
      SetYDefaultValue(ptzStatus.y)

      const jData = await ApiOnvifDeviceStreamurl(apikey, device.profile.name, profile.token, device.profile.device.username, device.profile.device.password)
      if (jData.status === 1) {
        setStreamName(device.profile.name)
        setStreamTitle(profile.name)
        setStreamUrl(jData.data.url)
      } else {
        notify("获取播放地址失败", "danger")
      }
      setStreamPanelOpen(true)
    }
  }

  const StreamPlay = async (stream_url: string) => {
    if (apikey) {
      let data = new (player.PlayParas)
      data.url = stream_url
      data.width = "1280"
      data.height = "720"
      const jData = await ApiOnvifPlay(apikey, data)
      if (jData.status === 1) {
        notify("开始播放", "primary")
      } else {
        notify("打开播放器失败:" + jData.message, "danger")
      }
    }
  }

  const PTZStatus = async (name: string) => {
    if (apikey) {
      const jData = await ApiOnvifDevicePtzStatus(apikey, name)
      if (jData.status === 1) {
        return jData.data
      }
    }
    return { x: 0, y: 0 }
  }

  const PTZAbsoluteMove = async (x: number, y: number) => {
    if (Math.abs(x) > 1 || Math.abs(y) > 1) {
      notify("位移距离必须在 0 ~ 1 之间", "danger")
      return
    }

    let data = new (ui.PtzControl)
    data.name = streamName
    data.axes.x = Number(x)
    data.axes.y = Number(y)

    if (apikey) {
      const jData = await ApiOnvifDevicePtzMoveAbsolute(apikey, data)
      if (jData.status === 1) {
        notify(`移动距离 X=${x} Y=${y}`, "primary")
      }
    }
  }

  const PTZRelativeMove = async (x: number, y: number) => {
    if (Math.abs(x) > 1 || Math.abs(y) > 1) {
      notify("位移距离必须在 0 ~ 1 之间", "danger")
      return
    }

    let data = new (ui.PtzControl)
    let axes = new (ui.PtzAxes)
    axes.x = Number(x)
    axes.y = Number(y)

    data.name = streamName
    data.axes = axes

    if (apikey) {
      const jData = await ApiOnvifDevicePtzMoveRelative(apikey, data)
      if (jData.status === 1) {
        notify(`移动距离 X=${x} Y=${y}`, "primary")
      }
    }
  }

  useEffect(() => {
    CheckAuth()
  }, [CheckAuth])

  useEffect(() => {
    ListDevices()
  }, [ListDevices])

  return (
    <Box
      sx={{
        maxWidth: 320,
        m: "0 auto",
        height: "100vh"
      }}
    >
      {Notifybar}
      <Box sx={{ pt: 2 }}>
        <Typography level="h2" sx={{ pb: 2 }}>设备列表</Typography>
        {devices.map((device: DevicesData, index: number) => {
          return (
            <Box key={"device" + index}>
              <Card>
                <Typography level="title-lg">{device.profile.name}</Typography>
                <Typography level="body-sm">{device.details.network.addr}</Typography>
                <CardContent>
                  <Box sx={{ pb: 2 }}>
                    <Typography level="body-sm">{device.details.info.manufacturer}</Typography>
                    <Typography level="body-sm">{device.details.info.model} v{device.details.info.hardware_id}</Typography>
                    <Typography level="body-sm">{device.details.info.firmware_version}</Typography>
                  </Box>
                  <Button
                    variant="solid"
                    size="md"
                    color="primary"
                    loading={isPendingForDeviceProfile}
                    onClick={() => GetMediaStreamInfo(device.profile.name)}
                  >
                    获取视频流
                  </Button>
                  {deviceProfile.length !== 0 && <Box sx={{ p: 2 }}>
                    <Stack
                      direction="row"
                      spacing={1}
                      sx={{ justifyContent: "center", alignItems: "center" }}
                    >
                      {deviceProfile.map((profile: ProfileData, index: number) => {
                        return (
                          <Chip
                            key={"profile" + index}
                            color="neutral"
                            variant="outlined"
                            onClick={() => StremPanelOpen(profile, device)}
                          >
                            {profile.name}
                          </Chip>
                        )
                      })}
                    </Stack>
                  </Box>}
                </CardContent>
              </Card>
            </Box>
          )
        })}
      </Box>

      <Modal open={streamPanelOpen} onClose={StremPanelClose}>
        <ModalDialog layout="fullscreen">
          <DialogTitle sx={{ justifyContent: "center" }}>{streamName} - {streamTitle}</DialogTitle>
          <DialogContent sx={{ textAlign: "center" }}>
            <Box sx={{ p: 2 }}>
              <Button onClick={() => StreamPlay(streamUrl)}>
                打开播放器
              </Button>

              <Box>
                <Typography sx={{ p: 2 }} level="h3">云台控制</Typography>

                <Box sx={{ width: 240, margin: "0 auto" }}>
                  <Typography level="body-lg">相对坐标</Typography>
                  <Stack spacing={2} sx={{ justifyContent: "center", alignItems: "center", pt: 2 }}>
                    <Stack spacing={2} sx={{ justifyContent: "center", alignItems: "center" }}>
                      <Input
                        placeholder="步进"
                        type="number"
                        defaultValue={XStepValue}
                        onChange={XStepValueChange}
                      />
                      <Stack direction="row" spacing={2} sx={{ justifyContent: "center", alignItems: "center" }}>
                        <IconButton onClick={() => PTZRelativeMove(-XStepValue, 0)}>
                          <LeftArrowIcon color="primary" fontSize="xl4" />
                        </IconButton>

                        <IconButton onClick={() => PTZRelativeMove(XStepValue, 0)}>
                          <RightArrowIcon color="primary" fontSize="xl4" />
                        </IconButton>
                      </Stack>
                    </Stack>

                    <Stack spacing={1} sx={{ justifyContent: "center", alignItems: "center" }}>
                      <Input
                        placeholder="步进"
                        type="number"
                        defaultValue={YStepValue}
                        onChange={YStepValueChange}
                      />
                      <Stack direction="row" spacing={2} sx={{ justifyContent: "center", alignItems: "center" }}>
                        <IconButton onClick={() => PTZRelativeMove(0, YStepValue)}>
                          <UpArrowIcon color="primary" fontSize="xl4" />
                        </IconButton>
                        <IconButton onClick={() => PTZRelativeMove(0, -YStepValue)}>
                          <DownArrowIcon color="primary" fontSize="xl4" />
                        </IconButton>
                      </Stack>
                    </Stack>
                  </Stack>
                </Box>

                <Box sx={{ width: 240, margin: "0 auto" }}>
                  <Typography sx={{ p: 2 }} level="body-lg">绝对坐标</Typography>
                  <Stack
                    spacing={2}
                    sx={{
                      justifyContent: "center",
                      alignItems: "center",
                    }}
                  >
                    <Input
                      fullWidth
                      placeholder="X"
                      type="number"
                      defaultValue={XDefaultValue}
                      onChange={XDefaultValueChange}
                      startDecorator="X"
                    />
                    <Input
                      fullWidth
                      placeholder="Y"
                      type="number"
                      defaultValue={YDefaultValue}
                      onChange={YDefaultValueChange}
                      startDecorator="Y"
                    />
                    <Button fullWidth onClick={() => PTZAbsoluteMove(XDefaultValue, YDefaultValue)}>确定</Button>
                  </Stack>
                </Box>

              </Box>
            </Box>
          </DialogContent>
          <ModalClose />
        </ModalDialog>
      </Modal>
    </Box>
  )
}