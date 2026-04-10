import React, { useState } from 'react';

interface DescriptorProps {
  descriptors: string[] // non empty array trick
}

const Descriptor: React.FC<DescriptorProps> = ({ descriptors }) => {
  const [display, setDisplay] = useState<boolean>(false);

  const toggle = () => {
    setDisplay(!display)
  }
  return (
    <div className="descriptor-containers">
      <div
        className={`${display ? 'descriptors-opened' : 'descriptors-closed'} descriptor-list`}
        onClick={toggle}
      >
        {descriptors[0]}
      </div>
      {display && (
        <ul className='descriptors-list'>
          {descriptors.map((e) => { return (<li key={e}>{e}</li>) })}
        </ul>
      )}
    </div >
  );
};

export default Descriptor;

