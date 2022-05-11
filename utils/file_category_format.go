package utils

//FileCategory 文件类型
var FileCategory = map[string]struct{}{
	"philosophy":     {}, //哲学
	"religion":       {}, //宗教
	"ethics":         {}, //伦理
	"logic":          {}, //逻辑
	"aesthetics":     {}, //美学
	"psychology":     {}, //心理
	"language":       {}, //语言
	"literature":     {}, //文学
	"art":            {}, //艺术
	"political":      {}, // 政治
	"economic":       {}, //经济
	"military":       {}, //军事
	"law":            {}, //法律
	"education":      {}, //教育
	"sports":         {}, //体育
	"media":          {}, //传媒
	"information":    {}, //资讯
	"management":     {}, //管理
	"business":       {}, //商贸
	"history":        {}, //历史
	"archaeological": {}, //考古
	"nation":         {}, //民族
	"life":           {}, //生活
	"financial":      {}, //财经
	"statistics":     {}, //统计
	"social":         {}, //社会
	"music":          {}, //音乐
	"technology":     {}, //科技
	"pet":            {}, // 宠物
}

const OtherFormat = "other"

// FileFormat 文件格式
var FileFormat = map[string]struct{}{
	"text":       {}, //包括 文本文件 电子书文件 压缩文件
	"data":       {}, //包括 数据文件
	"video":      {},
	"audio":      {},
	"image":      {}, //  包括 3D图像文件 位图文件 矢量图文件
	"executable": {}, //可执行文件
	OtherFormat:  {},
}

var FileFormat2Suffix = map[string][]string{
	"text":       {"adoc", "ans", "apkg", "asc", "ass", "bbl", "bib", "bibtex", "csk", "csv", "des", "doc", "docm", "docx", "fdf", "fdx", "fdxt", "hwp", "info", "log", "lst", "ltx", "markdn", "markdown", "mbox", "md", "mdown", "msg", "nfo", "odm", "odt", "ott", "pages", "psb", "rtf", "smi", "srt", "ssa", "strings", "sxw", "tex", "txt", "vmg", "vnt", "wp5", "wpd", "wps", "wps", "wri", "xfdf", "acsm", "apnx", "azw", "azw1", "azw3", "cb7", "cba", "cbr", "cbt", "cbz", "ceb", "cebx", "chm", "epub", "fb2", "ibooks", "lit", "mobi", "pdg", "snb", "teb", "tpz", "umd", "pdf", "1", "7z", "ace", "alz", "arc", "arj", "b1", "br", "bz", "bz2", "bzip", "bzip2", "cab", "cb7", "cbr", "cbt", "cbz", "cpgz", "cpio", "dd", "deb", "edxz", "egg", "emz", "enlx", "gz", "gzip", "hqx", "isz", "jar", "kz", "lha", "lz", "lz4", "lzh", "lzma", "lzo", "mpq", "pak", "pea", "pet", "pkg", "pup", "r00", "r01", "r02", "r03", "r04", "r05", "rar", "rpm", "shar", "shr", "sit", "sitx", "tar", "tar", "bz2", "tar", "gz", "tar", "lz", "tar", "lzma", "tar", "xz", "tar", "z", "taz", "tb2", "tbz", "tbz2", "tgz", "tlz", "tlzma", "tpz", "txz", "tz", "uue", "whl", "xar", "xip", "xz", "z", "zip", "zipx"},
	"data":       {"aae", "aca", "adt", "aifb", "approj", "bdic", "bin", "blg", "cert", "crtx", "csv", "dat", "data", "dcr", "ddb", "def", "dif", "dps", "dsl", "dtp", "efx", "em", "enl", "enlx", "enw", "fcpevent", "flipchart", "flo", "flo", "flp", "frm", "gan", "gcw", "ged", "gedcom", "gms", "grd", "h4", "h5", "hdf", "hdf4", "hdf5", "he4", "he5", "hsc", "idx", "jms", "jpr", "json", "key", "ld2", "lib", "lsd", "m", "marc", "mbx", "mdx", "mm", "mmf", "mpp", "mtb", "notebook", "obb", "odf", "odp", "ofx", "otp", "ova", "ovf", "pdb", "pes", "pps", "ppsm", "ppsx"},
	"video":      {"264", "3g2", "3gp", "3gp2", "3gpp", "3gpp2", "aaf", "aep", "aepx", "aet", "aetx", "amv", "arf", "asf", "asx", "avi", "bik", "camproj", "camrec", "ced", "cmproj", "cmrec", "csf", "dat", "divx", "dv", "dzm", "f4p", "f4v", "fla", "flc", "flh", "fli", "flic", "flp", "flv", "flx", "h264", "iff", "ifo", "ismc", "ismv", "m1v", "m2p", "m2t", "m2ts", "m2v", "m4v", "mk3d", "mks", "mkv", "mov", "mp2v", "mp4", "mp4v", "mpe", "mpeg", "mpeg1", "mpeg2", "mpeg4", "mpg", "mpv", "mpv2", "mswmm", "mt2s", "mts", "nut", "ogm", "ogx", "ovg", "pds", "qlv", "qmv", "qsv", "qt", "r3d", "ram", "rm", "rmd", "rmhd", "rmm", "rmvb", "rp", "rv", "srt", "stl", "swf", "trec", "ts", "usf", "vob", "vro", "vtt", "webm", "wmmp"},
	"audio":      {"3ga", "aa", "aac", "aax", "ac3", "act", "adpcm", "adt", "adts", "aif", "aifc", "aiff", "alac", "amr", "ape", "asd", "au", "au", "aup", "aup3", "caf", "cda", "cdr", "dts", "dvf", "f4a", "flac", "gpx", "gsm", "isma", "m1a", "m2a", "m3u", "m3u8", "m4a", "m4b", "m4p", "m4r", "mid", "midi", "mka", "mmf", "mp1", "mp2", "mp3", "mpa", "mpc", "mpg2", "mui", "nsf", "oga", "ogg", "oma", "opus", "ptb", "ptx", "ptxt", "ra", "raw", "rmi", "sid", "wav", "wma", "wpl", "wve", "xmf", "xspf"},
	"image":      {"3dl", "3dm", "3ds", "abc", "asm", "bip", "blend", "bvh", "c4d", "cg", "cmf", "csm", "dae", "egg", "fbx", "glb", "gltf", "igs", "ldr", "lxf", "ma", "max", "mb", "mix", "mtl", "obj", "pcd", "ply", "pmd", "r3d", "skp", "srf", "step", "stp", "u3d", "vob", "xaf", "apng", "art", "avif", "bmp", "cur", "dcm", "dds", "dic", "dicom", "djvu", "fits", "flif", "fpx", "frm", "fts", "gbr", "gif", "hdp", "hdr", "heic", "heif", "icn", "icns", "ico", "icon", "iff", "img", "ithmb", "j2c", "j2k", "jfif", "jp2", "jpc", "jpeg", "jpf", "jpg", "jpx", "jxr", "mac", "mng", "pam", "pbm", "pcd", "pct", "pcx", "pdd", "pgf", "pgm", "pic", "pict", "pictclipping", "pjp", "pjpeg", "pjpg", "png", "pnm", "ppm", "psb", "psd", "psp", "pspimage", "sfw", "tbi", "tga", "thm", "thm", "tif", "tiff", "vst", "wbmp", "wdp", "webp", "xbm", "xcf", "ai", "art", "cdr", "cdt", "cgm", "cvs", "emf", "emz", "eps", "epsf", "epsi", "fxg", "gvdesign", "odg", "otg", "pic", "ps", "sketch", "std", "svg", "svgz", "vdx", "vsd", "vsdm", "vsdx", "vss", "vst", "vsx", "wmf", "wmz", "wpg", "xar"},
	"executable": {"aab", "air", "apk", "app", "appx", "bat", "cgi", "cmd", "com", "dex", "dmg", "ds", "dsa", "ex4", "ex5", "exe", "gadget", "gpk", "jar", "js", "jsf", "msix", "nexe", "pkg", "run", "scr", "udf", "vb", "vbs", "wsf", "xap", "xapk", "xbe", "xex"},
}

