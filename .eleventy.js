export default function(cfg) {
  cfg.addPassthroughCopy("xmit.toml");
  return {
    dir: {
      output: "dist",
    },
  };
};
