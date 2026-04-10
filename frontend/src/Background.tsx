import React, { useState } from 'react';
import { BackgroundBackground } from './types';

interface BackgroundProps {
  background: BackgroundBackground
}

const Background: React.FC<BackgroundProps> = ({ background }) => {
  const [display, setDisplay] = useState<boolean>(false);

  const toggle = () => {
    setDisplay(!display)
  }
  return (
    <div className="background-containers">
      <div
        className={`${display ? 'background-opened' : 'background-closed'} background`}
        onClick={toggle}
      >
        {background.name}
      </div>
      {display && (
        <div className='background-details'>
          <div className='requirements-container'>
            <div className='requirements'>
              Requirements: {background.requirements}
            </div>
          </div>
          <div className='abilities-container'>
            <div className='abilities'>
              <ul>
                {background.abilities.map((e) => { return (<li key={e}>{e}</li>) })}
              </ul>
            </div>
          </div>
        </div>
      )}
    </div >
  );
};

export default Background;

