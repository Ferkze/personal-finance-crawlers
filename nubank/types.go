package main

import "encoding/xml"

type NubankEntry struct {
	Category string
	Time 	 string
	Place 	 string
	Value 	 string
}

// Hierarchy - XML's hierarchy of nodes
type Hierarchy struct {
	XMLName  xml.Name `xml:"hierarchy"`
	Text     string   `xml:",chardata"`
	Rotation string   `xml:"rotation,attr"`
	Node     struct {
		Text          string `xml:",chardata"`
		Index         string `xml:"index,attr"`
		AttrText      string `xml:"text,attr"`
		ResourceID    string `xml:"resource-id,attr"`
		Class         string `xml:"class,attr"`
		Package       string `xml:"package,attr"`
		ContentDesc   string `xml:"content-desc,attr"`
		Checkable     string `xml:"checkable,attr"`
		Checked       string `xml:"checked,attr"`
		Clickable     string `xml:"clickable,attr"`
		Enabled       string `xml:"enabled,attr"`
		Focusable     string `xml:"focusable,attr"`
		Focused       string `xml:"focused,attr"`
		Scrollable    string `xml:"scrollable,attr"`
		LongClickable string `xml:"long-clickable,attr"`
		Password      string `xml:"password,attr"`
		Selected      string `xml:"selected,attr"`
		Bounds        string `xml:"bounds,attr"`
		Node          []struct {
			Text          string `xml:",chardata"`
			Index         string `xml:"index,attr"`
			AttrText      string `xml:"text,attr"`
			ResourceID    string `xml:"resource-id,attr"`
			Class         string `xml:"class,attr"`
			Package       string `xml:"package,attr"`
			ContentDesc   string `xml:"content-desc,attr"`
			Checkable     string `xml:"checkable,attr"`
			Checked       string `xml:"checked,attr"`
			Clickable     string `xml:"clickable,attr"`
			Enabled       string `xml:"enabled,attr"`
			Focusable     string `xml:"focusable,attr"`
			Focused       string `xml:"focused,attr"`
			Scrollable    string `xml:"scrollable,attr"`
			LongClickable string `xml:"long-clickable,attr"`
			Password      string `xml:"password,attr"`
			Selected      string `xml:"selected,attr"`
			Bounds        string `xml:"bounds,attr"`
			Node          struct {
				Text          string `xml:",chardata"`
				Index         string `xml:"index,attr"`
				AttrText      string `xml:"text,attr"`
				ResourceID    string `xml:"resource-id,attr"`
				Class         string `xml:"class,attr"`
				Package       string `xml:"package,attr"`
				ContentDesc   string `xml:"content-desc,attr"`
				Checkable     string `xml:"checkable,attr"`
				Checked       string `xml:"checked,attr"`
				Clickable     string `xml:"clickable,attr"`
				Enabled       string `xml:"enabled,attr"`
				Focusable     string `xml:"focusable,attr"`
				Focused       string `xml:"focused,attr"`
				Scrollable    string `xml:"scrollable,attr"`
				LongClickable string `xml:"long-clickable,attr"`
				Password      string `xml:"password,attr"`
				Selected      string `xml:"selected,attr"`
				Bounds        string `xml:"bounds,attr"`
				Node          struct {
					Text          string `xml:",chardata"`
					Index         string `xml:"index,attr"`
					AttrText      string `xml:"text,attr"`
					ResourceID    string `xml:"resource-id,attr"`
					Class         string `xml:"class,attr"`
					Package       string `xml:"package,attr"`
					ContentDesc   string `xml:"content-desc,attr"`
					Checkable     string `xml:"checkable,attr"`
					Checked       string `xml:"checked,attr"`
					Clickable     string `xml:"clickable,attr"`
					Enabled       string `xml:"enabled,attr"`
					Focusable     string `xml:"focusable,attr"`
					Focused       string `xml:"focused,attr"`
					Scrollable    string `xml:"scrollable,attr"`
					LongClickable string `xml:"long-clickable,attr"`
					Password      string `xml:"password,attr"`
					Selected      string `xml:"selected,attr"`
					Bounds        string `xml:"bounds,attr"`
					Node          struct {
						Text          string `xml:",chardata"`
						Index         string `xml:"index,attr"`
						AttrText      string `xml:"text,attr"`
						ResourceID    string `xml:"resource-id,attr"`
						Class         string `xml:"class,attr"`
						Package       string `xml:"package,attr"`
						ContentDesc   string `xml:"content-desc,attr"`
						Checkable     string `xml:"checkable,attr"`
						Checked       string `xml:"checked,attr"`
						Clickable     string `xml:"clickable,attr"`
						Enabled       string `xml:"enabled,attr"`
						Focusable     string `xml:"focusable,attr"`
						Focused       string `xml:"focused,attr"`
						Scrollable    string `xml:"scrollable,attr"`
						LongClickable string `xml:"long-clickable,attr"`
						Password      string `xml:"password,attr"`
						Selected      string `xml:"selected,attr"`
						Bounds        string `xml:"bounds,attr"`
						Node          struct {
							Text          string `xml:",chardata"`
							Index         string `xml:"index,attr"`
							AttrText      string `xml:"text,attr"`
							ResourceID    string `xml:"resource-id,attr"`
							Class         string `xml:"class,attr"`
							Package       string `xml:"package,attr"`
							ContentDesc   string `xml:"content-desc,attr"`
							Checkable     string `xml:"checkable,attr"`
							Checked       string `xml:"checked,attr"`
							Clickable     string `xml:"clickable,attr"`
							Enabled       string `xml:"enabled,attr"`
							Focusable     string `xml:"focusable,attr"`
							Focused       string `xml:"focused,attr"`
							Scrollable    string `xml:"scrollable,attr"`
							LongClickable string `xml:"long-clickable,attr"`
							Password      string `xml:"password,attr"`
							Selected      string `xml:"selected,attr"`
							Bounds        string `xml:"bounds,attr"`
							Node          []struct {
								Text          string `xml:",chardata"`
								Index         string `xml:"index,attr"`
								AttrText      string `xml:"text,attr"`
								ResourceID    string `xml:"resource-id,attr"`
								Class         string `xml:"class,attr"`
								Package       string `xml:"package,attr"`
								ContentDesc   string `xml:"content-desc,attr"`
								Checkable     string `xml:"checkable,attr"`
								Checked       string `xml:"checked,attr"`
								Clickable     string `xml:"clickable,attr"`
								Enabled       string `xml:"enabled,attr"`
								Focusable     string `xml:"focusable,attr"`
								Focused       string `xml:"focused,attr"`
								Scrollable    string `xml:"scrollable,attr"`
								LongClickable string `xml:"long-clickable,attr"`
								Password      string `xml:"password,attr"`
								Selected      string `xml:"selected,attr"`
								Bounds        string `xml:"bounds,attr"`
								NAF           string `xml:"NAF,attr"`
								Node          struct {
									Text          string `xml:",chardata"`
									Index         string `xml:"index,attr"`
									AttrText      string `xml:"text,attr"`
									ResourceID    string `xml:"resource-id,attr"`
									Class         string `xml:"class,attr"`
									Package       string `xml:"package,attr"`
									ContentDesc   string `xml:"content-desc,attr"`
									Checkable     string `xml:"checkable,attr"`
									Checked       string `xml:"checked,attr"`
									Clickable     string `xml:"clickable,attr"`
									Enabled       string `xml:"enabled,attr"`
									Focusable     string `xml:"focusable,attr"`
									Focused       string `xml:"focused,attr"`
									Scrollable    string `xml:"scrollable,attr"`
									LongClickable string `xml:"long-clickable,attr"`
									Password      string `xml:"password,attr"`
									Selected      string `xml:"selected,attr"`
									Bounds        string `xml:"bounds,attr"`
									Node          []struct {
										Text          string `xml:",chardata"`
										Index         string `xml:"index,attr"`
										AttrText      string `xml:"text,attr"`
										ResourceID    string `xml:"resource-id,attr"`
										Class         string `xml:"class,attr"`
										Package       string `xml:"package,attr"`
										ContentDesc   string `xml:"content-desc,attr"`
										Checkable     string `xml:"checkable,attr"`
										Checked       string `xml:"checked,attr"`
										Clickable     string `xml:"clickable,attr"`
										Enabled       string `xml:"enabled,attr"`
										Focusable     string `xml:"focusable,attr"`
										Focused       string `xml:"focused,attr"`
										Scrollable    string `xml:"scrollable,attr"`
										LongClickable string `xml:"long-clickable,attr"`
										Password      string `xml:"password,attr"`
										Selected      string `xml:"selected,attr"`
										Bounds        string `xml:"bounds,attr"`
										Node          struct {
											Text          string `xml:",chardata"`
											Index         string `xml:"index,attr"`
											AttrText      string `xml:"text,attr"`
											ResourceID    string `xml:"resource-id,attr"`
											Class         string `xml:"class,attr"`
											Package       string `xml:"package,attr"`
											ContentDesc   string `xml:"content-desc,attr"`
											Checkable     string `xml:"checkable,attr"`
											Checked       string `xml:"checked,attr"`
											Clickable     string `xml:"clickable,attr"`
											Enabled       string `xml:"enabled,attr"`
											Focusable     string `xml:"focusable,attr"`
											Focused       string `xml:"focused,attr"`
											Scrollable    string `xml:"scrollable,attr"`
											LongClickable string `xml:"long-clickable,attr"`
											Password      string `xml:"password,attr"`
											Selected      string `xml:"selected,attr"`
											Bounds        string `xml:"bounds,attr"`
											Node          []struct {
												Text          string `xml:",chardata"`
												Index         string `xml:"index,attr"`
												AttrText      string `xml:"text,attr"`
												ResourceID    string `xml:"resource-id,attr"`
												Class         string `xml:"class,attr"`
												Package       string `xml:"package,attr"`
												ContentDesc   string `xml:"content-desc,attr"`
												Checkable     string `xml:"checkable,attr"`
												Checked       string `xml:"checked,attr"`
												Clickable     string `xml:"clickable,attr"`
												Enabled       string `xml:"enabled,attr"`
												Focusable     string `xml:"focusable,attr"`
												Focused       string `xml:"focused,attr"`
												Scrollable    string `xml:"scrollable,attr"`
												LongClickable string `xml:"long-clickable,attr"`
												Password      string `xml:"password,attr"`
												Selected      string `xml:"selected,attr"`
												Bounds        string `xml:"bounds,attr"`
												Node          []struct {
													Text          string `xml:",chardata"`
													Index         string `xml:"index,attr"`
													AttrText      string `xml:"text,attr"`
													ResourceID    string `xml:"resource-id,attr"`
													Class         string `xml:"class,attr"`
													Package       string `xml:"package,attr"`
													ContentDesc   string `xml:"content-desc,attr"`
													Checkable     string `xml:"checkable,attr"`
													Checked       string `xml:"checked,attr"`
													Clickable     string `xml:"clickable,attr"`
													Enabled       string `xml:"enabled,attr"`
													Focusable     string `xml:"focusable,attr"`
													Focused       string `xml:"focused,attr"`
													Scrollable    string `xml:"scrollable,attr"`
													LongClickable string `xml:"long-clickable,attr"`
													Password      string `xml:"password,attr"`
													Selected      string `xml:"selected,attr"`
													Bounds        string `xml:"bounds,attr"`
													Node          []struct {
														Text          string `xml:",chardata"`
														NAF           string `xml:"NAF,attr"`
														Index         string `xml:"index,attr"`
														AttrText      string `xml:"text,attr"`
														ResourceID    string `xml:"resource-id,attr"`
														Class         string `xml:"class,attr"`
														Package       string `xml:"package,attr"`
														ContentDesc   string `xml:"content-desc,attr"`
														Checkable     string `xml:"checkable,attr"`
														Checked       string `xml:"checked,attr"`
														Clickable     string `xml:"clickable,attr"`
														Enabled       string `xml:"enabled,attr"`
														Focusable     string `xml:"focusable,attr"`
														Focused       string `xml:"focused,attr"`
														Scrollable    string `xml:"scrollable,attr"`
														LongClickable string `xml:"long-clickable,attr"`
														Password      string `xml:"password,attr"`
														Selected      string `xml:"selected,attr"`
														Bounds        string `xml:"bounds,attr"`
														Node          []struct {
															Text          string `xml:",chardata"`
															Index         string `xml:"index,attr"`
															AttrText      string `xml:"text,attr"`
															ResourceID    string `xml:"resource-id,attr"`
															Class         string `xml:"class,attr"`
															Package       string `xml:"package,attr"`
															ContentDesc   string `xml:"content-desc,attr"`
															Checkable     string `xml:"checkable,attr"`
															Checked       string `xml:"checked,attr"`
															Clickable     string `xml:"clickable,attr"`
															Enabled       string `xml:"enabled,attr"`
															Focusable     string `xml:"focusable,attr"`
															Focused       string `xml:"focused,attr"`
															Scrollable    string `xml:"scrollable,attr"`
															LongClickable string `xml:"long-clickable,attr"`
															Password      string `xml:"password,attr"`
															Selected      string `xml:"selected,attr"`
															Bounds        string `xml:"bounds,attr"`
														} `xml:"node"`
													} `xml:"node"`
												} `xml:"node"`
											} `xml:"node"`
										} `xml:"node"`
									} `xml:"node"`
								} `xml:"node"`
							} `xml:"node"`
						} `xml:"node"`
					} `xml:"node"`
				} `xml:"node"`
			} `xml:"node"`
		} `xml:"node"`
	} `xml:"node"`
}

