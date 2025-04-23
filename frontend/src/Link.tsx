import React from 'react';

interface LinkProps {
  href: string
  text: string
  className?: string;
  target?: '_blank' | '_self' | '_parent' | '_top';
  rel?: string;
  onClick?: (e: React.MouseEvent<HTMLAnchorElement>) => void;
}

const Link: React.FC<LinkProps> = ({
  href,
  text,
  className = '',
  target = '_self',
  rel = target === '_blank' ? 'noopener noreferrer' : '',
  onClick
}) => {
  return (
    <a
      href={href}
      className={`text-blue-600 hover:text-blue-800 underline ${className}`}
      target={target}
      rel={rel}
      onClick={onClick}
    >
      {text}
    </a>
  );
};

export default Link;
