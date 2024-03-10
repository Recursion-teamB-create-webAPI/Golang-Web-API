import { VStack, useClipboard } from "@chakra-ui/react";
import { useEffect } from "react";

type Props = {
  imageURL: string;
};

const CopyURL = ({ imageURL }: Props) => {
  const { onCopy, value, setValue, hasCopied } = useClipboard("");

  useEffect(() => {
    setValue(imageURL);
  }, []);

  return (
    <>
      <VStack mb={2} mr={3}>
        <textarea
          value={value}
          readOnly={true}
          style={{ resize: "none" }}
          className="h-40 w-80 p-3 font-bold outline-none rounded-md "
        />
        <button
          onClick={onCopy}
          className="bg-blue-500 rounded-lg px-3 py-2 text-white shadow-md"
        >
          {hasCopied ? "Copied!" : "Copy"}
        </button>
      </VStack>
    </>
  );
};

export default CopyURL;
