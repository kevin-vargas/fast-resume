import Box from '@mui/material/Box';
import Modal from '@mui/material/Modal';

interface ModalProps {
    open: boolean;
    onClose: () => void,
    children: React.ReactNode
}

const style = {
    position: 'absolute' as 'absolute',
    top: '50%',
    left: '50%',
    transform: 'translate(-50%, -50%)',
    backgroundColor: 'background.paper',
    border: '2px solid #000',
    overflow: "hidden",
    overflowY: "scroll",
    height: 600,
    width: 1200,
    p: 4,
  };

export default ({ children, ...rest }:ModalProps) => {
    return (
        <Modal
        {...rest}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
        >
            <Box sx={style}>
                {children}
            </Box>
        </Modal>
    )
}