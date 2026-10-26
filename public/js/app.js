(function () {
  "use strict";

  var config = {};
  var configEl = document.getElementById("site-config");
  if (configEl) {
    try {
      config = JSON.parse(configEl.textContent || "{}");
    } catch (e) {
      config = {};
    }
  }

  var LANG_KEY = config.langKey || "portfolio-language";
  var currentLang = "id";

  /* ----------------------------------------------------------------
     Theme toggle (dark is the default; choice persists in localStorage)
     ---------------------------------------------------------------- */
  (function theme() {
    var toggle = document.getElementById("theme-toggle");
    if (!toggle) return;

    var isDark = localStorage.getItem("theme") !== "light";
    var darkIcon = toggle.querySelector(".theme-icon-dark");
    var lightIcon = toggle.querySelector(".theme-icon-light");

    function apply() {
      document.documentElement.classList.toggle("dark", isDark);
      if (darkIcon) darkIcon.classList.toggle("hidden", !isDark);
      if (lightIcon) lightIcon.classList.toggle("hidden", isDark);
      updateThemeAria();
    }

    toggle.addEventListener("click", function () {
      isDark = !isDark;
      localStorage.setItem("theme", isDark ? "dark" : "light");
      apply();
    });

    apply();
  })();

  function updateThemeAria() {
    var toggle = document.getElementById("theme-toggle");
    if (!toggle) return;
    var isDark = document.documentElement.classList.contains("dark");
    var dict = (config.dict && config.dict[currentLang]) || {};
    var label = isDark
      ? (dict.theme_to_light || "Aktifkan tema terang")
      : (dict.theme_to_dark || "Aktifkan tema gelap");
    toggle.setAttribute("aria-label", label);
  }

  /* ----------------------------------------------------------------
     Bilingual i18n system (Indonesian default, English alternative)
     ---------------------------------------------------------------- */
  var i18nListeners = [];

  function t(key, fallback) {
    var dict = (config.dict && config.dict[currentLang]) || {};
    return dict[key] || fallback || "";
  }

  function applyLanguage(lang) {
    if (lang !== "id" && lang !== "en") lang = "id";
    currentLang = lang;
    document.documentElement.setAttribute("lang", lang);

    var dict = (config.dict && config.dict[lang]) || {};

    // 1. Static text elements: data-i18n="key"
    document.querySelectorAll("[data-i18n]").forEach(function (el) {
      var key = el.getAttribute("data-i18n");
      if (dict[key] !== undefined) {
        el.textContent = dict[key];
      }
    });

    // 2. Elements with inline data-i18n-id and data-i18n-en
    document.querySelectorAll("[data-i18n-id]").forEach(function (el) {
      var val = lang === "en" ? el.getAttribute("data-i18n-en") : el.getAttribute("data-i18n-id");
      if (val !== null && val !== undefined) {
        el.textContent = val;
      }
    });

    // 3. Attribute translations: aria-label
    document.querySelectorAll("[data-i18n-aria]").forEach(function (el) {
      var key = el.getAttribute("data-i18n-aria");
      if (dict[key] !== undefined) {
        el.setAttribute("aria-label", dict[key]);
      }
    });

    // 4. Attribute translations: title
    document.querySelectorAll("[data-i18n-title]").forEach(function (el) {
      var key = el.getAttribute("data-i18n-title");
      if (dict[key] !== undefined) {
        el.setAttribute("title", dict[key]);
      }
    });

    // 5. Attribute translations: placeholder
    document.querySelectorAll("[data-i18n-placeholder]").forEach(function (el) {
      var key = el.getAttribute("data-i18n-placeholder");
      if (dict[key] !== undefined) {
        el.setAttribute("placeholder", dict[key]);
      }
    });

    // 6. Image alt attributes
    document.querySelectorAll("[data-i18n-alt-id]").forEach(function (el) {
      var val = lang === "en" ? el.getAttribute("data-i18n-alt-en") : el.getAttribute("data-i18n-alt-id");
      if (val) el.setAttribute("alt", val);
    });

    // 7. Lightbox trigger aria-label
    document.querySelectorAll("[data-i18n-aria-alt-id]").forEach(function (el) {
      var alt = lang === "en" ? el.getAttribute("data-i18n-aria-alt-en") : el.getAttribute("data-i18n-aria-alt-id");
      var prefix = lang === "en" ? "Enlarge " : "Perbesar ";
      if (alt) el.setAttribute("aria-label", prefix + alt);
    });

    // 8. Marquee spans: data-marquee-id and data-marquee-en
    document.querySelectorAll("[data-marquee-id]").forEach(function (el) {
      var val = lang === "en" ? el.getAttribute("data-marquee-en") : el.getAttribute("data-marquee-id");
      if (val) el.innerHTML = val;
    });

    // 9. Update language switcher UI
    var label = document.getElementById("lang-label");
    if (label) label.textContent = lang.toUpperCase();

    document.querySelectorAll("[data-lang-choice]").forEach(function (btn) {
      var choice = btn.getAttribute("data-lang-choice");
      var active = choice === lang;
      btn.classList.toggle("text-text-primary", active);
      btn.classList.toggle("text-text-secondary", !active);
      var check = btn.querySelector(".lang-check");
      if (check) {
        check.classList.toggle("text-text-primary", active);
        check.classList.toggle("text-text-secondary/50", !active);
      }
    });

    // 10. Update theme aria-label for active language
    updateThemeAria();

    // 11. Notify subscribers (hero typing, chat)
    i18nListeners.forEach(function (fn) {
      try {
        fn(lang);
      } catch (e) {}
    });
  }

  (function langSwitcher() {
    var switcher = document.getElementById("lang-switcher");
    var toggle = document.getElementById("lang-toggle");
    var menu = document.getElementById("lang-menu");
    if (!switcher || !toggle || !menu) return;

    function isOpen() {
      return !menu.classList.contains("opacity-0");
    }

    function setMenu(open) {
      if (open) {
        menu.classList.remove("scale-0", "opacity-0", "pointer-events-none");
        menu.classList.add("scale-100", "opacity-100");
      } else {
        menu.classList.remove("scale-100", "opacity-100");
        menu.classList.add("scale-0", "opacity-0", "pointer-events-none");
      }
      toggle.setAttribute("aria-expanded", open ? "true" : "false");
    }

    toggle.addEventListener("click", function (e) {
      e.preventDefault();
      e.stopPropagation();
      setMenu(!isOpen());
    });

    document.querySelectorAll("[data-lang-choice]").forEach(function (btn) {
      btn.addEventListener("click", function (e) {
        e.preventDefault();
        e.stopPropagation();
        var choice = btn.getAttribute("data-lang-choice");
        setLanguage(choice);
        setMenu(false);
      });
    });

    document.addEventListener("click", function (e) {
      if (!switcher.contains(e.target)) {
        setMenu(false);
      }
    });

    document.addEventListener("keydown", function (e) {
      if (e.key === "Escape" && isOpen()) {
        setMenu(false);
        toggle.focus();
      }
    });
  })();

  function setLanguage(lang) {
    if (lang !== "id" && lang !== "en") lang = "id";
    try {
      localStorage.setItem(LANG_KEY, lang);
    } catch (e) {}
    applyLanguage(lang);
  }
  window.setPortfolioLanguage = setLanguage;

  // Initial language resolution: stored choice -> fallback "id"
  (function initLanguage() {
    var stored = null;
    try {
      stored = localStorage.getItem(LANG_KEY);
    } catch (e) {}
    if (stored !== "id" && stored !== "en") stored = "id";
    applyLanguage(stored);
  })();

  /* ----------------------------------------------------------------
     Smooth scroll + scroll spy + mobile menu
     ---------------------------------------------------------------- */
  function scrollToSection(id) {
    var section = document.getElementById(id);
    if (section) section.scrollIntoView({ behavior: "smooth", block: "start" });
  }

  (function navigation() {
    var menuToggle = document.getElementById("menu-toggle");
    var mobileMenu = document.getElementById("mobile-menu");

    function setMenu(open) {
      if (!mobileMenu || !menuToggle) return;
      mobileMenu.classList.toggle("scale-100", open);
      mobileMenu.classList.toggle("opacity-100", open);
      mobileMenu.classList.toggle("scale-0", !open);
      mobileMenu.classList.toggle("opacity-0", !open);
      mobileMenu.classList.toggle("pointer-events-none", !open);
      var openIcon = menuToggle.querySelector(".menu-icon-open");
      var closeIcon = menuToggle.querySelector(".menu-icon-close");
      if (openIcon) openIcon.classList.toggle("hidden", open);
      if (closeIcon) closeIcon.classList.toggle("hidden", !open);
      menuToggle.setAttribute("aria-expanded", open ? "true" : "false");
      var dict = (config.dict && config.dict[currentLang]) || {};
      var label = open ? (dict.menu_close || "Tutup menu navigasi") : (dict.menu_open || "Buka menu navigasi");
      menuToggle.setAttribute("aria-label", label);
    }

    if (menuToggle) {
      menuToggle.addEventListener("click", function () {
        setMenu(mobileMenu.classList.contains("opacity-0"));
      });
    }

    var scrollButtons = document.querySelectorAll("[data-scroll]");
    scrollButtons.forEach(function (btn) {
      btn.addEventListener("click", function () {
        scrollToSection(btn.getAttribute("data-scroll"));
        setMenu(false);
      });
    });

    var navLinks = document.querySelectorAll(".nav-link");
    var sections = document.querySelectorAll("section[id]");

    function setActive(current) {
      navLinks.forEach(function (link) {
        var active = link.getAttribute("data-section") === current;
        link.classList.toggle("text-text-primary", active);
        link.classList.toggle("font-bold", active);
        link.classList.toggle("text-text-secondary", !active);
        link.classList.toggle("font-medium", !active);
      });
    }

    function onScroll() {
      var current = "home";
      sections.forEach(function (section) {
        if (window.scrollY >= section.offsetTop - section.clientHeight / 3) {
          current = section.id;
        }
      });
      setActive(current);
    }

    window.addEventListener("scroll", onScroll, { passive: true });
    onScroll();
  })();

  /* ----------------------------------------------------------------
     Scroll reveals (Fade / FadeUp / FadeDown / FadeLeft / FadeRight)
     ---------------------------------------------------------------- */
  (function reveals() {
    var nodes = document.querySelectorAll(".reveal");
    if (!nodes.length) return;

    if (!("IntersectionObserver" in window)) {
      nodes.forEach(function (n) {
        n.classList.add("reveal-in");
      });
      return;
    }

    var observer = new IntersectionObserver(
      function (entries) {
        entries.forEach(function (entry) {
          if (entry.isIntersecting) {
            entry.target.classList.add("reveal-in");
            observer.unobserve(entry.target);
          }
        });
      },
      { threshold: 0.2 },
    );

    nodes.forEach(function (n) {
      observer.observe(n);
    });
  })();

  /* ----------------------------------------------------------------
     Footer year
     ---------------------------------------------------------------- */
  (function footerYear() {
    var el = document.getElementById("footer-year");
    if (el) el.textContent = String(new Date().getFullYear());
  })();

  /* ----------------------------------------------------------------
     Opening loader
     ---------------------------------------------------------------- */
  (function loader() {
    var el = document.getElementById("page-loader");
    if (!el) return;

    document.body.style.overflow = "hidden";

    var title = el.querySelector(".loader-title");
    var sub = el.querySelector(".loader-sub");
    var bar = el.querySelector(".loader-bar");

    if (title) {
      title.style.transition = "none";
      title.style.opacity = "0";
      title.style.transform = "translateY(50px)";
    }
    if (sub) {
      sub.style.transition = "none";
      sub.style.opacity = "0";
      sub.style.transform = "translateY(30px)";
    }
    if (bar) {
      bar.style.transition = "none";
      bar.style.transform = "translateX(-100%)";
    }

    requestAnimationFrame(function () {
      if (title) {
        title.style.transition = "opacity .8s ease-out .2s, transform .8s ease-out .2s";
        title.style.opacity = "1";
        title.style.transform = "translateY(0)";
      }
      if (sub) {
        sub.style.transition = "opacity .8s ease-out .4s, transform .8s ease-out .4s";
        sub.style.opacity = "1";
        sub.style.transform = "translateY(0)";
      }
      if (bar) {
        bar.style.transition = "transform 1.5s ease-in-out";
        bar.style.transform = "translateX(0)";
      }
    });

    function dismiss() {
      el.style.transition = "opacity .8s ease-in-out, transform .8s ease-in-out, filter .8s ease-in-out";
      el.style.opacity = "0";
      el.style.transform = "translateY(-100px)";
      el.style.filter = "blur(20px)";
      document.body.style.overflow = "";
      window.setTimeout(function () {
        el.style.display = "none";
      }, 800);
    }

    window.setTimeout(dismiss, 1800);
  })();

  /* ----------------------------------------------------------------
     Hero: typing effect + CV availability check
     ---------------------------------------------------------------- */
  (function hero() {
    var roleEl = document.getElementById("hero-role");
    function getRoles() {
      if (config.roles && config.roles[currentLang]) {
        return config.roles[currentLang].slice();
      }
      return ["Pengembang Web", "Dukungan IT", "Frontend Developer"];
    }

    var roles = getRoles();
    var index = 0;
    var subIndex = 0;
    var deleting = false;

    if (roleEl) {
      var tick = function () {
        if (!roles.length) return;
        var text = roles[index % roles.length];
        if (!deleting && subIndex < text.length) {
          subIndex++;
        } else if (deleting && subIndex > 0) {
          subIndex--;
        } else if (!deleting && subIndex === text.length) {
          deleting = true;
        } else if (deleting && subIndex === 0) {
          deleting = false;
          index = (index + 1) % roles.length;
        }
        roleEl.textContent = roles[index % roles.length].substring(0, subIndex);
        window.setTimeout(tick, deleting ? 75 : 150);
      };

      window.setTimeout(tick, 150);

      // Reset typing seamlessly on language change
      i18nListeners.push(function () {
        roles = getRoles();
        index = 0;
        subIndex = 0;
        deleting = false;
      });
    }

    var link = document.getElementById("hero-cv-link");
    var disabled = document.getElementById("hero-cv-disabled");
    if (link && disabled && config.cvPath) {
      fetch(config.cvPath, { method: "HEAD" })
        .then(function (res) {
          if (res.ok) {
            link.classList.remove("hidden");
            disabled.classList.add("hidden");
          }
        })
        .catch(function () {});
    }
  })();

  /* ----------------------------------------------------------------
     Graduation gallery: expand + lightbox
     ---------------------------------------------------------------- */
  (function graduation() {
    var toggle = document.getElementById("graduation-toggle");
    var toggleLabel = document.getElementById("graduation-toggle-label");
    var toggleIcon = document.getElementById("graduation-toggle-icon");
    var items = document.querySelectorAll("[data-graduation-item]");
    var lightbox = document.getElementById("graduation-lightbox");
    var lightboxImg = document.getElementById("lightbox-img");
    var current = null;

    function getPhotos() {
      return Array.prototype.map.call(
        document.querySelectorAll("[data-graduation-item]"),
        function (item) {
          var img = item.querySelector("img");
          return {
            src: img.getAttribute("src"),
            alt: img.getAttribute("alt"),
          };
        },
      );
    }

    function updateToggleLabel() {
      if (!toggle || !toggleLabel) return;
      var expanded = toggle.getAttribute("aria-expanded") === "true";
      toggleLabel.textContent = expanded ? t("grad_show_less", "Tampilkan Lebih Sedikit") : t("grad_view_all", "Lihat Semua Foto");
    }

    if (toggle) {
      toggle.addEventListener("click", function () {
        var expanded = toggle.getAttribute("aria-expanded") === "true";
        expanded = !expanded;
        toggle.setAttribute("aria-expanded", expanded ? "true" : "false");
        updateToggleLabel();
        if (toggleIcon) toggleIcon.classList.toggle("rotate-180", expanded);
        items.forEach(function (item, i) {
          if (i >= 3) item.classList.toggle("hidden", !expanded);
        });
      });
    }

    i18nListeners.push(updateToggleLabel);

    function showLightbox(i) {
      var photos = getPhotos();
      if (!lightbox || !lightboxImg || !photos[i]) return;
      current = i;
      lightboxImg.setAttribute("src", photos[i].src);
      lightboxImg.setAttribute("alt", photos[i].alt);
      lightbox.classList.remove("hidden");
      lightbox.classList.add("flex");
      document.body.style.overflow = "hidden";
    }

    function hideLightbox() {
      if (!lightbox) return;
      lightbox.classList.add("hidden");
      lightbox.classList.remove("flex");
      document.body.style.overflow = "";
      current = null;
    }

    function step(delta) {
      var photos = getPhotos();
      if (current === null || !photos.length) return;
      showLightbox((current + delta + photos.length) % photos.length);
    }

    document.querySelectorAll("[data-lightbox-open]").forEach(function (btn) {
      btn.addEventListener("click", function () {
        showLightbox(parseInt(btn.getAttribute("data-lightbox-open"), 10));
      });
    });
    document.querySelectorAll("[data-lightbox-close]").forEach(function (btn) {
      btn.addEventListener("click", hideLightbox);
    });
    var prev = document.querySelector("[data-lightbox-prev]");
    var next = document.querySelector("[data-lightbox-next]");
    if (prev) prev.addEventListener("click", function () { step(-1); });
    if (next) next.addEventListener("click", function () { step(1); });

    document.addEventListener("keydown", function (e) {
      if (current === null) return;
      if (e.key === "Escape") hideLightbox();
      if (e.key === "ArrowRight") step(1);
      if (e.key === "ArrowLeft") step(-1);
    });
  })();

  /* ----------------------------------------------------------------
     Experience: scroll-driven timeline progress
     ---------------------------------------------------------------- */
  (function experienceTimeline() {
    var track = document.getElementById("experience-track");
    var progress = document.getElementById("experience-progress");
    if (!track || !progress) return;

    var target = 0;
    var current = 0;
    var running = false;

    function measure() {
      var rect = track.getBoundingClientRect();
      var center = window.innerHeight / 2;
      var value = (center - rect.top) / rect.height;
      target = Math.max(0, Math.min(1, value));
      if (!running) {
        running = true;
        requestAnimationFrame(frame);
      }
    }

    function frame() {
      current += (target - current) * 0.18;
      if (Math.abs(target - current) < 0.001) {
        current = target;
        running = false;
      }
      progress.style.transform = "scaleY(" + current + ")";
      if (running) requestAnimationFrame(frame);
    }

    window.addEventListener("scroll", measure, { passive: true });
    window.addEventListener("resize", measure);
    measure();
  })();

  /* ----------------------------------------------------------------
     ScrollVelocity marquee (scroll-reactive, hover-independent)
     ---------------------------------------------------------------- */
  (function marquee() {
    var marquees = document.querySelectorAll(".scroll-velocity");
    if (!marquees.length) return;

    var lastY = window.scrollY;
    var lastT = performance.now();
    var velocity = 0;

    window.addEventListener(
      "scroll",
      function () {
        var now = performance.now();
        var dt = Math.max(1, now - lastT);
        velocity = ((window.scrollY - lastY) / dt) * 1000;
        lastY = window.scrollY;
        lastT = now;
      },
      { passive: true },
    );

    marquees.forEach(function (marquee) {
      var track = marquee.querySelector(".scroll-velocity-track");
      var measure = marquee.querySelector("[data-measure]");
      if (!track || !measure) return;

      var baseVelocity = parseFloat(marquee.getAttribute("data-velocity") || "50");
      var baseX = 0;
      var direction = 1;
      var copyWidth = 0;
      var lastFrame = performance.now();

      function updateWidth() {
        copyWidth = measure.offsetWidth;
      }

      function wrap(min, max, v) {
        var range = max - min;
        return (((v - min) % range) + range) % range + min;
      }

      function frame(now) {
        var factor = velocity / 200;
        if (factor < 0) direction = -1;
        else if (factor > 0) direction = 1;

        var dt = Math.min(64, Math.max(1, now - lastFrame));
        lastFrame = now;
        var moveBy = direction * baseVelocity * (dt / 1000);
        moveBy += direction * moveBy * factor;
        baseX += moveBy;

        if (copyWidth > 0) {
          track.style.transform = "translateX(" + wrap(-copyWidth, 0, baseX) + "px)";
        }
        velocity *= 0.9;
        requestAnimationFrame(frame);
      }

      window.addEventListener("resize", updateWidth);
      i18nListeners.push(function () {
        // give browser a moment to layout new translated string width
        requestAnimationFrame(updateWidth);
      });
      updateWidth();
      requestAnimationFrame(frame);
    });
  })();

  /* ----------------------------------------------------------------
     GlareHover sweep
     ---------------------------------------------------------------- */
  (function glare() {
    document.querySelectorAll(".glare-hover").forEach(function (host) {
      var overlay = host.querySelector(".glare-overlay");
      if (!overlay) return;
      host.addEventListener("mouseenter", function () {
        overlay.style.transition = "none";
        overlay.style.backgroundPosition = "-100% -100%, 0 0";
        requestAnimationFrame(function () {
          overlay.style.transition = "650ms ease";
          overlay.style.backgroundPosition = "100% 100%, 0 0";
        });
      });
      host.addEventListener("mouseleave", function () {
        overlay.style.transition = "650ms ease";
        overlay.style.backgroundPosition = "-100% -100%, 0 0";
      });
    });
  })();

  /* ----------------------------------------------------------------
     Project modals + gallery
     ---------------------------------------------------------------- */
  (function projectModals() {
    var modals = document.querySelectorAll(".project-modal");
    if (!modals.length) return;

    function openModal(modal) {
      modal.classList.remove("hidden");
      modal.classList.add("flex");
      requestAnimationFrame(function () {
        var backdrop = modal.querySelector(".project-modal-backdrop");
        var panel = modal.querySelector(".project-modal-panel");
        if (backdrop) backdrop.classList.replace("opacity-0", "opacity-100");
        if (panel) {
          panel.classList.remove("opacity-0", "scale-95", "translate-y-5");
          panel.classList.add("opacity-100", "scale-100", "translate-y-0");
        }
      });
      document.body.style.overflow = "hidden";
    }

    function closeModal(modal) {
      var backdrop = modal.querySelector(".project-modal-backdrop");
      var panel = modal.querySelector(".project-modal-panel");
      if (backdrop) backdrop.classList.replace("opacity-100", "opacity-0");
      if (panel) {
        panel.classList.remove("opacity-100", "scale-100", "translate-y-0");
        panel.classList.add("opacity-0", "scale-95", "translate-y-5");
      }
      window.setTimeout(function () {
        modal.classList.add("hidden");
        modal.classList.remove("flex");
      }, 300);
      document.body.style.overflow = "";
    }

    document.querySelectorAll("[data-project-open]").forEach(function (btn) {
      btn.addEventListener("click", function () {
        var modal = document.querySelector('[data-project-modal="' + btn.getAttribute("data-project-open") + '"]');
        if (modal) openModal(modal);
      });
    });

    modals.forEach(function (modal) {
      modal.querySelectorAll("[data-project-close]").forEach(function (btn) {
        btn.addEventListener("click", function () {
          closeModal(modal);
        });
      });

      var main = modal.querySelector("[data-gallery-main]");
      var thumbs = modal.querySelectorAll("[data-gallery-thumb]");
      thumbs.forEach(function (thumb) {
        thumb.addEventListener("click", function () {
          var img = thumb.querySelector("img");
          if (main && img) main.setAttribute("src", img.getAttribute("src"));
          thumbs.forEach(function (t) {
            var active = t === thumb;
            t.classList.toggle("border-text-primary", active);
            t.classList.toggle("opacity-100", active);
            t.classList.toggle("border-text-secondary/20", !active);
            t.classList.toggle("opacity-60", !active);
            t.setAttribute("aria-current", active ? "true" : "false");
          });
        });
      });
    });

    document.addEventListener("keydown", function (e) {
      if (e.key !== "Escape") return;
      modals.forEach(function (modal) {
        if (!modal.classList.contains("hidden")) closeModal(modal);
      });
    });
  })();

  /* ----------------------------------------------------------------
     Certificate modals
     ---------------------------------------------------------------- */
  (function certificateModals() {
    var modals = document.querySelectorAll(".certificate-modal");
    if (!modals.length) return;

    function openModal(modal) {
      modal.classList.remove("hidden");
      modal.classList.add("flex");
      document.body.style.overflow = "hidden";
      var close = modal.querySelector("[data-certificate-close]");
      if (close) close.focus();
    }

    function closeModal(modal) {
      modal.classList.add("hidden");
      modal.classList.remove("flex");
      document.body.style.overflow = "";
    }

    document.querySelectorAll("[data-certificate-open]").forEach(function (btn) {
      btn.addEventListener("click", function () {
        var modal = document.querySelector('[data-certificate-modal="' + btn.getAttribute("data-certificate-open") + '"]');
        if (modal) openModal(modal);
      });
    });

    modals.forEach(function (modal) {
      modal.querySelectorAll("[data-certificate-close]").forEach(function (btn) {
        btn.addEventListener("click", function () {
          closeModal(modal);
        });
      });
    });

    document.addEventListener("keydown", function (e) {
      if (e.key !== "Escape") return;
      modals.forEach(function (modal) {
        if (!modal.classList.contains("hidden")) closeModal(modal);
      });
    });
  })();

  /* ----------------------------------------------------------------
     Chat assistant (bilingual greeting, prompt instructions, errors)
     ---------------------------------------------------------------- */
  (function chat() {
    var overlay = document.getElementById("chat-overlay");
    var panel = document.getElementById("chat-panel");
    var list = document.getElementById("chat-messages");
    var typing = document.getElementById("chat-typing");
    var warning = document.getElementById("chat-config-warning");
    var form = document.getElementById("chat-form");
    var input = document.getElementById("chat-input");
    var sendBtn = document.getElementById("chat-send");
    var end = document.getElementById("chat-end");
    if (!overlay || !list) return;

    var STORAGE_KEY = config.storageKey || "deryl_chat_messages";
    var loading = false;
    var configured = null;

    function welcomeText() {
      return t("chat_welcome_text", "Halo! Gue asisten AI Deryl. Ada yang bisa gue bantu soal portofolio, pengalaman, atau project Deryl?");
    }

    function welcome() {
      return {
        id: "welcome-msg",
        sender: "bot",
        text: welcomeText(),
        timestamp: new Date(0),
      };
    }

    function read() {
      try {
        var raw = localStorage.getItem(STORAGE_KEY);
        if (!raw) return [];
        var parsed = JSON.parse(raw);
        return parsed.map(function (m) {
          return { id: m.id, sender: m.sender, text: m.text, timestamp: new Date(m.timestamp) };
        });
      } catch (e) {
        return [];
      }
    }

    function write(messages) {
      try {
        localStorage.setItem(STORAGE_KEY, JSON.stringify(messages));
      } catch (e) {}
    }

    var messages = read();
    if (!messages.length) {
      messages = [welcome()];
    } else if (messages.length === 1 && messages[0].id === "welcome-msg") {
      messages[0].text = welcomeText();
    }

    // When language changes and only welcome message is present, update it seamlessly
    i18nListeners.push(function () {
      if (messages.length === 1 && messages[0].id === "welcome-msg") {
        messages[0].text = welcomeText();
        render();
      }
    });

    function escapeHtml(s) {
      return String(s).replace(/[&<>"']/g, function (c) {
        return { "&": "&amp;", "<": "&lt;", ">": "&gt;", '"': "&quot;", "'": "&#39;" }[c];
      });
    }

    function inline(s) {
      var out = escapeHtml(s);
      out = out.replace(/`([^`]+)`/g, function (m, c) {
        return '<code class="bg-thirdary/30 px-1.5 py-0.5 rounded text-xs text-text-primary font-mono font-bold">' + c + "</code>";
      });
      out = out.replace(/\*\*([^*]+)\*\*/g, '<strong class="font-bold text-text-primary">$1</strong>');
      out = out.replace(/\[([^\]]+)\]\(([^)\s]+)\)/g, function (m, label, url) {
        if (!/^(https?:|mailto:)/i.test(url)) return label;
        return '<a class="text-text-primary underline hover:opacity-80 font-bold" target="_blank" rel="noopener noreferrer" href="' + url + '">' + label + "</a>";
      });
      return out;
    }

    function renderMarkdown(src) {
      var lines = String(src).replace(/\r\n/g, "\n").split("\n");
      var html = "";
      var i = 0;
      while (i < lines.length) {
        var line = lines[i];
        if (/^\s*$/.test(line)) {
          i++;
          continue;
        }
        if (/^\s*[-*]\s+/.test(line)) {
          var items = "";
          while (i < lines.length && /^\s*[-*]\s+/.test(lines[i])) {
            items += "<li>" + inline(lines[i].replace(/^\s*[-*]\s+/, "")) + "</li>";
            i++;
          }
          html += '<ul class="list-disc ml-4 mb-2 space-y-1">' + items + "</ul>";
          continue;
        }
        if (/^\s*\d+\.\s+/.test(line)) {
          var oitems = "";
          while (i < lines.length && /^\s*\d+\.\s+/.test(lines[i])) {
            oitems += "<li>" + inline(lines[i].replace(/^\s*\d+\.\s+/, "")) + "</li>";
            i++;
          }
          html += '<ol class="list-decimal ml-4 mb-2 space-y-1">' + oitems + "</ol>";
          continue;
        }
        html += '<p class="mb-2 last:mb-0">' + inline(line) + "</p>";
        i++;
      }
      return html;
    }

    function buildMessage(msg) {
      var wrap = document.createElement("div");
      wrap.className = "flex w-full " + (msg.sender === "user" ? "justify-end" : "justify-start");
      wrap.setAttribute("data-msg-id", msg.id);

      var bubble = document.createElement("div");
      bubble.className =
        "max-w-[85%] rounded-2xl px-5 py-3 " +
        (msg.sender === "user"
          ? "bg-text-primary text-background rounded-tr-sm"
          : "bg-background border border-text-secondary/10 text-text-primary rounded-tl-sm shadow-sm");

      var body = document.createElement("div");
      if (msg.sender === "user") {
        body.innerHTML = '<p class="text-sm font-medium leading-relaxed whitespace-pre-wrap">' + escapeHtml(msg.text) + "</p>";
      } else {
        body.className = "text-sm font-medium leading-relaxed overflow-hidden";
        body.innerHTML = renderMarkdown(msg.text);
      }

      var time = document.createElement("span");
      time.className =
        "text-[10px] uppercase font-bold tracking-wider mt-2 block " +
        (msg.sender === "user" ? "text-background/70" : "text-text-secondary/70");
      time.textContent = new Date(msg.timestamp).toLocaleTimeString(currentLang === "en" ? "en-US" : "id-ID", { hour: "2-digit", minute: "2-digit" });

      bubble.appendChild(body);
      bubble.appendChild(time);
      wrap.appendChild(bubble);
      return { wrap: wrap, body: body };
    }

    function render() {
      list.innerHTML = "";
      messages.forEach(function (msg) {
        list.appendChild(buildMessage(msg).wrap);
      });
    }

    function renderWarning() {
      if (!warning) return;
      warning.classList.toggle("hidden", configured !== false);
    }

    function scrollToBottom() {
      if (end) end.scrollIntoView({ behavior: "smooth" });
    }

    function openChat() {
      overlay.classList.remove("opacity-0", "invisible", "pointer-events-none");
      overlay.classList.add("opacity-100", "visible");
      var backdrop = overlay.querySelector("[data-chat-close]");
      if (backdrop) backdrop.classList.replace("opacity-0", "opacity-100");
      if (panel) {
        panel.classList.remove("translate-y-12", "scale-95", "opacity-0");
        panel.classList.add("translate-y-0", "scale-100", "opacity-100");
      }
      document.body.style.overflow = "hidden";
      window.setTimeout(function () {
        scrollToBottom();
        if (input) input.focus();
      }, 100);
    }

    function closeChat() {
      overlay.classList.add("opacity-0", "invisible", "pointer-events-none");
      overlay.classList.remove("opacity-100", "visible");
      var backdrop = overlay.querySelector("[data-chat-close]");
      if (backdrop) backdrop.classList.replace("opacity-100", "opacity-0");
      if (panel) {
        panel.classList.remove("translate-y-0", "scale-100", "opacity-100");
        panel.classList.add("translate-y-12", "scale-95", "opacity-0");
      }
      document.body.style.overflow = "";
    }

    document.querySelectorAll("[data-chat-open]").forEach(function (btn) {
      btn.addEventListener("click", openChat);
    });
    overlay.querySelectorAll("[data-chat-close]").forEach(function (btn) {
      btn.addEventListener("click", closeChat);
    });
    document.querySelectorAll("[data-chat-clear]").forEach(function (btn) {
      btn.addEventListener("click", function () {
        messages = [welcome()];
        try {
          localStorage.removeItem(STORAGE_KEY);
        } catch (e) {}
        render();
      });
    });

    document.addEventListener("keydown", function (e) {
      if (e.key === "Escape" && !overlay.classList.contains("invisible")) closeChat();
    });

    if (input) {
      input.addEventListener("input", function () {
        if (sendBtn) sendBtn.disabled = !input.value.trim() || loading;
      });
    }

    function systemPrompt() {
      if (config.systemPrompts && config.systemPrompts[currentLang]) {
        return config.systemPrompts[currentLang];
      }
      return config.systemPrompts ? config.systemPrompts.id : "";
    }

    function setLoading(value) {
      loading = value;
      if (typing) typing.classList.toggle("hidden", !value);
      if (input) input.disabled = value;
      if (sendBtn) sendBtn.disabled = value || !input || !input.value.trim();
    }

    if (form) {
      form.addEventListener("submit", function (e) {
        e.preventDefault();
        if (!input || !input.value.trim() || loading) return;

        var userText = input.value.trim();
        var history = [
          { role: "system", content: systemPrompt() },
        ].concat(
          messages
            .filter(function (m) {
              return m.id !== "welcome-msg";
            })
            .map(function (m) {
              return { role: m.sender === "user" ? "user" : "assistant", content: m.text };
            }),
        );

        messages.push({ id: Date.now().toString(), sender: "user", text: userText, timestamp: new Date() });
        write(messages);
        render();
        scrollToBottom();
        input.value = "";
        setLoading(true);

        fetch("/api/chat", {
          method: "POST",
          headers: { "Content-Type": "application/json" },
          body: JSON.stringify({ text: userText, history: history }),
        })
          .then(function (response) {
            if (!response.ok) throw new Error("Network response was not ok");
            var contentType = response.headers.get("Content-Type") || "";
            if (contentType.indexOf("application/json") !== -1) {
              return response.json().then(function (data) {
                throw new Error((data && data.message) || "Error dari server");
              });
            }
            configured = true;
            renderWarning();
            setLoading(false);

            var botMsg = { id: (Date.now() + 1).toString(), sender: "bot", text: "", timestamp: new Date() };
            messages.push(botMsg);
            write(messages);
            var node = buildMessage(botMsg);
            list.appendChild(node.wrap);
            scrollToBottom();

            var reader = response.body.getReader();
            var decoder = new TextDecoder();
            var botText = "";
            function pump() {
              return reader.read().then(function (result) {
                if (result.done) return;
                botText += decoder.decode(result.value, { stream: true });
                botMsg.text = botText;
                node.body.innerHTML = renderMarkdown(botText);
                write(messages);
                scrollToBottom();
                return pump();
              });
            }
            return pump();
          })
          .catch(function (error) {
            var message = error && error.message ? error.message : "Terjadi kesalahan.";
            if (message.toLowerCase().indexOf("konfigurasi") !== -1) {
              configured = false;
              renderWarning();
            }
            var errText = message.toLowerCase().indexOf("konfigurasi") !== -1
              ? t("chat_unconfigured_err", "Asisten AI belum dikonfigurasi pada server ini, jadi gue belum bisa menjawab otomatis.")
              : t("chat_conn_err_prefix", "Maaf, terjadi kesalahan koneksi atau server AI.") + " (" + message + ")";
            messages.push({
              id: (Date.now() + 1).toString(),
              sender: "bot",
              text: errText,
              timestamp: new Date(),
            });
            write(messages);
            render();
            scrollToBottom();
          })
          .finally(function () {
            setLoading(false);
          });
      });
    }

    render();
    renderWarning();
  })();
})();
