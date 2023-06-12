const backendUri = import.meta.env.VITE_BACKEND_URI
const authorizeUri = import.meta.env.VITE_AUTHORIZE_URI

interface Config {
    authorizeUri: string,
    backendUri: string;
    uriPrefix: string;
    channelsUri: string;
    summarizeUri: string;
    summarizeIAUri: string;
}

const config: Config = {
    authorizeUri,
    backendUri,
    uriPrefix: import.meta.env.VITE_URI_PREFIX,
    channelsUri: `${backendUri}/channels`,
    summarizeUri: `${backendUri}/summarize`,
    summarizeIAUri: `${backendUri}/summarize-ia`
}

export default config
