import { Image } from "@chakra-ui/react";
import { useState } from "react";
import { useLocation } from "react-router-dom";
import CopyURL from "./CopyURL";

const Description = () => {
  const location = useLocation();
  const [imageURL, _] = useState<string>(location.state.imageURL);

  return (
    <div>
      <div className="flex justify-center items-center mr-5">
        <div>
          <p className="text-blue-500 font-bold text-xl text-center">
            Here is the Image
          </p>
          <Image
            boxSize={"300px"}
            src={imageURL}
            alt={`Image: ${imageURL}`}
            borderRadius={"lg"}
          />
        </div>
      </div>
      <div>
        <p className="flex justify-center text-blue-500 font-bold text-xl mt-4 mb-2">
          {" "}
          Here is the URL{" "}
        </p>
        <CopyURL imageURL={imageURL} />
      </div>
    </div>
  );
};

export default Description;
