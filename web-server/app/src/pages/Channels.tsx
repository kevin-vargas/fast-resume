import Grid from "@mui/material/Grid";
import useQueryChannels from "../hooks/useQueryChannels";
import Card from "../components/Card";
import Box from "@mui/material/Box";
import { useState } from "react";
import Modal from "../components/Modal";
import ChannelDetail from "../components/ChannelDetail";

export default function ChannelsPage() {
  const [selected, setSelected] = useState<string>()
  const [open, setOpen] = useState(false);
  
    const result = useQueryChannels()
    if (!result) {
      return null
    }
    const handleOpen = () => setOpen(true);
    const handleClose = () => setOpen(false);

    return (
      <>
      <Box sx={{ flexGrow: 1 }}>
        <Grid container spacing={3}>
          {
            result.data.map((c) => 
              <Grid item xs={12} sm={6} md={3}>
                <Card title={c.title} description={c.description} onClick={() => {
                  setSelected(c.id)
                  handleOpen()
                }}/>
            </Grid>
            )
          }
        </Grid>
      </Box>
      <Modal onClose={handleClose} open={open}>
        <ChannelDetail id={selected}/>
      </Modal>
      </>
    );
  }