var OtherFormatExcludeSuffix = []string{"adoc", "ans", "apkg", "asc", "ass", "bbl", "bib", "bibtex", "csk", "csv", "des", "doc", "docm", "docx", "fdf", "fdx", "fdxt", "hwp", "info", "log", "lst", "ltx", "markdn", "markdown", "mbox", "md", "mdown", "msg", "nfo", "odm", "odt", "ott", "pages", "psb", "rtf", "smi", "srt", "ssa", "strings", "sxw", "tex", "txt", "vmg", "vnt", "wp5", "wpd", "wps", "wps", "wri", "xfdf", "acsm", "apnx", "azw", "azw1", "azw3", "cb7", "cba", "cbr", "cbt", "cbz", "ceb", "cebx", "chm", "epub", "fb2", "ibooks", "lit", "mobi", "pdg", "snb", "teb", "tpz", "umd", "pdf", "1", "7z", "ace", "alz", "arc", "arj", "b1", "br", "bz", "bz2", "bzip", "bzip2", "cab", "cb7", "cbr", "cbt", "cbz", "cpgz", "cpio", "dd", "deb", "edxz", "egg", "emz", "enlx", "gz", "gzip", "hqx", "isz", "jar", "kz", "lha", "lz", "lz4", "lzh", "lzma", "lzo", "mpq", "pak", "pea", "pet", "pkg", "pup", "r00", "r01", "r02", "r03", "r04", "r05", "rar", "rpm", "shar", "shr", "sit", "sitx", "tar", "tar", "bz2", "tar", "gz", "tar", "lz", "tar", "lzma", "tar", "xz", "tar", "z", "taz", "tb2", "tbz", "tbz2", "tgz", "tlz", "tlzma", "tpz", "txz", "tz", "uue", "whl", "xar", "xip", "xz", "z", "zip", "zipx", "aae", "aca", "adt", "aifb", "approj", "bdic", "bin", "blg", "cert", "crtx", "csv", "dat", "data", "dcr", "ddb", "def", "dif", "dps", "dsl", "dtp", "efx", "em", "enl", "enlx", "enw", "fcpevent", "flipchart", "flo", "flo", "flp", "frm", "gan", "gcw", "ged", "gedcom", "gms", "grd", "h4", "h5", "hdf", "hdf4", "hdf5", "he4", "he5", "hsc", "idx", "jms", "jpr", "json", "key", "ld2", "lib", "lsd", "m", "marc", "mbx", "mdx", "mm", "mmf", "mpp", "mtb", "notebook", "obb", "odf", "odp", "ofx", "otp", "ova", "ovf", "pdb", "pes", "pps", "ppsm", "ppsx", "264", "3g2", "3gp", "3gp2", "3gpp", "3gpp2", "aaf", "aep", "aepx", "aet", "aetx", "amv", "arf", "asf", "asx", "avi", "bik", "camproj", "camrec", "ced", "cmproj", "cmrec", "csf", "dat", "divx", "dv", "dzm", "f4p", "f4v", "fla", "flc", "flh", "fli", "flic", "flp", "flv", "flx", "h264", "iff", "ifo", "ismc", "ismv", "m1v", "m2p", "m2t", "m2ts", "m2v", "m4v", "mk3d", "mks", "mkv", "mov", "mp2v", "mp4", "mp4v", "mpe", "mpeg", "mpeg1", "mpeg2", "mpeg4", "mpg", "mpv", "mpv2", "mswmm", "mt2s", "mts", "nut", "ogm", "ogx", "ovg", "pds", "qlv", "qmv", "qsv", "qt", "r3d", "ram", "rm", "rmd", "rmhd", "rmm", "rmvb", "rp", "rv", "srt", "stl", "swf", "trec", "ts", "usf", "vob", "vro", "vtt", "webm", "wmmp", "3ga", "aa", "aac", "aax", "ac3", "act", "adpcm", "adt", "adts", "aif", "aifc", "aiff", "alac", "amr", "ape", "asd", "au", "au", "aup", "aup3", "caf", "cda", "cdr", "dts", "dvf", "f4a", "flac", "gpx", "gsm", "isma", "m1a", "m2a", "m3u", "m3u8", "m4a", "m4b", "m4p", "m4r", "mid", "midi", "mka", "mmf", "mp1", "mp2", "mp3", "mpa", "mpc", "mpg2", "mui", "nsf", "oga", "ogg", "oma", "opus", "ptb", "ptx", "ptxt", "ra", "raw", "rmi", "sid", "wav", "wma", "wpl", "wve", "xmf", "xspf", "3dl", "3dm", "3ds", "abc", "asm", "bip", "blend", "bvh", "c4d", "cg", "cmf", "csm", "dae", "egg", "fbx", "glb", "gltf", "igs", "ldr", "lxf", "ma", "max", "mb", "mix", "mtl", "obj", "pcd", "ply", "pmd", "r3d", "skp", "srf", "step", "stp", "u3d", "vob", "xaf", "apng", "art", "avif", "bmp", "cur", "dcm", "dds", "dic", "dicom", "djvu", "fits", "flif", "fpx", "frm", "fts", "gbr", "gif", "hdp", "hdr", "heic", "heif", "icn", "icns", "ico", "icon", "iff", "img", "ithmb", "j2c", "j2k", "jfif", "jp2", "jpc", "jpeg", "jpf", "jpg", "jpx", "jxr", "mac", "mng", "pam", "pbm", "pcd", "pct", "pcx", "pdd", "pgf", "pgm", "pic", "pict", "pictclipping", "pjp", "pjpeg", "pjpg", "png", "pnm", "ppm", "psb", "psd", "psp", "pspimage", "sfw", "tbi", "tga", "thm", "thm", "tif", "tiff", "vst", "wbmp", "wdp", "webp", "xbm", "xcf", "ai", "art", "cdr", "cdt", "cgm", "cvs", "emf", "emz", "eps", "epsf", "epsi", "fxg", "gvdesign", "odg", "otg", "pic", "ps", "sketch", "std", "svg", "svgz", "vdx", "vsd", "vsdm", "vsdx", "vss", "vst", "vsx", "wmf", "wmz", "wpg", "xar", "aab", "air", "apk", "app", "appx", "bat", "cgi", "cmd", "com", "dex", "dmg", "ds", "dsa", "ex4", "ex5", "exe", "gadget", "gpk", "jar", "js", "jsf", "msix", "nexe", "pkg", "run", "scr", "udf", "vb", "vbs", "wsf", "xap", "xapk", "xbe", "xex"}