// Node - XML's node structure
type ParentNode struct {
	Text          string `xml:",chardata"`
	Index         string `xml:"index,attr"`
	AttrText      string `xml:"text,attr"`
	ResourceID    string `xml:"resource-id,attr"`
	Class         string `xml:"class,attr"`
	Package       string `xml:"package,attr"`
	ContentDesc   string `xml:"content-desc,attr"`
	Checkable     string `xml:"checkable,attr"`
	Checked       string `xml:"checked,attr"`
	Clickable     string `xml:"clickable,attr"`
	Enabled       string `xml:"enabled,attr"`
	Focusable     string `xml:"focusable,attr"`
	Focused       string `xml:"focused,attr"`
	Scrollable    string `xml:"scrollable,attr"`
	LongClickable string `xml:"long-clickable,attr"`
	Password      string `xml:"password,attr"`
	Selected      string `xml:"selected,attr"`
	Bounds        string `xml:"bounds,attr"`
	Node          interface{} `xml:"node"`
}
// Node - XML's node structure
type LeafNode struct {
	Text          string `xml:",chardata"`
	Index         string `xml:"index,attr"`
	AttrText      string `xml:"text,attr"`
	ResourceID    string `xml:"resource-id,attr"`
	Class         string `xml:"class,attr"`
	Package       string `xml:"package,attr"`
	ContentDesc   string `xml:"content-desc,attr"`
	Checkable     string `xml:"checkable,attr"`
	Checked       string `xml:"checked,attr"`
	Clickable     string `xml:"clickable,attr"`
	Enabled       string `xml:"enabled,attr"`
	Focusable     string `xml:"focusable,attr"`
	Focused       string `xml:"focused,attr"`
	Scrollable    string `xml:"scrollable,attr"`
	LongClickable string `xml:"long-clickable,attr"`
	Password      string `xml:"password,attr"`
	Selected      string `xml:"selected,attr"`
	Bounds        string `xml:"bounds,attr"`
	Node          interface{} `xml:"node"`
}
